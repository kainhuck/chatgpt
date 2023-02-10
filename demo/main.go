package main

import (
	"fmt"
	"github.com/kainhuck/chatgpt"
)

func main() {
	gpt := chatgpt.NewChatGPT()
	fmt.Println(gpt.Reply("Write a tagline for an ice cream shop. in chinese"))
}
