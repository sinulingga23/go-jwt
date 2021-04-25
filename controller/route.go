package controller

import (
	"net/http"
	"github.com/sinulingga23/go-jwt/middleware"
	"github.com/gorilla/mux"
)

func RunServer() {
	router := mux.NewRouter()

	// endpoints for authtenticaiton & authorization purpose
	router.Handle("/login", LoginUser).Methods("POST")

	// endpoints categories
	router.Handle("/categories", GetCategories).Methods("GET")
	router.Handle("/categories", middleware.CheckAuthenticationMiddleware(CreateCategory)).Methods("POST")
	router.Handle("/categories/{categoryId}", GetCategoryByCategoryId).Methods("GET")
	router.Handle("/categories/{categoryId}", middleware.CheckAuthenticationMiddleware(UpdateCategoryByCategoryId)).Methods("PUT")
	router.Handle("/categories/{categoryId}", middleware.CheckAuthenticationMiddleware(DeleteCategoryByCategoryId)).Methods("DELETE")
	router.Handle("/categories/{categoryId}/products", GetProductsByCategoryId).Methods("GET")

	// endpoints products
	router.Handle("/products", GetProducts).Methods("GET")
	router.Handle("/products", middleware.CheckAuthenticationMiddleware(CreateProduct)).Methods("POST")
	router.Handle("/products/{productId}", GetProductByProductId).Methods("GET")
	router.Handle("/products/{productId}", middleware.CheckAuthenticationMiddleware(UpdateProductByProductId)).Methods("PUT")
	router.Handle("/products/{productId}", middleware.CheckAuthenticationMiddleware(DeleteProductByProductId)).Methods("DELETE")

	http.ListenAndServe(":8080", router)
}
