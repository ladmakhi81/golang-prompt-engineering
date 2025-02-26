package newssearchdto

type FetchNewsDTO struct {
	JobTitle string
	Industry string
	Company  string
}

func NewFetchNewsDTO(jobTitle, industry, company string) FetchNewsDTO {
	return FetchNewsDTO{
		JobTitle: jobTitle,
		Industry: industry,
		Company:  company,
	}
}
