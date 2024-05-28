package api

import (
	"tongyiqwen/package/openai_model"
)

type RequestBody struct {
	Id      string `json:"id"`
	Preset  string `json:"preset,omitempty"`
	Content string `json:"content"`
}

type UploadPreset struct {
	Name string                 `json:"name"`
	Type string                 `json:"type"`
	Text string                 `json:"text,omitempty"`
	Map  []openai_model.Message `json:"map,omitempty"`
}
