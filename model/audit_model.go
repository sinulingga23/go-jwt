package model

import (
	"time"
)

type Audit struct {
	CreatedAt	time.Time	`json:"createdAt"`
	UpdatedAt	*time.Time	`json:"updatedAt"`
}
