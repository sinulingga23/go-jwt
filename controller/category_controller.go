package controller

import (
	"net/http"
)

var GetCategories = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("GetCategories"))
})

var CreateCategory = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("CreateCategory"))
})

var GetCategoryByCategoryId = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("GetCategoryByCategoryId"))
})

var UpdateCategoryByCategoryId = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("UpdateCategoryByCategoryId"))
})

var DeleteCategoryByCategoryId = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("DeleteCategoryByCategoryId"))
})
