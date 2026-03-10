package web

import (
	"SysTrace_Server/data/static"
	handler2 "SysTrace_Server/services/handler"
	"fmt"
	"net/http"
)

type Server struct {
	Devices map[string]*static.Device
}

func (s *Server) Start() {
	handler := handler2.NewHandler()

	http.HandleFunc("/", handler.Dashboard)
	http.HandleFunc("/metrics", handler.Metrics)
	http.HandleFunc("/api/metrics", handler.Metrics)
	http.HandleFunc("/status", handler.Status)
	http.HandleFunc("/api/status", handler.Status)
	http.HandleFunc("/devices", handler.Devices)
	http.HandleFunc("/api/devices", handler.Devices)
	http.HandleFunc("/api/health", handler.DevicesHealth)
	http.HandleFunc("/api/device/", handler.DeviceHistory)
	http.HandleFunc("/api/ws/send", handler.SendToClient)
	http.HandleFunc("/device/", handler.DeviceDetailsPage)
	http.HandleFunc("/ws", handler.WebSocketHandler)

	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("templates"))))

	fmt.Println("Server läuft auf :8080")

	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println("Server error:", err)
	}
}
