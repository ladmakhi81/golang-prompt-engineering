package newssearch

import (
	"fmt"
	newssearchdto "questions-generators/internal/providers/newssearch/dto"
	newssearchservice "questions-generators/internal/providers/newssearch/service"
)

type NewsSearchService struct {
	newsSearchProvider newssearchservice.NewsSearchServiceProvider
}

func NewNewsSearchService(
	newsSearchProvider newssearchservice.NewsSearchServiceProvider,
) NewsSearchService {
	return NewsSearchService{
		newsSearchProvider: newsSearchProvider,
	}
}

func (svc NewsSearchService) FetchNews(
	dto newssearchdto.FetchNewsDTO,
) string {
	query := fmt.Sprintf("q=+%s OR +%s OR +%s", dto.JobTitle, dto.Industry, dto.Company)
	return svc.newsSearchProvider.FetchNews(query)
}
