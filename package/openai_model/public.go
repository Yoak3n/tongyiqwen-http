package openai_model

type Message struct {
	Role    string `json:"role"` //system,user or assistant
	Content string `json:"content"`
}

type Parameters struct {
	ResultFormat      string   `json:"ResultFormat,omitempty"`
	TopK              int      `json:"TopK,omitempty"`
	Seed              int      `json:"Seed,omitempty"`
	Temperature       float64  `json:"Temperature,omitempty"`
	MaxNewTokens      int      `json:"MaxNewTokens,omitempty"`
	Stop              []string `json:"Stop,omitempty"`
	IncrementalOutput bool     `json:"IncrementalOutput,omitempty"`
}
type RequestBody struct {
	Id         string      `json:"id"`
	Model      string      `json:"model"`
	Stream     bool        `json:"stream"`
	Messages   []Message   `json:"messages"`
	Prompt     string      `json:"prompt"`
	TopP       float64     `json:"topP,omitempty"`
	Parameters *Parameters `json:"parameters,omitempty"`
}

type Choice struct {
	Index        int     `json:"index"`
	FinishReason string  `json:"finish_reason"`
	Message      Message `json:"message"`
}

type Usage struct {
	CompletionTokens int `json:"completion_tokens"`
	PromptTokens     int `json:"prompt_tokens"`
	TotalTokens      int `json:"total_tokens"`
}
type Result struct {
	Type string `json:"type"`
}

type ResponseBody struct {
	FinishReason string    `json:"finish_reason"`
	Text         string    `json:"text"`
	Content      []Result  `json:"content"`
	Object       string    `json:"object"`
	Id           string    `json:"id"`
	Created      int64     `json:"created"`
	Model        string    `json:"model"`
	Choices      []*Choice `json:"choices"`
	*Usage       `json:"usage"`
}
