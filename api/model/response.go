package api

import "tongyiqwen/package/ali_model"

type RequestBody struct {
	Id      string `json:"id"`
	Preset  string `json:"preset,omitempty"`
	Content string `json:"content"`
}

type UploadPreset struct {
	Name string              `json:"name"`
	Type string              `json:"type"`
	Text string              `json:"text,omitempty"`
	Map  []ali_model.Message `json:"map,omitempty"`
}
