package entities

import (
	"time"
)

type Log struct {
	ID          int       `json:"id"`
	ReferenceID int       `json:"reference_id"`
	Entity      string    `json:"entity"`
	Payload     string    `json:"payload"`
	CreatedAt   time.Time `json:"created_at"`
}
