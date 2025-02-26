package question

import "questions-generators/internal/dto"

type QuestionInterface interface {
	GetQuestions(dto dto.GenerateInterviewQuestionReqBody) []string
}
