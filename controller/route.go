package controller

import (
	"net/http"
	"github.com/gorilla/mux"
)

func RunServer() {
	router := mux.NewRouter()

	router.Handle("/categories", GetCategories).Methods("GET")
	router.Handle("/categories", CreateCategory).Methods("POST")
	router.Handle("/categories/{categoryId}", GetCategoryByCategoryId).Methods("GET")
	router.Handle("/categories/{categoryId}", UpdateCategoryByCategoryId).Methods("PUT")
	router.Handle("/categories/{categoryId}", DeleteCategoryByCategoryId).Methods("DELETE")
	router.Handle("/categories/{categoryId}/products", GetProductsByCategoryId).Methods("GET")

	http.ListenAndServe(":8080", router)
}

var NotImplemented = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Not Implemented"))
})
