package newsservice

import (
	"encoding/json"
	"fmt"
	newstype "interview-generator/src/news/type"

	"github.com/go-resty/resty/v2"
)

const (
	API_KEY = "b4433c3f2c9d4d9badb8db417268b91c"
	API_URL = "https://newsapi.org/v2/everything"
)

type NewsService struct{}

func NewNews() NewsService {
	return NewsService{}
}

func (NewsService) FetchLatestNews(topic string) (*newstype.NewsResponse, error) {
	client := resty.New()
	requestURL := fmt.Sprintf("%s?q=%s&apiKey=%s&language=en&pageSize=100&sortBy=publishedAt", API_URL, topic, API_KEY)
	resp, respErr := client.R().
		Get(requestURL)
	if respErr != nil {
		return nil, fmt.Errorf("error in response : %v", respErr)
	}
	result := new(newstype.NewsResponse)
	if err := json.Unmarshal(resp.Body(), result); err != nil {
		return nil, fmt.Errorf("error in convert data : %v", err)
	}
	return result, nil
}

func (NewsService) GetTitlesOfArticles(articles []newstype.NewsArticle) []string {
	titles := make([]string, len(articles))
	for index, article := range articles {
		titles[index] = article.Title
	}
	return titles
}
