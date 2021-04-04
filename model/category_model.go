package model

type Category struct {
	CategoryId	string	`json:"categoryId"`
	Category	string	`json:"category"`
	Audit		Audit	`json:"audit"`
}
