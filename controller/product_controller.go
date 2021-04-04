package controller

import (
	"net/http"
)

var GetProducts = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("GetProducts"))
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
