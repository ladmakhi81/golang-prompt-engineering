package newssearchservice

type NewsSearchServiceProvider interface {
	FetchNews(query string) string
}
