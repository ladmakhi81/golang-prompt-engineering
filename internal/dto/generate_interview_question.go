package dto

type GenerateInterviewQuestionReqBody struct {
	JobTitle       string         `json:"jobTitle"`
	Industry       string         `json:"industry"`
	Company        string         `json:"company"`
	JobDescription string         `json:"jobDescription"`
	AiModel        string         `json:"model"`
	CV             map[string]any `json:"cv"`
	Features       string         `json:"features"`
	CvAsText       string         `json:"-"`
}
