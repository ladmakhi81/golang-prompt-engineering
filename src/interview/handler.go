package interview

import (
	"fmt"
	newsservice "interview-generator/src/news/service"
	promptservice "interview-generator/src/prompt/service"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

type InterviewHandler struct {
	newsSvc   newsservice.NewsService
	promptSvc promptservice.PromptService
}

func NewInterview(
	newsSvc newsservice.NewsService,
	promptSvc promptservice.PromptService,
) InterviewHandler {
	return InterviewHandler{
		newsSvc:   newsSvc,
		promptSvc: promptSvc,
	}
}

func (svc InterviewHandler) GenerateInterviewQuestion(ctx *gin.Context) {
	questionTopic := ctx.Query("topic")
	questionIndustry := ctx.Query("industry")
	questionPlatform := ctx.Query("platform")
	questionJobDesc := ctx.Query("jobDesc")

	isProvideParameter := questionTopic != "" && questionIndustry != "" && questionPlatform != "" && questionJobDesc != ""

	if !isProvideParameter {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "You must provide topic - industry - platform as important parameter"})

		return
	}

	news, newsErr := svc.newsSvc.FetchLatestNews(questionTopic)
	if newsErr != nil {
		fmt.Println(newsErr)
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "something went wrong"})

		return
	}

	newsTopics := svc.newsSvc.GetTitlesOfArticles(news.Articles)
	promptToken := map[string]any{
		"{topic}":           questionTopic,
		"{industry}":        questionIndustry,
		"{platform}":        questionPlatform,
		"{articles}":        strings.Join(newsTopics, "\n"),
		"{job_description}": questionJobDesc,
	}

	prompt, promptErr := svc.promptSvc.GetPrompt("prompt-with-industry-platform", promptToken)
	if promptErr != nil {
		fmt.Println(promptErr)
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "something went wrong"})
		return
	}

	promptRes, promptResErr := svc.promptSvc.SendPromptToChatgpt(prompt)
	if promptResErr != nil {
		fmt.Println(promptResErr)
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "something went wrong"})
		return
	}
	// if promptRes.Choices == nil {
	// 	ctx.JSON(http.StatusOK, gin.H{"message": "can't generate interview question based on provided information"})
	// 	return
	// }

	// choice := promptRes.Choices[0]
	ctx.JSON(http.StatusOK, gin.H{"questions": promptRes})
}
