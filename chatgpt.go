package chatgpt

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"os"
	"strings"
)

type ChatGPT struct {
	apiKey       string
	organization string
	client       http.Client
}

func (c *ChatGPT) Init() {
	c.apiKey = os.Getenv("OPENAI_API_KEY")
	c.client = http.Client{}
}

func NewChatGPT(opts ...func(*ChatGPT)) *ChatGPT {
	c := &ChatGPT{}
	c.Init()

	for _, opt := range opts {
		opt(c)
	}

	return c
}

func SetApiKey(apiKey string) func(*ChatGPT) {
	return func(c *ChatGPT) {
		c.apiKey = apiKey
	}
}

func SetOrganization(organization string) func(*ChatGPT) {
	return func(c *ChatGPT) {
		c.organization = organization
	}
}

func (c *ChatGPT) sendRequest(apiEndpoint string, req interface{}, resp interface{}) error {
	if c.apiKey == "" {
		return fmt.Errorf("please set openapi apikey first, click `https://platform.openai.com/account/api-keys`")
	}
	baseEndpoint, err := url.Parse(OpenaiBaseUrl)
	if err != nil {
		return err
	}

	if !strings.HasSuffix(baseEndpoint.Path, "/") {
		baseEndpoint.Path += "/"
	}

	body, err := json.Marshal(req)
	if err != nil {
		return err
	}

	request, err := http.NewRequest(http.MethodPost, baseEndpoint.String()+apiEndpoint, bytes.NewReader(body))
	if err != nil {
		return err
	}

	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.apiKey))
	if c.organization != "" {
		request.Header.Set("OpenAI-Organization", c.organization)
	}

	response, err := c.client.Do(request)
	if err != nil {
		return err
	}
	defer func() {
		_ = response.Body.Close()
	}()

	return json.NewDecoder(response.Body).Decode(&resp)
}

func (c *ChatGPT) CreateCompletion(req CompletionRequest) (resp CompletionResponse, err error) {
	err = c.sendRequest("completions", req, &resp)

	return
}

func (c *ChatGPT) Reply(prompt string) string {
	resp, err := c.CreateCompletion(CompletionRequest{
		Model:       Gpt3TextDavinci003,
		Prompt:      prompt,
		Temperature: 0.6,
		MaxTokens:   128,
	})
	if err != nil {
		return fmt.Sprintf("Oops! got a error: (%v).", err)
	}
	if len(resp.Choices) > 0 {
		return resp.Choices[0].RealText()
	}
	return "Oops! got a mistake."
}
