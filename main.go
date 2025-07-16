package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	openai "github.com/sashabaranov/go-openai"
	"github.com/SAIKRUSHNAGOUDA/Golang-based-AI-agent-azure/azure"
)

func main() {
	_ = godotenv.Load() // Load .env file

	apiKey := os.Getenv("OPENAI_API_KEY")
	if apiKey == "" {
		log.Fatal("❌ Missing OPENAI_API_KEY environment variable")
	}

	http.HandleFunc("/api/resources", handleResources)
	http.HandleFunc("/api/ask", handleAsk(apiKey))

	// Serve frontend files
	fs := http.FileServer(http.Dir("./frontend"))
	http.Handle("/", fs)

	fmt.Println("🌐 Server running at: http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func handleResources(w http.ResponseWriter, r *http.Request) {
	subID := "98f3c311-5766-420d-a7d5-7ef36868b7ef" // Replace with your actual subscription ID
	resources := azure.FetchResources(subID)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resources)
}

func handleAsk(apiKey string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var body struct {
			Question string `json:"question"`
		}
		json.NewDecoder(r.Body).Decode(&body)

		subID := "98f3c311-5766-420d-a7d5-7ef36868b7ef"
		resources := azure.FetchResources(subID)
		resourceSummary, _ := json.MarshalIndent(resources, "", "  ")

		client := openai.NewClient(apiKey)
		resp, err := client.CreateChatCompletion(r.Context(), openai.ChatCompletionRequest{
			Model: openai.GPT3Dot5Turbo,
			Messages: []openai.ChatCompletionMessage{
				{
					Role:    "system",
					Content: "You are an Azure assistant. Answer questions about Azure resources provided by the user.",
				},
				{
					Role:    "user",
					Content: fmt.Sprintf("Resources:\n%s\n\nQuestion: %s", resourceSummary, body.Question),
				},
			},
		})
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		answer := resp.Choices[0].Message.Content
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{"answer": answer})
	}
}
