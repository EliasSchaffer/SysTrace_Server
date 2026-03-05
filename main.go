package main

import (
	"SysTrace_Server/services"
	"fmt"
	"net/http"
)

func main() {
	handler := services.NewHandler()

	http.HandleFunc("/", handler.Dashboard)
	http.HandleFunc("/metrics", handler.Metrics)
	http.HandleFunc("/status", handler.Status)

	fmt.Println("Server läuft auf :8080")

	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println("Server error:", err)
	}
}
