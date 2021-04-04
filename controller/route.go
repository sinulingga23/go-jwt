package controller

import (
	"net/http"
	"github.com/gorilla/mux"
)

func RunServer() {
	router := mux.NewRouter()

	// endpoints categories
	router.Handle("/categories", GetCategories).Methods("GET")
	router.Handle("/categories", CreateCategory).Methods("POST")
	router.Handle("/categories/{categoryId}", GetCategoryByCategoryId).Methods("GET")
	router.Handle("/categories/{categoryId}", UpdateCategoryByCategoryId).Methods("PUT")
	router.Handle("/categories/{categoryId}", DeleteCategoryByCategoryId).Methods("DELETE")
	router.Handle("/categories/{categoryId}/products", GetProductsByCategoryId).Methods("GET")

	router.Handle("/products", GetProducts).Methods("GET")
	router.Handle("/products", CreateProduct).Methods("POST")
	router.Handle("/products/{productId}", GetProductByProductId).Methods("GET")
	router.Handle("/products/{productId}", UpdateProductByProductId).Methods("PUT")
	router.Handle("/products/{productId}", DeleteProductByProductId).Methods("DELETE")

	// endpoints products
	http.ListenAndServe(":8080", router)
}
