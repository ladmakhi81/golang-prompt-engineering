package interviewdto

type GetInterviewMetadataDTO struct {
	Topic          string `json:"topic"`
	Platform       string `json:"platform"`
	Industry       string `json:"industry"`
	JobDescription string `json:"jobDescription"`
}

func NewGetInterviewMetadataDTO() *GetInterviewMetadataDTO {
	return new(GetInterviewMetadataDTO)
}
