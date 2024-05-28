package qianwen

import (
	"tongyiqwen/package/openai_model"
)

type Parameters struct {
	ResultFormat      string   `json:"result_format,omitempty"`
	TopK              int      `json:"topK,omitempty"`
	Seed              int      `json:"seed,omitempty"`
	Temperature       float64  `json:"temperature,omitempty"`
	MaxNewTokens      int      `json:"max_newTokens,omitempty"`
	Stop              []string `json:"stop,omitempty"`
	IncrementalOutput bool     `json:"incremental_output,omitempty"`
}
type RequestBody struct {
	RequestId  string      `json:"requestId"`
	Stream     bool        `json:"stream"`
	Model      string      `json:"model"`
	Input      Input       `json:"input"`
	TopP       float64     `json:"topP,omitempty"`
	Parameters *Parameters `json:"parameters,omitempty"`
}

type Input struct {
	Messages []openai_model.Message `json:"messages"`
}

type Choice struct {
	Index        int
	Content      string
	FinishReason string
	Role         string
}

type Usage struct {
	InputTokens  int
	OutputTokens int
}
type ResponseBody struct {
	Success   bool
	Code      int
	Message   string
	RequestId string
	Data      struct {
		ResponseId string
		Text       string
		Choices    []Choice
		Usage      []Usage
	}
}
