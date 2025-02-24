package promptservice

import (
	"encoding/json"
	"fmt"
	promptdto "interview-generator/src/prompt/dto"
	"os"
	"strings"

	"github.com/go-resty/resty/v2"
)

type PromptService struct{}

func NewPromptService() PromptService {
	return PromptService{}
}

func (PromptService) GetPrompt(promptName string, promptToken map[string]any) (string, error) {
	file, fileErr := os.ReadFile(fmt.Sprintf("./prompt-sample-message/%s", promptName))
	if fileErr != nil {
		return "", fmt.Errorf("Unable to load file : %v", fileErr)
	}
	content := string(file)
	for tokenName, tokenValue := range promptToken {
		content = strings.Replace(content, tokenName, tokenValue.(string), -1)
	}
	return content, nil
}

func (PromptService) SendPromptToChatgpt(prompt string) (*promptdto.GetPromptMessageResponse, error) {
	// TODO: update api key from env
	apiKey := ""
	apiUrl := "https://api.openai.com/v1/chat/completions"
	body := promptdto.NewSendPromptPayloadDTO("gpt-4o-mini", promptdto.NewMessage(prompt))
	client := resty.New()
	resp, respErr := client.R().
		SetHeader("content-type", "application/json").
		SetHeader("Authorization", fmt.Sprintf("Bearer %s", apiKey)).
		SetBody(body).
		Post(apiUrl)

	if respErr != nil {
		return nil, fmt.Errorf("error in response : %v", respErr)
	}
	result := new(promptdto.GetPromptMessageResponse)
	if err := json.Unmarshal(resp.Body(), result); err != nil {
		return nil, fmt.Errorf("error in convert data : %v", err)
	}
	return result, nil
}
