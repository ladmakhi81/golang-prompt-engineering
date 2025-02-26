package questionsvc

import (
	"questions-generators/internal/dto"
	"questions-generators/internal/providers/ai"
	aidto "questions-generators/internal/providers/ai/dto"
	"questions-generators/internal/providers/newssearch"
	newssearchdto "questions-generators/internal/providers/newssearch/dto"
	"questions-generators/internal/providers/websearch"
	websearchdto "questions-generators/internal/providers/websearch/dto"
	"questions-generators/internal/utils"
	"strings"
)

type QuestionService struct {
	websearchSvc  websearch.WebSearchService
	newssearchSvc newssearch.NewsSearchService
	aiSvc         ai.AiService
}

func NewQuestionService(
	websearchSvc websearch.WebSearchService,
	newssearchSvc newssearch.NewsSearchService,
	aiSvc ai.AiService,
) *QuestionService {
	return &QuestionService{
		websearchSvc:  websearchSvc,
		newssearchSvc: newssearchSvc,
		aiSvc:         aiSvc,
	}
}

func (svc *QuestionService) GetQuestions(dto dto.GenerateInterviewQuestionReqBody) []string {
	var news string
	var trends string
	features := strings.Split(dto.Features, ",")

	includeCV := utils.IncludeString(features, "cv")
	includeWebSearch := utils.IncludeString(features, "websearch")
	includeNews := utils.IncludeString(features, "news")

	if includeWebSearch {
		trends = svc.websearchSvc.FetchLatestTrends(
			websearchdto.NewFetchTrendsDTO(
				dto.JobTitle,
				dto.Industry,
				dto.Company,
			),
		)
	}

	if includeNews {
		news = svc.newssearchSvc.FetchNews(
			newssearchdto.NewFetchNewsDTO(
				dto.JobTitle,
				dto.Industry,
				dto.Company,
			),
		)
	}

	questions := svc.aiSvc.GenerateInterviewQuestions(
		aidto.NewGenerateInterviewQuestionDTO(
			dto.JobTitle,
			dto.Industry,
			dto.Company,
			dto.AiModel,
			dto.JobDescription,
			trends,
			news,
			dto.CvAsText,
			includeCV,
			includeWebSearch,
			includeNews,
			dto.CustomPrompt,
		),
	)

	return questions
}
