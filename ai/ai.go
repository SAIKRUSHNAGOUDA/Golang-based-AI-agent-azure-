package ai

import (
	"context"
	"fmt"
	"os"
	"strings"

	"github.com/SAIKRUSHNAGOUDA/Golang-based-AI-agent-azure/azure"
	"github.com/sashabaranov/go-openai"
	"github.com/joho/godotenv"
)

func GenerateAIResponse(question string, resources []azure.Resource) string {
	// Load .env if present
	_ = godotenv.Load()

	apiKey := os.Getenv("OPENAI_API_KEY")
	if apiKey == "" {
		return "❌ Missing OpenAI API key in environment variables."
	}

	client := openai.NewClient(apiKey)

	// Build resource summary for prompt context
	var resourceDetails []string
	for _, res := range resources {
		resourceDetails = append(resourceDetails,
			fmt.Sprintf("- Name: %s | Type: %s | Location: %s", res.Name, res.Type, res.Location))
	}
	resourceSummary := strings.Join(resourceDetails, "\n")

	// Construct prompt
	prompt := fmt.Sprintf(`
You are a helpful assistant with access to Azure cloud resource data.
Here's the list of resources:
%s

Answer the following user question based on this data:
Q: %s
A:`, resourceSummary, question)

	// Create chat request
	resp, err := client.CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest{
			Model: openai.GPT3Dot5Turbo,
			Messages: []openai.ChatCompletionMessage{
				{Role: "system", Content: "You are a cloud assistant that answers based on provided resource data."},
				{Role: "user", Content: prompt},
			},
		},
	)

	if err != nil {
		return fmt.Sprintf("❌ OpenAI API error: %v", err)
	}

	return strings.TrimSpace(resp.Choices[0].Message.Content)
}
