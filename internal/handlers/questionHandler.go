package handlers

import (
	"encoding/json"
	"net/http"
	v1 "questions-generators/internal/services/question/v1"
	"strings"
)

func QuestionHandler(w http.ResponseWriter, r *http.Request) {
	var body struct {
		JobTitle       string         `json:"jobTitle"`
		Industry       string         `json:"industry"`
		Company        string         `json:"company"`
		JobDescription string         `json:"jobDescription"`
		AiModel        string         `json:"model"`
		CV             map[string]any `json:"cv"`
		Features       string         `json:"features"`
	}

	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		http.Error(w, "Unable to parse request body", http.StatusBadRequest)
		return
	}

	if body.JobTitle == "" || body.Industry == "" || body.Company == "" || body.AiModel == "" {
		http.Error(w, "Missing required parameters", http.StatusBadRequest)
		return
	}

	transformedCV, cvErr := json.Marshal(body.CV)
	if cvErr != nil {
		http.Error(w, "Unable to parse cv", http.StatusBadRequest)
		return
	}

	featuresSegments := strings.Split(body.Features, ",")

	service := v1.NewQuestionService()
	questions := service.GetQuestions(body.JobTitle, body.Industry, body.Company, body.JobDescription, body.AiModel, string(transformedCV), featuresSegments)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(questions)
}
