package models

type WSMessage struct {
	Type     string `json:"type"`
	Code     string `json:"code,omitempty"`
	Language string `json:"language,omitempty"`
	Count    int    `json:"count,omitempty"`
}
