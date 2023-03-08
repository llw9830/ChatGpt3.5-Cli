package services

import (
	"context"
	"fmt"
	"log"

	gogpt "github.com/sashabaranov/go-gpt3"
)

var client *gogpt.Client
var request gogpt.ChatCompletionRequest

func InitClient(api, system string) {
	gptConfig := gogpt.DefaultConfig(api)
	client = gogpt.NewClientWithConfig(gptConfig)

	request.Messages = append(request.Messages, gogpt.ChatCompletionMessage{
		Role: "system", Content: system,
	})
	request.Model = "gpt-3.5-turbo-0301"
}

func GetAnswer(question string) (reply string, ok bool) {
	fmt.Print("Bot: ")
	request.Messages = append(request.Messages, gogpt.ChatCompletionMessage{
		Role: "user", Content: question,
	})
	ctx := context.Background()
	resp, err := client.CreateChatCompletion(ctx, request)
	if err != nil {
		log.Fatalln(err)
	}

	reply = resp.Choices[0].Message.Content
	if reply != "" {
		request.Messages = append(request.Messages, gogpt.ChatCompletionMessage{
			Role: "assistant", Content: reply,
		})

		for _, v := range reply {
			fmt.Print(string(v))
		}
		fmt.Println()
		ok = true
	}

	return reply, ok
}

func FormatQuestion(question string) string {
	return "Answer:" + question
}
