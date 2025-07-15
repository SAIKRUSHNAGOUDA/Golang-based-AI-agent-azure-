package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"path/filepath"

	"github.com/SAIKRUSHNAGOUDA/Golang-based-AI-agent-azure/azure"
)

func main() {
	// Serve static frontend files (index.html, etc.)
	fs := http.FileServer(http.Dir("./frontend"))
	http.Handle("/", fs)

	// API endpoint for resource data
	http.HandleFunc("/api/resources", func(w http.ResponseWriter, r *http.Request) {
		subID := "98f3c311-5766-420d-a7d5-7ef36868b7ef"
		resources := azure.FetchResources(subID)

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(resources)
	})

	fmt.Println("🌐 Server running at: http://localhost:8080")
	fmt.Println("📊 Open UI at:        http://localhost:8080/index.html")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
