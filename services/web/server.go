package web

import (
	"fmt"
	"net/http"
)

type Server struct {
}

func (s *Server) Start() {
	handler := NewHandler()

	http.HandleFunc("/", handler.Dashboard)
	http.HandleFunc("/metrics", handler.Metrics)
	http.HandleFunc("/status", handler.Status)
	http.HandleFunc("/devices", handler.Devices)

	// Static Dateien (CSS, JS, etc.) servieren
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("templates"))))

	fmt.Println("Server läuft auf :8080")

	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println("Server error:", err)
	}
}
