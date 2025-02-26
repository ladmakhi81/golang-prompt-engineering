package aiservice

import (
	"bytes"
	"encoding/json"
	"net/http"
	"questions-generators/internal/config"
)

type ChatgptService struct{}

func NewChatgptService() ChatgptService {
	return ChatgptService{}
}

func (svc ChatgptService) GenerateInterviewQuestions(model, prompt string) []string {
	config.LoadEnv()
	openAIKey := config.GetEnv("OPENAI_API_KEY", "")
	req, _ := http.NewRequest("POST", "https://api.openai.com/v1/chat/completions", bytes.NewBuffer([]byte(prompt)))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+openAIKey)
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return []string{"Error: Unable to fetch AI-generated questions."}
	}
	defer resp.Body.Close()
	var result struct {
		Choices []struct {
			Message struct {
				Content string `json:"content"`
			} `json:"message"`
		} `json:"choices"`
	}
	json.NewDecoder(resp.Body).Decode(&result)
	if len(result.Choices) == 0 {
		return []string{"Error: No response from AI."}
	}
	var questions []string
	err = json.Unmarshal([]byte(result.Choices[0].Message.Content), &questions)
	if err != nil {
		return []string{"Error: AI response is not in JSON array format."}
	}
	return questions
}
