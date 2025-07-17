package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/SAIKRUSHNAGOUDA/Golang-based-AI-agent-azure/ai"
	"github.com/SAIKRUSHNAGOUDA/Golang-based-AI-agent-azure/azure"
)

func main() {
	subscriptionID := "98f3c311-5766-420d-a7d5-7ef36868b7ef"

	// Load resources once for both UI and AI
	resources := azure.FetchResources(subscriptionID)

	http.HandleFunc("/api/resources", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(resources)
	})

	http.HandleFunc("/api/ask", func(w http.ResponseWriter, r *http.Request) {
		var req struct {
			Question string `json:"question"`
		}
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, "Invalid request", http.StatusBadRequest)
			return
		}

		answer := ai.GenerateAIResponse(req.Question, resources)

		resp := map[string]string{"answer": answer}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(resp)
	})

	// Serve frontend
	fs := http.FileServer(http.Dir("./frontend"))
	http.Handle("/", fs)

	fmt.Println("🌐 Server running at: http://localhost:8080")
	fmt.Println("📊 UI available at:   http://localhost:8080/index.html")

	log.Fatal(http.ListenAndServe(":8080", nil))
}
