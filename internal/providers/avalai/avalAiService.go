package avalai

import (
	"context"
	"encoding/json"
	"questions-generators/internal/config"
	"questions-generators/internal/utils"
	"strings"
	"time"

	"github.com/sashabaranov/go-openai"
)

func GenerateInterviewQuestions(
	jobTitle,
	industry,
	company,
	aiModel,
	jobDescription string,
	trends,
	articles,
	cv string,
	includeCV,
	includeWebSearch,
	includeNews bool,
	externalPrompt string,

) []string {
	config.LoadEnv()
	openAIKey := config.GetEnv("OPENAI_API_KEY", "")
	token := map[string]any{
		"Company":          company,
		"Topic":            jobTitle,
		"Industry":         industry,
		"JobDescription":   jobDescription,
		"News":             articles,
		"CV":               cv,
		"WebSearchContent": trends,
		"IncludeCV":        includeCV,
		"IncludeWebSearch": includeWebSearch,
		"IncludeNews":      includeNews,
		"Year":             time.Now().Year(),
	}
	prompt, promptErr := utils.ParsePromptTemplate("prompts/avalai_prompt_template.txt", token, externalPrompt)
	if promptErr != nil {
		return []string{"Error: Unable to parse prompt"}
	}
	config := openai.DefaultConfig(openAIKey)
	config.BaseURL = "https://api.avalai.ir/v1"
	client := openai.NewClientWithConfig(config)
	resp, respErr := client.CreateChatCompletion(context.Background(), openai.ChatCompletionRequest{
		Model: aiModel,
		Messages: []openai.ChatCompletionMessage{
			{
				Role:    "system",
				Content: "You are an expert in interview preparation across all industries. return questions in JSON array format without any extra words and sentences.",
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

	messageContent := resp.Choices[0].Message.Content
	messageContent = strings.Replace(messageContent, "json", "", -1)
	messageContent = strings.Replace(messageContent, "`", "", -1)

	var questions []string
	parseErr := json.Unmarshal([]byte(messageContent), &questions)
	if parseErr != nil {
		return []string{"Error: AI response is not in JSON array format."}
	}
	return questions
}
