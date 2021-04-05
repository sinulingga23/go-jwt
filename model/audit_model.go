package model

import (
	"time"
	"database/sql"
)

type Audit struct {
	CreatedAt	time.Time	`json:"createdAt"`
	UpdatedAt	sql.NullTime	`json:"updatedAt"`
}
