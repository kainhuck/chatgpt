package chatgpt

import "strings"

type CompletionRequest struct {
	Model       string  `json:"model"`
	Prompt      string  `json:"prompt"`
	Temperature float32 `json:"temperature"`
	MaxTokens   int     `json:"max_tokens"`
}

type Choice struct {
	Text         string `json:"text"`
	Index        int    `json:"index"`
	FinishReason string `json:"finish_reason"`
}

func (c Choice) RealText() string {
	return strings.TrimPrefix(c.Text, "\n\n")
}

type Usage struct {
	PromptTokens     int `json:"prompt_tokens"`
	CompletionTokens int `json:"completion_tokens"`
	TotalTokens      int `json:"total_tokens"`
}

type CompletionResponse struct {
	ID      string   `json:"id"`
	Object  string   `json:"object"`
	Created int      `json:"created"`
	Model   string   `json:"model"`
	Choices []Choice `json:"choices"`
	Usage   Usage    `json:"usage"`
}
