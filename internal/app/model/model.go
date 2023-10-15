package model

import "time"

type Event struct {
	EventID   int64     `json:"eventID"`
	EventType string    `json:"eventType"`
	UserID    int64     `json:"userID"`
	EvenTime  time.Time `json:"evenTime"`
	Payload   string    `json:"payload"`
}
