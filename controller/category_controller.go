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
		w.WriteHeader(http.StatusNotFound)
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
	w.WriteHeader(http.StatusOK)
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
		w.WriteHeader(http.StatusBadRequest)
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
		w.WriteHeader(http.StatusBadRequest)
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
		w.WriteHeader(http.StatusInternalServerError)
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
	w.WriteHeader(http.StatusOK)
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
		w.WriteHeader(http.StatusBadRequest)
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
			http.StatusNotFound, "Category can't be found", fmt.Sprintf("%s", err),
		})
		w.WriteHeader(http.StatusNotFound)
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
		w.WriteHeader(http.StatusBadRequest)
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
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(payload))
	return
})

var UpdateCategoryByCategoryId = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	var categoryId int

	var err error
	if categoryId, err = strconv.Atoi(vars["categoryId"]); err != nil {
		payload, _ := json.Marshal(struct {
			StatusCode	int	`json:"statusCode"`
			Message		string	`json:"message"`
			Errors		string	`json:"errors"`
		}{
			http.StatusBadRequest, "Invalid request", "BadRequest",
		})
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(payload))
		return
	}

	var categoryRequest model.Category
	if err = json.NewDecoder(r.Body).Decode(&categoryRequest); err != nil {
		payload, _ := json.Marshal(struct {
			StatusCode	int	`json:"statusCode"`
			Message		string	`json:"message"`
			Errors		string	`json:"errors"`
		}{
			http.StatusBadRequest, "invalid", fmt.Sprintf("%s", err),
		})
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(payload))
		return
	}

	if categoryRequest.CategoryId != categoryId {
		payload, _ := json.Marshal(struct {
			StatusCode	int	`json:"statusCode"`
			Message		string	`json:"message"`
			Errors 		string 	`json:"errors"`
		}{
			http.StatusBadRequest, "CategoryId is not same", "BadRequest",
		})
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(payload))
		return
	}

	var categoryModel model.Category
	var isExist bool = false
	if isExist, err = categoryModel.IsCategoryExistByCategoryId(categoryId); err != nil {
		payload, _ := json.Marshal(struct {
			StatusCode	int	`json:"statusCode"`
			Message		string	`json:"message"`
			Errors		string	`json:"errors"`
		}{
			http.StatusNotFound, "Category can't be found", fmt.Sprintf("%s", err),
		})
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(payload))
		return
	}

	if isExist {
		var currentCategory model.Category
		if currentCategory, err = categoryModel.FindCategoryByCategoryId(categoryId); err != nil {
			payload, _ := json.Marshal(struct {
				StatusCode	int	`json:"statusCode"`
				Message		string	`json:"message"`
				Errors		string	`json:"errors"`
			}{
				http.StatusInternalServerError, "Somethings wrong!", fmt.Sprintf("%s", err),
			})
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(payload))
			return
		}

		currentCategory.Category = categoryRequest.Category
		var updatedAt string = time.Now().Format("2006-01-02 15:05:03")
		currentCategory.Audit.UpdatedAt = &updatedAt

		var updatedCategory model.Category
		if updatedCategory, err = currentCategory.UpdateCategoryByCategoryId(categoryId); err != nil {
			payload, _ := json.Marshal(struct {
				StatusCode	int	`json:"statusCode"`
				Message		string	`json:"message"`
				Errors		string	`json:"errors"`
			}{
				http.StatusInternalServerError, "Somethings wrong!", fmt.Sprintf("%s", err),
			})
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(payload))
			return
		} else {
			payload, _ := json.Marshal(struct {
				StatusCode	int		`json:"statusCode"`
				Message		string		`json:"message"`
				Data		model.Category	`json:"category"`
			}{
				http.StatusOK, "success to update the category", updatedCategory,
			})
			w.WriteHeader(http.StatusOK)
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
	w.WriteHeader(http.StatusInternalServerError)
	w.Write([]byte(payload))
	return
})

var DeleteCategoryByCategoryId = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	var categoryId int

	var err error
	if categoryId, err = strconv.Atoi(vars["categoryId"]); err != nil || len(strings.Trim(vars["categoryId"], " ")) == 0 {
		payload, _ := json.Marshal(struct {
			StatusCode	int	`json:"statusCode"`
			Message		string	`json:"message"`
			Errors		string	`json:"errors"`
		}{
			http.StatusBadRequest, "Invalid request", "BadRequest",
		})
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(payload))
		return
	}

	var categoryModel model.Category
	var isExist bool = false
	if isExist, err = categoryModel.IsCategoryExistByCategoryId(categoryId); err != nil {
		payload, _ := json.Marshal(struct {
			StatusCode	int	`json:"statusCode"`
			Message		string	`json:"message"`
			Errors		string	`json:"errors"`
		}{
			http.StatusNotFound, "Category can't be found", fmt.Sprintf("%s", err),
		})
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(payload))
		return
	}

	if isExist {
		var isDeleted bool = false
		if isDeleted, err = categoryModel.DeleteCategoryByCategoryId(categoryId); err != nil {
			payload, _ := json.Marshal(struct {
				StatusCode	int	`json:"statusCode"`
				Message		string	`json:"message"`
				Errors		string	`json:"errors"`
			}{
				http.StatusInternalServerError, "Somethings wrong!", fmt.Sprintf("%s", err),
			})
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(payload))
			return
		}

		if isDeleted {
			payload, _ := json.Marshal(struct {
				StatusCode	int 	`json:"statusCode"`
				Message		string 	`json:"message"`
			}{
				http.StatusOK, "success to delete the category",
			})
			w.WriteHeader(http.StatusOK)
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
	w.WriteHeader(http.StatusInternalServerError)
	w.Write([]byte(payload))
	return
})

var GetProductsByCategoryId = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	var categoryId int

	var err error
	if categoryId, err = strconv.Atoi(vars["categoryId"]); err != nil {
		payload, _ := json.Marshal(struct {
			StatusCode	int	`json:"statusCode"`
			Message		string	`json:"message"`
			Errors		string	`json:"errors"`
		}{
			http.StatusBadRequest, "Invalid request", "BadRequest",
		})
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(payload))
		return
	}

	var categoryModel model.Category
	var isExist bool = false
	if isExist, err = categoryModel.IsCategoryExistByCategoryId(categoryId); err != nil {
		payload, _ := json.Marshal(struct {
			StatusCode	int	`json:"statusCode"`
			Message		string	`json:"message"`
			Errors		string	`json:"errors"`
		}{
			http.StatusNotFound, "Category can't be found", fmt.Sprintf("%s", err),
		})
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(payload))
		return
	}

	if isExist {
		var products []model.Product
		if products, err = categoryModel.FindProductsByCategoryId(categoryId); err != nil {
			payload, _ := json.Marshal(struct {
				StatusCode	int	`json:"statusCode"`
				Message		string	`json:"message"`
				Errors		string	`json:"errors"`
			}{
				http.StatusInternalServerError, "Somethings wrong!", fmt.Sprintf("%s", err),
			})
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(payload))
			return
		}

		if products != nil {
			payload, _ := json.Marshal(struct {
				StatusCode	int 		`json:"statusCode"`
				Message		string 		`json:"message"`
				CategoryId	int 		`json:"categoryId"`
				Data		[]model.Product	`json:"products"`
			}{
				http.StatusOK, "succes to get the products", categoryId, products,
			})
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(payload))
			return
		} else {
			payload, _ := json.Marshal(struct {
				StatusCode	int	`json:"statusCode"`
				Message		string	`json:"message"`
				Errors		string	`json:"errors"`
			}{
				http.StatusNotFound, "The category don't have the products", "NotFound",
			})
			w.WriteHeader(http.StatusNotFound)
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
	w.WriteHeader(http.StatusInternalServerError)
	w.Write([]byte(payload))
	return
})
