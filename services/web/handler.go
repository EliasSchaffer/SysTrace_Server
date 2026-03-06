package web

import (
	"SysTrace_Server/data/static"
	"SysTrace_Server/services/database"
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"
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

	return &Handler{
		devices: make(map[string]*static.Device),
	}
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

	h.mu.Lock()
	h.devices[m.ID] = &m
	h.mu.Unlock()

	fmt.Println("Device:", m.Hostname)
	fmt.Println("CPU:", m.Hardware.CPU)
	fmt.Println("RAM:", m.Hardware.MEMORY)

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

	fmt.Printf("Total devices: %d\n", len(h.devices))

	for id, device := range h.devices {
		// Nur Devices mit gültigen GPS-Daten zurückgeben (nicht 0,0)
		if device.GPS.Latitude == 0 && device.GPS.Longitude == 0 {
			fmt.Printf("Device %s hat keine gültigen GPS-Daten, wird übersprungen\n", id)
			continue
		}

		fmt.Printf("Device %s: Hostname=%s, GPS=(%.4f, %.4f)\n", id, device.Hostname, device.GPS.Latitude, device.GPS.Longitude)

		gpsData := map[string]interface{}{
			"lat":  device.GPS.Latitude,
			"lon":  device.GPS.Longitude,
			"name": device.Hostname,
		}
		gpsDataArray = append(gpsDataArray, gpsData)
	}

	fmt.Printf("Returning %d valid devices\n", len(gpsDataArray))

	jsonData, err := json.Marshal(gpsDataArray)
	if err != nil {
		http.Error(w, "Error encoding JSON", http.StatusInternalServerError)
		return
	}

	w.Write(jsonData)
}
