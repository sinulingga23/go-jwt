package controller

import (
	"fmt"
	"net/http"
	"encoding/json"

	"github.com/sinulingga23/go-jwt/model"
)



var GetCategories = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var categories []model.Category
	var categoryModel model.Category
	var err error

	if categories, err = categoryModel.GetAllCategory(); err != nil {
		payload, _ := json.Marshal(struct {
			StatusCode	int		`json:"statusCode"`
			Message		string		`json:"message"`
			Errors		string		`json:"errors"`
		}{
			http.StatusNotFound,
			"Somethings wrong!",
			fmt.Sprintf("%s", err),
		})
		w.Write([]byte(payload))
		return
	}

	payload, _ := json.Marshal(struct {
		StatusCode	int			`json:"statusCode"`
		Message		string			`json:"message"`
		Data		[]model.Category	`json:"categories"`
	}{
		http.StatusOK,
		"Success to get categories",
		categories,
	})
	w.Write([]byte(payload))
	return
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

var GetProductsByCategoryId = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("GetProductsByCategoryId"))
})
