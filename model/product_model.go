package model

import (
	"errros"
)

type Product struct {
	ProductId	int 	`json:"productId"`
	CategoryId	string  `json:"categoryProductId"`
	Name		string  `json:"product"`
	Unit		string  `json:"unit"`
	Price		float64 `json:"price"`
	Stock           int     `json:"stock"`
	Audit           Audit   `json:"audit"`
}

func (p *Product) IsProductExistByProductId(productId int) (bool, error) {
	db, err := database.ConnectDB()
	if err != nil {
		return false, err
	}
	defer db.Close()

	var check int64 = 0
	err = db.QueryRow("SELECT COUNT(product_id) FROM product WHERE product_id = ?", productId).Scan(&check)
	if err != nil {
		return false, err
	}

	if check != 1 {
		return false, errors.New("Somethings wrong!")
	}

	return true, nil
}

func (p *Product) SaveProduct() (Product, error) {
	db, err := database.ConnectDB()
	if err != nil {
		return Product{}, err
	}
	defer db.Close()

	result, err := db.Exec("INSERT INTO product (category_id, name, unit, price, stock, created_at) VALUES (?, ?, ?, ?, ?, ?)",
			p.CategoryId,
			p.Name,
			p.Unit,
			p.Price,
			p.Stock,
			p.Audit.CreatedAt)
	if err != nil {
		return Product{}, err
	}

	var lastInsertId int64
	if lastInsertId, err = result.LastInsertId(); err != nil {
		return Product{}, err
	}
	c.ProductId = int(lastInsertId)

	var product Product
	product = Product{
			ProductId: p.ProductId,
			CategoryId: p.CategoryId,
			Name: p.Name,
			Unit: p.Unit,
			Price: p.price,
			Stock: p.Stock,
			Audit: Audit {CreatedAt: p.Audit.CreatedAt},
		}
	return product, nil
}

func (p *Product) FindProductByProductId(productId int) (Product, error) {
	db, err := database.ConnectDB()
	if err != nil {
		return Product{}, err
	}
	defer db.Close()

	var product Product
	err = db.QueryRow("SELECT product_id, category_id, name, unit, price, stock, created_at, updated_at FROM product WHERE product_id = ?", productId).
		Scan(&product.ProductId, &product.CategoryId, &product.Name, &product.Unit, &product.Price, &product.Stock, &product.Audit.CreatedAt, &product.Audit.UpdatedAt)
	if err != nil {
		return Product{}, err
	}

	if Product == (Product{}) {
		return Product{}, errors.New("Product can't found")
	}

	return Product, nil
}
