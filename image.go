package chatgpt

// https://platform.openai.com/docs/api-reference/images

type ImageSize string

const (
	Small  ImageSize = "256x256"
	Middle ImageSize = "512x512"
	Big    ImageSize = "1024x1024"
)

type ResponseFormat string

const (
	FormatURL     = "url"
	FormatB64JSON = "b64_json"
)

type ImageRequest struct {
	Image          string         `json:"image,omitempty"`
	Mask           string         `json:"mask,omitempty"`
	Prompt         string         `json:"prompt,omitempty"`
	N              int            `json:"n,omitempty"` // 1-10
	Size           ImageSize      `json:"size,omitempty"`
	ResponseFormat ResponseFormat `json:"response_format,omitempty"`
	User           string         `json:"user,omitempty"`
}

type ImageResponse struct {
	Created int `json:"created"`
	Data    []struct {
		URL     string `json:"url"`
		B64Json string `json:"b64_json"`
	} `json:"data"`
}
