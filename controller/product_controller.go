package controller

import (
	"fmt"
	"time"
	"strings"
	"strconv"
	"net/http"
	"encoding/json"

	"github.com/sinulingga23/go-jwt/model"
	"github.com/gorilla/mux"
)

var GetProducts = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var productModel model.Product
	var products []model.Product
	var err error
	if products, err = productModel.FindAllProduct(); err != nil {
		payload, _ := json.Marshal(struct {
			StatusCode	int 	`json:"statusCode"`
			Message 	string 	`json:"message"`
			Errors 		string	`json:"errors"`
		}{
			http.StatusBadRequest, "Somethings wrong!", fmt.Sprintf("%s", err),
		})
		w.Write([]byte(payload))
		return
	}

	payload, _ := json.Marshal(struct {
		StatusCode	int		`json:"statusCode"`
		Message 	string 		`json:"message"`
		Data 		[]model.Product	`json:"products"`
	}{
		http.StatusOK, "Success to get the products", products,
	})
	w.Write([]byte(payload))
	return
})

var CreateProduct = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var productRequest model.Product
	var categoryModel model.Category
	var err error

	if err = json.NewDecoder(r.Body).Decode(&productRequest); err != nil {
		payload, _ := json.Marshal(struct {
			StatusCode	int	`json:"statusCode"`
			Message		string	`json:"message"`
			Errors		string	`json:"errors"`
		}{
			http.StatusBadRequest, "invalid", fmt.Sprintf("%s", err),
		})
		w.Write([]byte(payload))
		return
	}

	if productRequest.CategoryId <= 0 {
		payload, _ := json.Marshal(struct {
			StatusCode	int	`json:"statusCode"`
			Message		string	`json:"message"`
			Errors 		string 	`json:"errors"`
		}{
			http.StatusBadRequest, "Product should have the category", "BadRequest",
		})
		w.Write([]byte(payload))
		return
	}

	if len(strings.Trim(productRequest.Name, " ")) == 0 {
		payload, _ := json.Marshal(struct {
			StatusCode	int	`json:"statusCode"`
			Message		string	`json:"message"`
			Errors		string	`json:"errors"`
		}{
			http.StatusBadRequest, "Product name can't be empty", "BadRequest",
		})
		w.Write([]byte(payload))
		return
	}

	if len(strings.Trim(productRequest.Unit, " ")) == 0 {
		payload, _ := json.Marshal(struct {
			StatusCode	int	`json:"statusCode"`
			Message		string	`json:"message"`
			Errors		string	`json:"errors"`
		}{
			http.StatusBadRequest, "Unit can't be empty", "BadRequest",
		})
		w.Write([]byte(payload))
		return
	}

	if productRequest.Stock < 0 {
		payload, _ := json.Marshal(struct {
			StatusCode	int	`json:"statusCode"`
			Message		string	`json:"message"`
			Errors		string	`json:"errors"`
		}{
			http.StatusBadRequest, "Stock can't be negative", "BadRequest",
		})
		w.Write([]byte(payload))
		return
	}

	if productRequest.AddSotck < 0 {
		payload, _ := json.Marshal(struct {
			StatusCode	int	`json:"statusCode"`
			Message		string	`json:"message"`
			Errors		string	`json:"errors"`
		}{
			http.StatusBadRequest, "Stock Add can't be negative", "BadRequest",
		})
		w.Write([]byte(payload))
		return
	}

	if productRequest.Price < 0 {
		payload, _ := json.Marshal(struct {
			StatusCode	int	`json:"statusCode"`
			Message		string	`json:"message"`
			Errors		string	`json:"errors"`
		}{
			http.StatusBadRequest, "Price can't be negative", "BadRequest",
		})
		w.Write([]byte(payload))
		return
	}

	var isExist bool = false
	if isExist, err = categoryModel.IsCategoryExistByCategoryId(productRequest.CategoryId); err != nil {
		payload, _ := json.Marshal(struct {
			StatusCode	int	`json:"statusCode"`
			Message		string	`json:"message"`
			Errors		string	`json:"errors"`
		}{
			http.StatusNotFound, "Category can't be found", fmt.Sprintf("%s", err),
		})
		w.Write([]byte(payload))
		return
	}

	if isExist {
		var createdProduct model.Product
		productRequest.Audit.CreatedAt = time.Now().Format("2006-01-02 15:05:03")
		if createdProduct, err = productRequest.SaveProduct(); err != nil {
			payload, _ := json.Marshal(struct {
				StatusCode	int	`json:"statusCode"`
				Message		string 	`json:"message"`
				Errors		string	`json:"errors"`
			}{
				http.StatusInternalServerError, "Somethings wrong!", fmt.Sprintf("%s", err),
			})
			w.Write([]byte(payload))
			return
		}

		payload, _ := json.Marshal(struct {
			StatusCode	int 		`json:"statusCode"`
			Message		string		`json:"message"`
			Data		model.Product	`json:"product"`
		}{
			http.StatusOK, "Success to create a new product", createdProduct,
		})
		w.Write([]byte(payload))
		return
	}

	payload, _ := json.Marshal(struct {
		StatusCode	int	`json:"statusCode"`
		Message		string	`json:"message"`
	}{
		http.StatusInternalServerError, "Somethings wrong!",
	})
	w.Write([]byte(payload))
	return
})

var GetProductByProductId = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	var productId int

	var err error
	if productId, err = strconv.Atoi(vars["productId"]); err != nil {
		payload, _ := json.Marshal(struct {
			StatusCode	int	`json:"statusCode"`
			Message		string	`json:"message"`
			Errors		string	`json:"errors"`
		}{
			http.StatusBadRequest, "Invalid request", "BadRequest",
		})
		w.Write([]byte(payload))
		return
	}

	var productModel model.Product
	if _, err = productModel.IsProductExistByProductId(productId); err != nil {
		payload, _ := json.Marshal(struct {
			StatusCode	int	`json:"statusCode"`
			Message		string	`json:"message"`
			Errors		string	`json:"errors"`
		}{
			http.StatusNotFound, "Product can't be found", fmt.Sprintf("%s", err),
		})
		w.Write([]byte(payload))
		return
	}

	var currentProduct model.Product
	if currentProduct, err = productModel.FindProductByProductId(productId); err != nil {
		payload, _ := json.Marshal(struct {
			StatusCode	int	`json:"statusCode"`
			Message		string	`json:"message"`
			Errors		string	`json:"errors"`
		}{
			http.StatusBadRequest, "Somethings wrong!", fmt.Sprintf("%s", err),
		})
		w.Write([]byte(payload))
		return
	}

	if currentProduct != (model.Product{}) {
		payload, _ := json.Marshal(struct {
			StatusCode	int 		`json:"statusCode"`
			Message		string 		`json:"message"`
			Data 		model.Product	`json:"product"`
		}{
			http.StatusOK, "Product is found!", currentProduct,
		})
		w.Write([]byte(payload))
		return
	}

	payload, _ := json.Marshal(struct {
		StatusCode	int	`json:"statusCode"`
		Message		string	`json:"message"`
	}{
		http.StatusInternalServerError, "Somethings wrong!",
	})
	w.Write([]byte(payload))
	return
})

var UpdateProductByProductId = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	var productId int

	var err error
	if productId, err = strconv.Atoi(vars["productId"]); err != nil {
		payload, _ := json.Marshal(struct {
			StatusCode	int	`json:"statusCode"`
			Message		string	`json:"message"`
			Errors		string	`json:"errors"`
		}{
			http.StatusBadRequest, "Invalid request", "BadRequest",
		})
		w.Write([]byte(payload))
		return
	}

	var productRequest model.Product
	if err = json.NewDecoder(r.Body).Decode(&productRequest); err != nil {
		payload, _ := json.Marshal(struct {
			StatusCode	int	`json:"statusCode"`
			Message		string	`json:"message"`
			Errors		string	`json:"errors"`
		}{
			http.StatusBadRequest, "invalid", fmt.Sprintf("%s", err),
		})
		w.Write([]byte(payload))
		return
	}

	if productRequest.CategoryId <= 0 {
		payload, _ := json.Marshal(struct {
			StatusCode	int	`json:"statusCode"`
			Message		string	`json:"message"`
			Errors 		string 	`json:"errors"`
		}{
			http.StatusBadRequest, "Product should have the category", "BadRequest",
		})
		w.Write([]byte(payload))
		return
	}

	if productRequest.ProductId != productId {
		payload, _ := json.Marshal(struct {
			StatusCode	int	`json:"statusCode"`
			Message		string	`json:"message"`
			Errors 		string 	`json:"errors"`
		}{
			http.StatusBadRequest, "ProductId is not same", "BadRequest",
		})
		w.Write([]byte(payload))
		return
	}

	if len(strings.Trim(productRequest.Name, " ")) == 0 {
		payload, _ := json.Marshal(struct {
			StatusCode	int	`json:"statusCode"`
			Message		string	`json:"message"`
			Errors		string	`json:"errors"`
		}{
			http.StatusBadRequest, "Product name can't be empty", "BadRequest",
		})
		w.Write([]byte(payload))
		return
	}

	if len(strings.Trim(productRequest.Unit, " ")) == 0 {
		payload, _ := json.Marshal(struct {
			StatusCode	int	`json:"statusCode"`
			Message		string	`json:"message"`
			Errors		string	`json:"errors"`
		}{
			http.StatusBadRequest, "Unit can't be empty", "BadRequest",
		})
		w.Write([]byte(payload))
		return
	}

	if productRequest.Stock < 0 {
		payload, _ := json.Marshal(struct {
			StatusCode	int	`json:"statusCode"`
			Message		string	`json:"message"`
			Errors		string	`json:"errors"`
		}{
			http.StatusBadRequest, "Stock can't be negative", "BadRequest",
		})
		w.Write([]byte(payload))
		return
	}

	if productRequest.AddSotck < 0 {
		payload, _ := json.Marshal(struct {
			StatusCode	int	`json:"statusCode"`
			Message		string	`json:"message"`
			Errors		string	`json:"errors"`
		}{
			http.StatusBadRequest, "Stock Add can't be negative", "BadRequest",
		})
		w.Write([]byte(payload))
		return
	}

	if productRequest.Price < 0 {
		payload, _ := json.Marshal(struct {
			StatusCode	int	`json:"statusCode"`
			Message		string	`json:"message"`
			Errors		string	`json:"errors"`
		}{
			http.StatusBadRequest, "Price can't be negative", "BadRequest",
		})
		w.Write([]byte(payload))
		return
	}

	var productModel model.Product
	var isProductExist bool = false
	if isProductExist, err = productModel.IsProductExistByProductId(productId); err != nil {
		payload, _ := json.Marshal(struct {
			StatusCode	int	`json:"statusCode"`
			Message		string	`json:"message"`
			Errors		string	`json:"errors"`
		}{
			http.StatusNotFound, "Product can't be found", fmt.Sprintf("%s", err),
		})
		w.Write([]byte(payload))
		return
	}

	var categoryModel model.Category
	var isCategoryExist bool = false
	if isCategoryExist, err = categoryModel.IsCategoryExistByCategoryId(productRequest.CategoryId); err != nil {
		payload, _ := json.Marshal(struct {
			StatusCode	int	`json:"statusCode"`
			Message		string	`json:"message"`
			Errors		string	`json:"errors"`
		}{
			http.StatusNotFound, "Category can't be found", fmt.Sprintf("%s", err),
		})
		w.Write([]byte(payload))
		return
	}

	if isProductExist && isCategoryExist {
		var currentProduct model.Product
		if currentProduct, err = productModel.FindProductByProductId(productId); err != nil {
			payload, _ := json.Marshal(struct {
				StatusCode	int	`json:"statusCode"`
				Message		string	`json:"message"`
				Errors		string	`json:"errors"`
			}{
				http.StatusInternalServerError, "Somethings wrong!", fmt.Sprintf("%s", err),
			})
			w.Write([]byte(payload))
			return
		}

		currentProduct.CategoryId = productRequest.CategoryId
		currentProduct.Name = productRequest.Name
		currentProduct.Unit = productRequest.Unit
		currentProduct.Price = productRequest.Price
		currentProduct.Stock = currentProduct.Stock + productRequest.AddSotck
		var updatedAt = time.Now().Format("2006-01-02 15:05:03")
		currentProduct.Audit.UpdatedAt = &updatedAt

		var updatedProduct model.Product
		if updatedProduct, err = currentProduct.UpdateProductByProductId(productId); err != nil {
			payload, _ := json.Marshal(struct {
				StatusCode	int	`json:"statusCode"`
				Message		string	`json:"message"`
				Errors		string	`json:"errors"`
			}{
				http.StatusInternalServerError, "Somethings wrong!", fmt.Sprintf("%s", err),
			})
			w.Write([]byte(payload))
			return
		} else {
			payload, _ := json.Marshal(struct {
				StatusCode	int		`json:"statusCode"`
				Message		string		`json:"message"`
				Data		model.Product	`json:"product"`
			}{
				http.StatusOK, "Success to update the product", updatedProduct,
			})
			w.Write([]byte(payload))
			return
		}
	}
	payload, _ := json.Marshal(struct {
		StatusCode	int	`json:"statusCode"`
		Message		string	`json:"message"`
	}{
		http.StatusInternalServerError, "Somethings wrong!",
	})
	w.Write([]byte(payload))
	return
})

var DeleteProductByProductId = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	var productId int

	var err error
	if productId, err = strconv.Atoi(vars["productId"]); err != nil {
		payload, _ := json.Marshal(struct {
			StatusCode	int	`json:"statusCode"`
			Message		string	`json:"message"`
			Errors		string	`json:"errors"`
		}{
			http.StatusBadRequest, "Invalid request", "BadRequest",
		})
		w.Write([]byte(payload))
		return
	}

	var productModel model.Product
	var isExist bool = false
	if isExist, err = productModel.IsProductExistByProductId(productId); err != nil {
		payload, _ := json.Marshal(struct {
			StatusCode	int	`json:"statusCode"`
			Message		string	`json:"message"`
			Errors		string	`json:"errors"`
		}{
			http.StatusNotFound, "Product can't be found", fmt.Sprintf("%s", err),
		})
		w.Write([]byte(payload))
		return
	}

	if isExist {
		var isDeleted bool = false
		if isDeleted, err = productModel.DeleteProductByProductId(productId); err != nil {
			payload, _ := json.Marshal(struct {
				StatusCode	int	`json:"statusCode"`
				Message		string	`json:"message"`
				Errors		string	`json:"errors"`
			}{
				http.StatusInternalServerError, "Somethings wrong!", fmt.Sprintf("%s", err),
			})
			w.Write([]byte(payload))
			return
		}

		if isDeleted {
			payload, _ := json.Marshal(struct {
				StatusCode	int 	`json:"statusCode"`
				Message 	string	`json:"message"`
			}{
				http.StatusOK, "Success to delete the product",
			})
			w.Write([]byte(payload))
			return
		}
	}

	payload, _ := json.Marshal(struct {
		StatusCode	int	`json:"statusCode"`
		Message		string	`json:"message"`
	}{
		http.StatusInternalServerError, "Somethings wrong!",
	})
	w.Write([]byte(payload))
	return
})
