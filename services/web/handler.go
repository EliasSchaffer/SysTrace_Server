package web

import (
	"SysTrace_Server/data/static"
	"SysTrace_Server/services/database"
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"
	"strings"
	"sync"
)

type Handler struct {
	mu      sync.RWMutex
	devices map[string]*static.Device
}

func NewHandler() *Handler {
	err := database.InitDatabase()
	if err != nil {
		fmt.Println("Error initializing database:", err)
	}

	h := &Handler{
		devices: make(map[string]*static.Device),
	}

	if database.IsConnected() {
		devices, err := database.LoadDevicesFromDatabase()
		if err != nil {
			fmt.Printf("Error loading devices from database: %v\n", err)
		} else {
			h.devices = devices
			fmt.Printf("Loaded %d devices from database\n", len(devices))
		}
	}

	return h
}

func (h *Handler) DeviceCount() int {
	h.mu.RLock()
	defer h.mu.RUnlock()
	return len(h.devices)
}

func (h *Handler) DataInput() string {
	h.mu.RLock()
	defer h.mu.RUnlock()

	var output string
	for _, device := range h.devices {
		output += fmt.Sprintf("Hostname: %s, CPU: %v, RAM: %v\n", device.Hostname, device.Hardware.CPU, device.Hardware.MEMORY)
	}

	return output

}

func (h *Handler) Dashboard(w http.ResponseWriter, _ *http.Request) {
	tmpl, err := template.ParseFiles("templates/index.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	viewData := struct {
		Title       string
		Heading     string
		DeviceCount int
		Data        string
	}{
		Title:       "SysTrace Dashboard",
		Heading:     "Welcome",
		DeviceCount: h.DeviceCount(),
		Data:        h.DataInput(),
	}

	if err := tmpl.Execute(w, viewData); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func (h *Handler) Metrics(w http.ResponseWriter, r *http.Request) {
	var m static.Device
	if err := json.NewDecoder(r.Body).Decode(&m); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	if m.ID == "" {
		http.Error(w, "Device ID is required", http.StatusBadRequest)
		return
	}

	h.mu.Lock()
	m.Active = true
	h.devices[m.ID] = &m
	h.mu.Unlock()

	if database.IsConnected() {
		go database.InsertFullDataSet("localhost", m)
	}

	w.WriteHeader(http.StatusOK)
}

func (h *Handler) Status(w http.ResponseWriter, _ *http.Request) {
	w.WriteHeader(http.StatusOK)
}

func (h *Handler) Devices(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	h.mu.RLock()
	defer h.mu.RUnlock()

	gpsDataArray := make([]map[string]interface{}, 0)

	for _, device := range h.devices {
		// Zeige alle Devices an, auch ohne GPS (0,0)
		gpsData := map[string]interface{}{
			"lat":  device.GPS.Latitude,
			"lon":  device.GPS.Longitude,
			"name": device.Hostname,
			"ip":   device.IP,
			"id":   device.ID,
		}
		gpsDataArray = append(gpsDataArray, gpsData)
	}

	jsonData, err := json.Marshal(gpsDataArray)
	if err != nil {
		http.Error(w, "Error encoding JSON", http.StatusInternalServerError)
		return
	}

	w.Write(jsonData)
}

func (h *Handler) DeviceDetailsPage(w http.ResponseWriter, r *http.Request) {
	deviceName := r.URL.Path[len("/device/"):]

	h.mu.RLock()
	defer h.mu.RUnlock()

	var device *static.Device
	for _, dev := range h.devices {
		if dev.Hostname == deviceName {
			device = dev
			break
		}
	}

	if device == nil {
		http.Error(w, "Device not found", http.StatusNotFound)
		return
	}

	tmpl, err := template.ParseFiles("templates/device_details.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	viewData := struct {
		DeviceID   string
		DeviceName string
		OS         string
		IP         string
		CPU        string
		RAM        string
		CurrentGPS string
	}{
		DeviceID:   device.ID,
		DeviceName: device.Hostname,
		OS:         device.OS,
		IP:         device.IP,
		CPU:        fmt.Sprintf("%s (%d Cores, %d Threads)", device.Hardware.CPU.Model, device.Hardware.CPU.Cores, device.Hardware.CPU.Threads),
		RAM:        fmt.Sprintf("%d GB (%s)", device.Hardware.MEMORY.Total/1024/1024/1024, device.Hardware.MEMORY.Model),
		CurrentGPS: fmt.Sprintf("%.6f, %.6f", device.GPS.Latitude, device.GPS.Longitude),
	}

	if err := tmpl.Execute(w, viewData); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func (h *Handler) DeviceHistory(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if !strings.HasSuffix(r.URL.Path, "/gps-history") {
		http.Error(w, "Invalid API endpoint", http.StatusNotFound)
		return
	}

	path := strings.TrimPrefix(r.URL.Path, "/api/device/")
	deviceID := strings.TrimSuffix(path, "/gps-history")

	history, err := database.GetGPSHistory(deviceID)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error fetching GPS history: %v", err), http.StatusInternalServerError)
		return
	}

	if history == nil {
		history = make([]map[string]interface{}, 0)
	}

	jsonData, err := json.Marshal(history)
	if err != nil {
		http.Error(w, "Error encoding JSON", http.StatusInternalServerError)
		return
	}

	w.Write(jsonData)
}

func (h *Handler) DevicesHealth(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	h.mu.Lock()
	defer h.mu.Unlock()

	healthData := make([]map[string]interface{}, 0)

	// Markiere Devices als offline, wenn sie länger als 30 Sekunden keine Daten gesendet haben
	for _, device := range h.devices {
		status := "offline"
		active := device.Active

		// Falls Active ist, aber länger als 30 Sekunden keine neuen Metrics kommen, setze offline
		// (Hier könnte man einen Timestamp speichern und vergleichen)
		if device.Active {
			status = "online"
		}

		healthInfo := map[string]interface{}{
			"deviceID": device.ID,
			"hostname": device.Hostname,
			"ip":       device.IP,
			"os":       device.OS,
			"status":   status,
			"active":   active,
		}
		healthData = append(healthData, healthInfo)
	}

	jsonData, err := json.Marshal(healthData)
	if err != nil {
		http.Error(w, "Error encoding JSON", http.StatusInternalServerError)
		return
	}

	w.Write(jsonData)
}
