package websearch

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"questions-generators/internal/config"
	"questions-generators/internal/utils"
	"time"
)

// FetchLatestTrends fetches latest trends from Google Search API
func FetchLatestTrends(jobTitle, industry, company string) string {
	config.LoadEnv()
	googleAPIKey := config.GetEnv("GOOGLE_API_KEY", "")
	googleCX := config.GetEnv("GOOGLE_CX", "")

	currentYear := time.Now().Year()

	queryTemplate, err := utils.ReadFile("prompts/websearch_query.txt")
	if err != nil {
		return "Error: Could not read Web Search query file."
	}

	query := utils.ReplacePlaceholders(queryTemplate, map[string]string{
		"{jobTitle}": jobTitle,
		"{industry}": industry,
		"{company}":  company,
		"{year}":     fmt.Sprintf("%d", currentYear),
	})
	encodedQuery := url.QueryEscape(query)
	urlPath := fmt.Sprintf("https://www.googleapis.com/customsearch/v1?q=%s&key=%s&cx=%s", encodedQuery, googleAPIKey, googleCX)

	resp, err := http.Get(urlPath)
	if err != nil {
		return "Error fetching latest trends."
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)

	var result struct {
		Items []struct {
			Title   string `json:"title"`
			Snippet string `json:"snippet"`
		} `json:"items"`
	}
	json.Unmarshal(body, &result)

	if len(result.Items) == 0 {
		return "No relevant trends found."
	}

	trends := ""
	for _, item := range result.Items {
		trends += fmt.Sprintf("- %s: %s\n", item.Title, item.Snippet)
	}

	return trends
}
