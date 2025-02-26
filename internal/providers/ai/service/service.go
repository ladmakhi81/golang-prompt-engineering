package aiservice

type AiServiceProvider interface {
	GenerateInterviewQuestions(model, prompt string) []string
}
