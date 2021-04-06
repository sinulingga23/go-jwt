package model

import (
	"log"
	"errors"

	database "github.com/sinulingga23/go-jwt/db"
)

type Category struct {
	CategoryId	int	`json:"categoryId"`
	Category	string	`json:"category"`
	Audit		Audit	`json:"audit"`
}


func (c *Category) IsCategoryExistByCategoryId(categoryId int) (bool, error) {
	db, err := database.ConnectDB()
	if err != nil {
		return false, err
	}
	defer db.Close()

	var check int64 = 0
	err = db.QueryRow("SELECT COUNT(category_id) FROM category WHERE category_id = ?", categoryId).Scan(&check)
	if err != nil {
		return false, err
	}

	if check != 1 {
		return false, errors.New("Somethings wrong!")
	}

	return true, nil
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

func (c *Category) FindCategoryByCategoryId(categoryId int) (Category, error) {
	db, err := database.ConnectDB()
	if err != nil {
		return Category{}, err
	}
	defer db.Close()

	var category Category
	err = db.QueryRow("SELECT category_id, category, created_at, updated_at FROM category WHERE category_id = ?", categoryId).
		Scan(&category.CategoryId, &category.Category, &category.Audit.CreatedAt, &category.Audit.UpdatedAt)
	if err != nil {
		return Category{}, err
	}

	if category == (Category{}) {
		return Category{}, errors.New("Category can't found")
	}

	return category, nil
}

func (c *Category) UpdateCategoryByCategoryId(categoryId int) (Category, error) {
	db, err := database.ConnectDB()
	if err != nil {
		return Category{}, err
	}
	defer db.Close()

	result, err := db.Exec("UPDATE category SET category_id = ?, category = ?, created_at = ?, updated_at = ? WHERE category_id = ?",
		c.CategoryId,
		c.Category,
		c.Audit.CreatedAt,
		c.Audit.UpdatedAt,
		categoryId)
	if err != nil {
		return Category{}, err
	}

	var rowsAffected int64
	if rowsAffected, err = result.RowsAffected(); err != nil {
		return Category{}, err
	}

	if rowsAffected == 0 {
		return Category{}, errors.New("Maybe the category is not exist.")
	}

	if rowsAffected != 1 {
		return Category{}, errors.New("Somethings wrong!")
	}

	var category Category
	category = Category{
			CategoryId: c.CategoryId,
			Category: c.Category,
			Audit: Audit {CreatedAt: c.Audit.CreatedAt, UpdatedAt: c.Audit.UpdatedAt},}
	return category, nil
}

func (c *Category) DeleteCategoryByCategoryId(categoryId int) (bool, error) {
	db, err := database.ConnectDB()
	if err != nil {
		return false, err
	}
	defer db.Close()

	result, err := db.Exec("DELETE FROM category WHERE category_id = ?", categoryId)
	if err != nil {
		return false, err
	}

	var rowsAffected int64
	if rowsAffected, err = result.RowsAffected(); err != nil {
		return false, err
	}

	if rowsAffected == 0 {
		return false, errors.New("Maybe the category is not exist.")
	}

	if rowsAffected != 1 {
		return false, errors.New("Somethings wrong")
	}

	return true, nil
}

func (c *Category) FindProductsByCategoryId(categoryId int) ([]Product, error) {
	db, err := database.ConnectDB()
	if err != nil {
		return []Product{}, err
	}
	defer db.Close()

	rows, err := db.Query("SELECT p.product_id, p.category_id, p.name, p.unit, p.price, p.stock, p.created_at, p.updated_at FROM product p INNER JOIN category c ON p.product_id = c.category_id HAVING p.category_id = ?", categoryId)
	if err != nil {
		return []Product{}, err
	}
	defer rows.Close()

	var result []Product
	for rows.Next() {
		var each Product
		err = rows.Scan(&each.ProductId, &each.CategoryId, &each.Name, &each.Unit, &each.Price, &each.Stock, &each.Audit.CreatedAt, &each.Audit.UpdatedAt)
		if err != nil {
			return []Product{}, err
		}

		result = append(result, each)
	}

	if err = rows.Err(); err != nil {
		return []Product{}, err
	}

	return result, nil
}
