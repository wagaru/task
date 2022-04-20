package model

type Task struct {
	ID     uint32 `json:"id"`
	Name   string `json:"name"`
	Status uint8  `json:"status"`
}
