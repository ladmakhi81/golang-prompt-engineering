package main

import (
	"fmt"
	"interview-generator/src/interview"
	newsservice "interview-generator/src/news/service"
	promptservice "interview-generator/src/prompt/service"

	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()
	port := 8080

	newsSvc := newsservice.NewNews()
	promptSvc := promptservice.NewPromptService()
	interviewSvc := interview.NewInterview(newsSvc, promptSvc)

	server.POST("/interview", interviewSvc.GenerateInterviewQuestion)

	log.Fatalln(server.Run(fmt.Sprintf(":%d", port)))
}
