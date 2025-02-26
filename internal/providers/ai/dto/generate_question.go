package aidto

type GenerateInterviewQuestionDTO struct {
	JobTitle         string
	Industry         string
	Company          string
	AiModel          string
	JobDescription   string
	Trends           string
	News             string
	Cv               string
	IncludeCv        bool
	IncludeWebSearch bool
	IncludeNews      bool
}

func NewGenerateInterviewQuestionDTO(
	jobTitle string,
	industry string,
	company string,
	aiModel string,
	jobDescription string,
	trends string,
	news string,
	cv string,
	includeCv bool,
	includeWebSearch bool,
	includeNews bool,
) GenerateInterviewQuestionDTO {
	return GenerateInterviewQuestionDTO{
		JobTitle:         jobTitle,
		Industry:         industry,
		Company:          company,
		AiModel:          aiModel,
		JobDescription:   jobDescription,
		Trends:           trends,
		News:             news,
		Cv:               cv,
		IncludeCv:        includeCv,
		IncludeWebSearch: includeWebSearch,
		IncludeNews:      includeNews,
	}
}
