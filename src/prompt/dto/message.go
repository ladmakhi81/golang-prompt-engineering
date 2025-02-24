package promptdto

type Message struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

func NewMessage(role, content string) Message {
	return Message{
		Role:    role,
		Content: content,
	}
}

type PromptChoice struct {
	Message Message `json:"message"`
}

type GetPromptMessageResponse struct {
	Model   string         `json:"model"`
	Choices []PromptChoice `json:"choices"`
}

type SendPromptPayloadDTO struct {
	Model    string    `json:"model"`
	Messages []Message `json:"messages"`
}

func NewSendPromptPayloadDTO(model string, messages ...Message) SendPromptPayloadDTO {
	return SendPromptPayloadDTO{
		Model:    model,
		Messages: messages,
	}
}
