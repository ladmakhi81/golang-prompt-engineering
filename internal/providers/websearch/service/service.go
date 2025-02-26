package websearchservice

type WebSearchServiceProvider interface {
	FetchLatestTrends(query string) string
}
