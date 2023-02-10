package main

import (
	"fmt"
	"github.com/kainhuck/chatgpt"
)

func main() {
	gpt := chatgpt.NewChatGPT()
	ModelDemo(gpt)
}

func ModelDemo(gpt *chatgpt.ChatGPT) {
	resp, err := gpt.ListModels()
	if err != nil {
		panic(err)
	}

	for _, model := range resp.Data {
		fmt.Println(model.ID)
	}
}

func CompletionDemo(gpt *chatgpt.ChatGPT) {
	resp, err := gpt.CreateCompletion(chatgpt.CompletionRequest{
		Model:     chatgpt.Gpt3TextDavinci003,
		Prompt:    "路亚翘嘴怎么钓？",
		MaxTokens: 512,
	})
	if err != nil {
		panic(err)
	}
	fmt.Println(resp.Choices[0].RealText())
}

func EditDemo(gpt *chatgpt.ChatGPT) {
	resp, err := gpt.CreateEdit(chatgpt.EditRequest{
		Model:       chatgpt.Gpt3TextDavinciEdit001,
		Input:       "What day of the wek is it?",
		Instruction: "Fix the spelling mistakes",
	})
	if err != nil {
		panic(err)
	}
	fmt.Println(resp.Choices[0].RealText())
}

func ImageDemo(gpt *chatgpt.ChatGPT) {
	resp, err := gpt.CreateImage(chatgpt.ImageRequest{
		Prompt: "夕阳西下断肠人在天涯",
	})
	if err != nil {
		panic(err)
	}

	for _, d := range resp.Data {
		fmt.Println(d.URL)
	}
}
