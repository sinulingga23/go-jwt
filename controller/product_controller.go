package controller

import (
	"fmt"
	"net/http"
	"encoding/json"

	"github.com/sinulingga23/go-jwt/model"
	_"github.com/gorilla/mux"
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
	w.Write([]byte("CreateProduct"))
})

var GetProductByProductId = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("GetProductByProductId"))
})

var UpdateProductByProductId = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("UpdateProductByProductId"))
})

var DeleteProductByProductId = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("DeleteProductByProductId"))
})
