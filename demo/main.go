package main

import (
	"fmt"
	"github.com/kainhuck/chatgpt"
)

func main() {
	gpt := chatgpt.NewChatGPT()
	//fmt.Println(gpt.Reply("Write a tagline for an ice cream shop. in chinese"))
	models, err := gpt.ListModels()
	if err != nil {
		panic(err)
	}

	fmt.Println(models)

	m, err := gpt.RetrieveModel(chatgpt.Gpt3TextDavinci003)
	if err != nil {
		panic(err)
	}
	fmt.Println(m)
}
