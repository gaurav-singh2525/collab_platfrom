package dto

type ExecuteRequest struct {
	RoomID   string `json:"roomId"`
	Language string `json:"language"`
	Code     string `json:"code"`
}
