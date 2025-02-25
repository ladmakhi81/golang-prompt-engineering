package v1

import (
	"questions-generators/internal/providers/avalai"
	"questions-generators/internal/providers/newsapi"
	"questions-generators/internal/providers/websearch"
	"questions-generators/internal/utils"
)

type QuestionService struct{}

func NewQuestionService() *QuestionService {
	return &QuestionService{}
}

// func (q *QuestionService) GetQuestions(jobTitle, industry, company string, jobDescription string) []string {
// 	trends := websearch.FetchLatestTrends(jobTitle, industry, company)
// 	questions := ai.GenerateInterviewQuestions(jobTitle, industry, company, trends, jobDescription)
// 	return questions
// }

func (q *QuestionService) GetQuestions(jobTitle, industry, company, jobDescription, aiModel, cv string, features []string) []string {
	var news string
	var trends string

	includeCV := utils.IncludeString(features, "cv")
	includeWebSearch := utils.IncludeString(features, "websearch")
	includeNews := utils.IncludeString(features, "news")

	if includeNews {
		news = newsapi.FetchLatestTrends(jobTitle, industry, company)
	}

	if includeWebSearch {
		trends = websearch.FetchLatestTrends(jobTitle, industry, company)
	}

	questions := avalai.GenerateInterviewQuestions(
		jobTitle,
		industry,
		company,
		aiModel,
		jobDescription,
		trends,
		news,
		cv,
		includeCV,
		includeWebSearch,
		includeNews,
	)
	return questions
}
