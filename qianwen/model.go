package qianwen

import "tongyiqwen/package/model"

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
	RequestId  string `json:"RequestId"`
	AppId      string `json:"AppId"`
	Stream     bool   `json:"Stream"`
	Messages   []model.Message
	TopP       float64     `json:"TopP,omitempty"`
	Parameters *Parameters `json:"Parameters,omitempty"`
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
