package main

import (
	"fmt"
	"github.com/kainhuck/chatgpt"
)

func main() {
	gpt := chatgpt.NewChatGPT()
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
