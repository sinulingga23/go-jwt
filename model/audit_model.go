package model

type Audit struct {
	CreatedAt	string		`json:"createdAt"`
	UpdatedAt	*string		`json:"updatedAt"`
}

