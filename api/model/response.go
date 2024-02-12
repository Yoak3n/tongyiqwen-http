package api

type RequestBody struct {
	Id      string `json:"id"`
	Preset  string `json:"preset"`
	Content string `json:"content"`
}
