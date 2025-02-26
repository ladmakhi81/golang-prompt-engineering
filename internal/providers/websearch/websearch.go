package websearch

import (
	"fmt"
	"net/url"
	websearchdto "questions-generators/internal/providers/websearch/dto"
	websearchservice "questions-generators/internal/providers/websearch/service"
	"questions-generators/internal/utils"
	"time"
)

type WebSearchService struct {
	webSearchProvider websearchservice.WebSearchServiceProvider
}

func NewWebSearchService(
	webSearchProvider websearchservice.WebSearchServiceProvider,
) WebSearchService {
	return WebSearchService{
		webSearchProvider: webSearchProvider,
	}
}

func (svc WebSearchService) FetchLatestTrends(
	dto websearchdto.FetchTrendsDTO,
) string {
	currentYear := time.Now().Year()
	queryTemplate, err := utils.ReadFile("prompts/websearch_query.txt")
	if err != nil {
		return "Error: Could not read Web Search query file."
	}
	query := utils.ReplacePlaceholders(queryTemplate, map[string]string{
		"{jobTitle}": dto.JobTitle,
		"{industry}": dto.Industry,
		"{company}":  dto.Company,
		"{year}":     fmt.Sprintf("%d", currentYear),
	})
	encodedQuery := url.QueryEscape(query)
	result := svc.webSearchProvider.FetchLatestTrends(encodedQuery)
	return result
}
