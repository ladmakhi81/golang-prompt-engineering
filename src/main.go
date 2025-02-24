package main

import (
	"fmt"
	"interview-generator/src/interview"
	newsservice "interview-generator/src/news/service"
	promptservice "interview-generator/src/prompt/service"
	"os"

	"log"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func main() {
	rootDirectory, rootDirectoryErr := os.Getwd()
	if rootDirectoryErr != nil {
		panic("unable to find root directory")
	}
	viper.SetConfigFile(fmt.Sprintf("%s/.env", rootDirectory))

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file, %s", err)
	}

	server := gin.Default()
	port := 8080

	newsSvc := newsservice.NewNews()
	promptSvc := promptservice.NewPromptService()
	interviewSvc := interview.NewInterview(newsSvc, promptSvc)

	server.GET("/interview", interviewSvc.GenerateInterviewQuestion)

	log.Fatalln(server.Run(fmt.Sprintf(":%d", port)))
}
