package ali_model

type Message struct {
	Role    string //system,user or assistant
	Content string
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
	RequestId  string `json:"RequestId"`
	AppId      string `json:"AppId"`
	Stream     bool   `json:"Stream"`
	Messages   []Message
	TopP       float64     `json:"TopP,omitempty"`
	Parameters *Parameters `json:"Parameters,omitempty"`
}

type Choice struct {
	Index        int
	FinishReason string
	Message      *Message
}

type Usage struct {
	InputTokens  int
	OutputTokens int
	ModelId      string
}
type ResponseBody struct {
	Success   bool
	Code      int
	Message   string
	RequestId string
	Data      struct {
		ResponseId string
		SessionId  string
		Text       string
		Choices    []Choice
		Usage      []Usage
	}
}
