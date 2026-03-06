package services

import (
	"SysTrace_Server/data"
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"
	"sync"
)

type Handler struct {
	mu      sync.RWMutex
	devices map[string]*data.Device
}

func NewHandler() *Handler {

	return &Handler{
		devices: make(map[string]*data.Device),
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

// Dashboard renders the SysTrace dashboard page.
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
	var m data.Device
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

	w.WriteHeader(http.StatusOK)
}

func (h *Handler) Status(w http.ResponseWriter, _ *http.Request) {
	w.WriteHeader(http.StatusOK)
}
