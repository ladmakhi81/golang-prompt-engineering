package question

type QuestionInterface interface {
	GetQuestions(jobTitle, industry, company, aiModel, cv string, features []string) []string
}
