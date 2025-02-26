package main

import (
	"fmt"
	"log"
	"net/http"
	"questions-generators/internal/handlers"
	"questions-generators/internal/providers/ai"
	aiservice "questions-generators/internal/providers/ai/service"
	"questions-generators/internal/providers/newssearch"
	newssearchservice "questions-generators/internal/providers/newssearch/service"
	"questions-generators/internal/providers/websearch"
	websearchservice "questions-generators/internal/providers/websearch/service"
	questionsvc "questions-generators/internal/services/question/v1"
)

func main() {
	googleWebsearchProvider := websearchservice.NewGoogleWebSearch()
	websearchSvc := websearch.NewWebSearchService(
		googleWebsearchProvider,
	)

	newsapiSearchProvider := newssearchservice.NewNewsApiSearch()
	newsSearchSvc := newssearch.NewNewsSearchService(
		newsapiSearchProvider,
	)

	avalaiProvider := aiservice.NewAvalaiService()
	aiSvc := ai.NewAiService(
		avalaiProvider,
	)

	questionSvc := questionsvc.NewQuestionService(websearchSvc, newsSearchSvc, aiSvc)
	questionHandler := handlers.NewQuestionHandler(questionSvc)

	http.HandleFunc("/questions", questionHandler.GetInterviewQuestions)

	fmt.Println("Server running on port 8082...")
	log.Fatal(http.ListenAndServe(":8082", nil))
}
