package newsapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"questions-generators/internal/config"
	"time"
)

func FetchLatestTrends(jobTitle, industry, company string) string {
	config.LoadEnv()
	newsAPIKey := config.GetEnv("NEWS_API_KEY", "")
	query := fmt.Sprintf("q=+%s OR +%s OR +%s", jobTitle, industry, company)
	urlPath := fmt.Sprintf("https://newsapi.org/v2/everything?%s&apiKey=%s&language=en&pageSize=100&sortBy=publishedAt", query, newsAPIKey)
	resp, err := http.Get(urlPath)
	if err != nil {
		return "Error fetching latest trends"
	}
	defer resp.Body.Close()
	body, _ := io.ReadAll(resp.Body)
	var result struct {
		Articles []struct {
			Author      string    `json:"author"`
			Title       string    `json:"title"`
			Description string    `json:"description"`
			URL         string    `json:"url"`
			URLToImage  string    `json:"urlToImage"`
			PublishedAt time.Time `json:"publishedAt"`
			Content     string    `json:"content"`
		} `json:"articles"`
	}
	json.Unmarshal(body, &result)
	if len(result.Articles) == 0 {
		return "No relevant trends found."
	}
	trends := ""
	for _, item := range result.Articles {
		trends += fmt.Sprintf("- %s\n", item.Title)
	}
	return trends
}
