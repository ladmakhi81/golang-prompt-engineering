package main

import (
	"fmt"
	"log"
	"net/http"
	"questions-generators/internal/handlers"
)

func main() {
	http.HandleFunc("/questions", handlers.QuestionHandler)

	fmt.Println("Server running on port 8082...")
	log.Fatal(http.ListenAndServe(":8082", nil))
}
