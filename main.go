package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/SAIKRUSHNAGOUDA/Golang-based-AI-agent-azure/azure"
)

func main() {
	http.HandleFunc("/api/resources", func(w http.ResponseWriter, r *http.Request) {
		subID := "98f3c311-5766-420d-a7d5-7ef36868b7ef"
		resources := azure.FetchResources(subID)

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(resources)
	})

	fmt.Println("🌐 Server running on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
