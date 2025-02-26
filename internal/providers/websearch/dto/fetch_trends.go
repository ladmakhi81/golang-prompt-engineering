package websearchdto

type FetchTrendsDTO struct {
	JobTitle string
	Industry string
	Company  string
}

func NewFetchTrendsDTO(jobTitle, industry, company string) FetchTrendsDTO {
	return FetchTrendsDTO{
		JobTitle: jobTitle,
		Industry: industry,
		Company:  company,
	}
}
