package models

import "time"

type Room struct {
	ID int

	RoomID string

	Code string

	Language string

	UpdatedAt time.Time
}
