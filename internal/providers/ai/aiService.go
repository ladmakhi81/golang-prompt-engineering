package ai

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"questions-generators/internal/config"
	"questions-generators/internal/utils"
	"time"
)

func GenerateInterviewQuestions(jobTitle, industry, company, trends string, jobDescription string) []string {
	config.LoadEnv()
	openAIKey := config.GetEnv("OPENAI_API_KEY", "")
	currentYear := time.Now().Year()

	promptTemplate, err := utils.ReadFile("prompts/ai_prompt.txt")
	if err != nil {
		return []string{"Error: Could not read AI prompt file."}
	}

	prompt := utils.ReplacePlaceholders(promptTemplate, map[string]string{
		"{jobTitle}":       jobTitle,
		"{company}":        company,
		"{year}":           fmt.Sprintf("%d", currentYear),
		"{industry}":       industry,
		"{trends}":         trends,
		"{jobDescription}": jobDescription,
	})

	fmt.Println(prompt)

	requestBody, _ := json.Marshal(map[string]interface{}{
		"model": "gpt-4-turbo",
		"messages": []map[string]string{
			{"role": "system", "content": "You are an expert in interview preparation across all industries. Return questions in JSON array format."},
			{"role": "user", "content": prompt},
		},
	})

	req, _ := http.NewRequest("POST", "https://api.openai.com/v1/chat/completions", bytes.NewBuffer(requestBody))
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
