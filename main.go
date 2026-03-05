package main

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello SysTrace Dashboard")
}

func main() {
	http.HandleFunc("/", handler)

	fmt.Println("Server läuft auf :8080")

	http.HandleFunc("/metrics", metricsHandler)

	http.ListenAndServe(":8080", nil)

}

func metricsHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Metrics received")

	w.WriteHeader(http.StatusOK)
}
