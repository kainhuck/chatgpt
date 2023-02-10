package main

import (
	"fmt"
	"github.com/kainhuck/chatgpt"
)

func main() {
	gpt := chatgpt.NewChatGPT()
	ImageDemo(gpt)
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