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
	w.Header().Set("Content-Type", "application/json")

	var categoryRequest model.Category
	var err error

	if err = json.NewDecoder(r.Body).Decode(&categoryRequest); err != nil {
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

	if len(strings.Trim(categoryRequest.Category, " ")) == 0 {
		payload, _ := json.Marshal(struct {
			StatusCode	int	`json:"statusCode"`
			Message		string	`json:"message"`
			Errors		string	`json:"errors"`
		}{
			http.StatusBadRequest, "Category name can't be empty.", "BadRequest",
		})
		w.Write([]byte(payload))
		return
	}

	var createdCategory model.Category
	categoryRequest.Audit.CreatedAt = time.Now().Format("2006-01-02 15:05:03")
	if createdCategory, err = categoryRequest.SaveCategory(); err != nil {
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
		Data		model.Category	`json:"category"`
	}{
		http.StatusOK, "success to create a new category", createdCategory,
	})
	w.Write([]byte(payload))
	return
})

var GetCategoryByCategoryId = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	var categoryId int

	var err error
	if  categoryId, err = strconv.Atoi(vars["categoryId"]); err != nil {
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

	var categoryModel model.Category
	if _, err = categoryModel.IsCategoryExistByCategoryId(categoryId); err != nil {
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

	var currentCategory model.Category
	if currentCategory, err = categoryModel.FindCategoryByCategoryId(categoryId); err != nil {
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

	payload, _ := json.Marshal(struct {
		StatusCode	int 		`json:"statusCode"`
		Message 	string		`json:"message"`
		Data		model.Category	`json:"category"`
	}{
		http.StatusOK, "Category is found!", currentCategory,
	})
	w.Write([]byte(payload))
	return
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
