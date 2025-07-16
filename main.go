package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/SAIKRUSHNAGOUDA/Golang-based-AI-agent-azure/ai"
	"github.com/SAIKRUSHNAGOUDA/Golang-based-AI-agent-azure/azure"
)

func main() {
	_ = godotenv.Load()

	http.HandleFunc("/api/resources", func(w http.ResponseWriter, r *http.Request) {
		subID := "your-subscription-id" // Replace with your Azure subscription ID
		resources := azure.FetchResources(subID)

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(resources)
	})

	http.HandleFunc("/api/ask", func(w http.ResponseWriter, r *http.Request) {
		var input struct {
			Question string `json:"question"`
		}
		_ = json.NewDecoder(r.Body).Decode(&input)

		subID := "your-subscription-id"
		resources := azure.FetchResources(subID)
		answer := ai.ProcessQuestion(input.Question, resources)

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{"answer": answer})
	})

	// Serve frontend
	fs := http.FileServer(http.Dir("./frontend"))
	http.Handle("/", fs)

	fmt.Println("🌐 Server running at: http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
