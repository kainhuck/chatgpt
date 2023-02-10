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

	if c.apiKey == "" {
		fmt.Println("please set openapi apikey first")
		fmt.Println("set env `OPENAI_API_KEY=<your apikey>` or chatgpt.NewChatGPT(chatgpt.NewChatGPT(chatgpt.SetApiKey(\"your apikey\")))")
		fmt.Println("you can get your apikey from `https://platform.openai.com/account/api-keys`")
		os.Exit(1)
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

func (c *ChatGPT) newUrl(apiEndpoint string) (string, error) {
	baseEndpoint, err := url.Parse(OpenaiBaseUrl)
	if err != nil {
		return "", err
	}

	if !strings.HasSuffix(baseEndpoint.Path, "/") {
		baseEndpoint.Path += "/"
	}

	return baseEndpoint.String() + apiEndpoint, nil
}

func (c *ChatGPT) setBaseRequestHeader(request *http.Request) {
	request.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.apiKey))
	if c.organization != "" {
		request.Header.Set("OpenAI-Organization", c.organization)
	}
}

func (c *ChatGPT) sendRequest(request *http.Request, resp interface{}) error {
	c.setBaseRequestHeader(request)
	response, err := c.client.Do(request)
	if err != nil {
		return err
	}
	defer func() {
		_ = response.Body.Close()
	}()

	return json.NewDecoder(response.Body).Decode(&resp)
}

func (c *ChatGPT) sendPostRequest(apiEndpoint string, req interface{}, resp interface{}) error {
	u, err := c.newUrl(apiEndpoint)
	if err != nil {
		return err
	}

	body, err := json.Marshal(req)
	if err != nil {
		return err
	}

	request, err := http.NewRequest(http.MethodPost, u, bytes.NewReader(body))
	if err != nil {
		return err
	}
	request.Header.Set("Content-Type", "application/json")

	return c.sendRequest(request, resp)
}

func (c *ChatGPT) sendGetRequest(apiEndpoint string, resp interface{}) error {
	u, err := c.newUrl(apiEndpoint)
	if err != nil {
		return err
	}

	request, err := http.NewRequest(http.MethodGet, u, nil)
	if err != nil {
		return err
	}

	return c.sendRequest(request, resp)
}

func (c *ChatGPT) CreateCompletion(req CompletionRequest) (resp CompletionResponse, err error) {
	err = c.sendPostRequest("completions", req, &resp)

	return
}

func (c *ChatGPT) ListModels() (resp ModelResponse, err error) {
	err = c.sendGetRequest("models", &resp)

	return
}

func (c *ChatGPT) RetrieveModel(modelName string) (model Model, err error) {
	err = c.sendGetRequest("models/"+modelName, &model)

	return
}

func (c *ChatGPT) CreateEdit(req EditRequest) (resp EditResponse, err error) {
	err = c.sendPostRequest("edits", req, &resp)

	return
}

func (c *ChatGPT) CreateImage(req ImageRequest) (resp ImageResponse, err error) {
	err = c.sendPostRequest("images/generations", req, &resp)

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
