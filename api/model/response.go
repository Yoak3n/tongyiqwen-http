package api

import "tongyiqwen/package/model"

type RequestBody struct {
	Id      string `json:"id"`
	Preset  string `json:"preset"`
	Content string `json:"content"`
}

type UploadPreset struct {
	Name    string          `json:"name"`
	Type    string          `json:"type"`
	Content string          `json:"content,omitempty"`
	Map     []model.Message `json:"map,omitempty"`
}
