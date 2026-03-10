package handler

import (
	"SysTrace_Server/data/static"
	"SysTrace_Server/data/ws"
	"SysTrace_Server/services/database"
	"encoding/json"
	"fmt"
)

type WSEventHandler struct {
}

func HandleEvent(eventString string) {
	var event ws.WSEvent
	err := json.Unmarshal([]byte(eventString), &event)
	if err != nil {
		fmt.Printf("Fehler beim Parsen des JSON: %v", err)
		return
	}

	switch event.Type {
	case "update":
		handleUpdateEvent(event.Device)
	case "device_connected":
		// TODO: Handle device connected event
	case "device_disconnected":
		// TODO: Handle device disconnected event
	default:
		// Handle unknown event types if necessary
	}
}

func handleUpdateEvent(device static.Device) {
	var m static.Device
	m = device
	m.Active = true

	if database.IsConnected() {
		go database.InsertFullDataSet("localhost", m)
	}

}
