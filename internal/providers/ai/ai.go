package ai

import (
	aidto "questions-generators/internal/providers/ai/dto"
	aiservice "questions-generators/internal/providers/ai/service"
	"questions-generators/internal/utils"
	"time"
)

type AiService struct {
	aiGeneratorSvc aiservice.AiServiceProvider
}

func NewAiService(aiGeneratorSvc aiservice.AiServiceProvider) AiService {
	return AiService{
		aiGeneratorSvc: aiGeneratorSvc,
	}
}

func (svc AiService) GenerateInterviewQuestions(
	dto aidto.GenerateInterviewQuestionDTO,
) []string {
	promptName := "prompts/avalai_prompt_template.txt"
	promptToken := map[string]any{
		"Company":          dto.Company,
		"Topic":            dto.JobTitle,
		"Industry":         dto.Industry,
		"JobDescription":   dto.JobDescription,
		"News":             dto.News,
		"CV":               dto.Cv,
		"WebSearchContent": dto.Trends,
		"IncludeCV":        dto.IncludeCv,
		"IncludeWebSearch": dto.IncludeWebSearch,
		"IncludeNews":      dto.IncludeNews,
		"Year":             time.Now().Year(),
	}
	prompt, promptErr := utils.ParsePromptTemplate(promptName, promptToken)
	if promptErr != nil {
		return []string{"Error: Unable to parse prompt"}
	}
	result := svc.aiGeneratorSvc.GenerateInterviewQuestions(dto.AiModel, prompt)
	return result
}
