package promptservice

import (
	"context"
	"encoding/json"
	"fmt"
	promptdto "interview-generator/src/prompt/dto"
	"os"
	"strings"

	"github.com/go-resty/resty/v2"
	"github.com/sashabaranov/go-openai"
	"github.com/spf13/viper"
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
	apiUrl := viper.GetString("CHAT_API_URL")
	apiKey := viper.GetString("CHAT_API_KEY")
	body := promptdto.NewSendPromptPayloadDTO("gpt-4o-mini",
		promptdto.NewMessage("system", "You are an expert in interview preparation across all industries. return questions in JSON array format."),
		promptdto.NewMessage("user", prompt),
	)
	client := resty.New().SetProxy("http://127.0.0.1:12334/")
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

func x() {
	apiKey := "aa-SCmxsmr0nsR91N5mpWMfDGiBPN7ihlUbHqkMIZYhHD67KK8v"
	baseUrl := "https://api.avalai.ir/v1"

	config := openai.DefaultConfig(apiKey)
	config.BaseURL = baseUrl
	client := openai.NewClientWithConfig(config)
	client.CreateChatCompletion(context.Background(), openai.ChatCompletionRequest{
		Model:    openai.GPT4Turbo,
		Messages: []openai.ChatCompletionMessage{},
	})
}
