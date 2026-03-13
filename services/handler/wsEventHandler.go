package handler

import (
	"SysTrace_Server/data/static"
	"SysTrace_Server/services/database"
	"encoding/json"
	"fmt"
)

type WSEventHandler struct {
}

func (h *Handler) HandleEvent(eventString string) {
	raw := []byte(eventString)

	var header struct {
		Type string `json:"type"`
	}
	if err := json.Unmarshal(raw, &header); err != nil {
		fmt.Printf("Fehler beim Parsen des Event-Typs: %v", err)
		return
	}
	if header.Type == "" {
		fmt.Printf("Event ohne Typ empfangen")
		return
	}

	switch header.Type {
	case "update":
		var update struct {
			Device static.Device `json:"device"`
		}
		if err := json.Unmarshal(raw, &update); err != nil {
			fmt.Printf("Fehler beim Parsen des Update-Events: %v", err)
			return
		}
		h.handleUpdateEvent(update.Device)
	case "response":
		return
	case "device_connected":
		// TODO: Handle device connected event
	case "device_disconnected":
		// TODO: Handle device disconnected event
	default:
		fmt.Printf("Unknown event type: %s\n", header.Type)
	}
}

func (h *Handler) handleUpdateEvent(device static.Device) {
	device.Active = true

	h.mu.Lock()
	h.devices[device.ID] = &device
	h.mu.Unlock()

	if database.IsConnected() {
		go database.InsertFullDataSet("localhost", device)
	}
}
