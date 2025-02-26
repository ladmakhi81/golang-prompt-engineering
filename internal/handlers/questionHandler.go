package handlers

import (
	"encoding/json"
	"net/http"
	"questions-generators/internal/dto"
	"questions-generators/internal/providers/ai"
	aiservice "questions-generators/internal/providers/ai/service"
	"questions-generators/internal/providers/newssearch"
	newssearchservice "questions-generators/internal/providers/newssearch/service"
	"questions-generators/internal/providers/websearch"
	websearchservice "questions-generators/internal/providers/websearch/service"
	v1 "questions-generators/internal/services/question/v1"
)

type QuestionHandler struct {
	svc *v1.QuestionService
}

func NewQuestionHandler(
	svc *v1.QuestionService,
) QuestionHandler {
	return QuestionHandler{
		svc: svc,
	}
}

func (h QuestionHandler) GetInterviewQuestions(w http.ResponseWriter, r *http.Request) {
	var body dto.GenerateInterviewQuestionReqBody

	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		http.Error(w, "Unable to parse request body", http.StatusBadRequest)
		return
	}

	if body.JobTitle == "" || body.Industry == "" || body.Company == "" || body.AiModel == "" {
		http.Error(w, "Missing required parameters", http.StatusBadRequest)
		return
	}

	transformedCV, cvErr := json.Marshal(body.CV)
	body.CvAsText = string(transformedCV)
	if cvErr != nil {
		http.Error(w, "Unable to parse cv", http.StatusBadRequest)
		return
	}

	questions := h.svc.GetQuestions(body)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(questions)
}

func getSvc() *v1.QuestionService {
	websearchSvc := websearch.NewWebSearchService(
		websearchservice.NewGoogleWebSearch(),
	)

	newsSearchSvc := newssearch.NewNewsSearchService(
		newssearchservice.NewNewsApiSearch(),
	)

	aiSvc := ai.NewAiService(
		aiservice.NewAvalaiService(),
	)

	return v1.NewQuestionService(
		websearchSvc,
		newsSearchSvc,
		aiSvc,
	)
}
