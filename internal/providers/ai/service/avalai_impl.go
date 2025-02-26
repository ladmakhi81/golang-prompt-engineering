package aiservice

import (
	"context"
	"encoding/json"
	"questions-generators/internal/config"

	"github.com/sashabaranov/go-openai"
)

type AvalaiService struct{}

func NewAvalaiService() AvalaiService {
	return AvalaiService{}
}

func (svc AvalaiService) GenerateInterviewQuestions(model, prompt string) []string {
	config.LoadEnv()
	openAIKey := config.GetEnv("OPENAI_API_KEY", "")
	config := openai.DefaultConfig(openAIKey)
	config.BaseURL = "https://api.avalai.ir/v1"
	client := openai.NewClientWithConfig(config)
	resp, respErr := client.CreateChatCompletion(context.Background(), openai.ChatCompletionRequest{
		Model: model,
		Messages: []openai.ChatCompletionMessage{
			{
				Role:    "system",
				Content: "You are an expert in interview preparation across all industries. return questions in JSON array format.",
			},
			{
				Role:    "user",
				Content: prompt,
			},
		},
	})
	if respErr != nil || len(resp.Choices) == 0 {
		return []string{"Error: No response from AI."}
	}
	var questions []string
	parseErr := json.Unmarshal([]byte(resp.Choices[0].Message.Content), &questions)
	if parseErr != nil {
		return []string{"Error: AI response is not in JSON array format."}
	}
	return questions
}
