package ai

import (
	"context"
	"fmt"
	"os"

	openai "github.com/sashabaranov/go-openai"
	"github.com/SAIKRUSHNAGOUDA/Golang-based-AI-agent-azure/azure"
)

func ProcessQuestion(question string, resources []azure.Resource) string {
	apiKey := os.Getenv("OPENAI_API_KEY")
	if apiKey == "" {
		return "❌ Missing OpenAI API key"
	}

	client := openai.NewClient(apiKey)

	resourceList := ""
	for _, r := range resources {
		resourceList += fmt.Sprintf("Name: %s, Type: %s, Location: %s\n", r.Name, r.Type, r.Location)
	}

	prompt := fmt.Sprintf(`
You are a smart Azure assistant. These are the resources:

%s

User question: %s

If it matches any resource name/type/location, list them.
If not, say: ❌ No matching resource found.
`, resourceList, question)

	resp, err := client.CreateChatCompletion(context.Background(), openai.ChatCompletionRequest{
		Model: openai.GPT3Dot5Turbo,
		Messages: []openai.ChatCompletionMessage{
			{Role: "system", Content: "You are a helpful Azure cloud assistant."},
			{Role: "user", Content: prompt},
		},
	})

	if err != nil {
		return "❌ Error from OpenAI: " + err.Error()
	}

	return resp.Choices[0].Message.Content
}
