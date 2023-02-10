package chatgpt

// OpenaiBaseUrl base url
const OpenaiBaseUrl = "https://api.openai.com/v1"

const (
	// Gpt3TextDavinci003
	// Most capable GPT-3 model. Can do any task the other models can do,
	// often with higher quality, longer output and better instruction-following.
	// Also supports inserting completions within text.
	Gpt3TextDavinci003 = "text-davinci-003"

	// Gpt3TextCurie001
	// Very capable, but faster and lower cost than Davinci.
	Gpt3TextCurie001 = "text-curie-001"

	// Gpt3TextBabbage001
	// Capable of straightforward tasks, very fast, and lower cost.
	Gpt3TextBabbage001 = "text-babbage-001"

	// Gpt3TextAda001
	// Capable of very simple tasks, usually the fastest model in the GPT-3 series, and lowest cost.
	Gpt3TextAda001 = "text-ada-001"

	// CodexCodeDavinci002
	// Most capable Codex model. Particularly good at translating natural language to code.
	// In addition to completing code, also supports inserting completions within code.
	CodexCodeDavinci002 = "code-davinci-002"

	// CodexCodeCushman001
	// Almost as capable as Davinci Codex, but slightly faster.
	// This speed advantage may make it preferable for real-time applications.
	CodexCodeCushman001 = "code-cushman-001"

	Gpt3TextDavinciEdit001  = "text-davinci-edit-001"
	CodexCodeDavinciEdit001 = "code-davinci-edit-001"
)
