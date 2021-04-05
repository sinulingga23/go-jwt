package model

import (
	"log"
	database "github.com/sinulingga23/go-jwt/db"
)

type Category struct {
	CategoryId	int	`json:"categoryId"`
	Category	string	`json:"category"`
	Audit		Audit	`json:"audit"`
}

func (c *Category) SaveCategory() (Category, error) {
	db, err := database.ConnectDB()
	if err != nil {
		log.Printf("%v\n", err)
		return Category{}, err
	}
	defer db.Close()

	result, err := db.Exec("INSERT INTO category (category, created_at) VALUES (?, ?)", c.Category, c.Audit.CreatedAt)
	if err != nil {
		log.Printf("%v\n", err)
		return Category{}, err
	}

	var lastInsertId int64
	if lastInsertId, err = result.LastInsertId(); err != nil {
		log.Printf("%v\n", err)
		return Category{}, err
	}
	c.CategoryId = int(lastInsertId)

	var category Category
	category = Category{CategoryId: c.CategoryId, Category: c.Category, Audit: Audit {CreatedAt: c.Audit.CreatedAt}}
	return category, nil
}
