package models

type WSMessage struct {
	Type string `json:"type"`
	Code string `json:"code,omitempty"`
}
