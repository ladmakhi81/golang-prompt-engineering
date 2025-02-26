package websearchservice

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"questions-generators/internal/config"
)

type GoogleWebSearch struct{}

func NewGoogleWebSearch() GoogleWebSearch {
	return GoogleWebSearch{}
}

func (svc GoogleWebSearch) FetchLatestTrends(query string) string {
	config.LoadEnv()
	googleAPIKey := config.GetEnv("GOOGLE_API_KEY", "")
	googleCX := config.GetEnv("GOOGLE_CX", "")
	urlPath := fmt.Sprintf("https://www.googleapis.com/customsearch/v1?q=%s&key=%s&cx=%s", query, googleAPIKey, googleCX)
	resp, err := http.Get(urlPath)
	if err != nil {
		return "Error fetching latest trends."
	}
	defer resp.Body.Close()
	body, _ := io.ReadAll(resp.Body)
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
