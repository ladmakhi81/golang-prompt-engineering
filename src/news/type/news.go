package newstype

import "time"

type NewsSource struct {
	Name string `json:"name"`
}

type NewsArticle struct {
	Source      NewsSource `json:"source"`
	Author      string     `json:"author"`
	Title       string     `json:"title"`
	Description string     `json:"description"`
	URL         string     `json:"url"`
	URLToImage  string     `json:"urlToImage"`
	PublishedAt time.Time  `json:"publishedAt"`
	Content     string     `json:"content"`
}

type NewsResponse struct {
	Status       string        `json:"status"`
	TotalResults uint          `json:"totalResults"`
	Articles     []NewsArticle `json:"articles"`
}
