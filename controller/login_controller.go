package controller

import (
	"fmt"
	"net/http"
	"encoding/json"

	"github.com/sinulingga23/go-jwt/model"
	"github.com/sinulingga23/go-jwt/auth"
)


var LoginUser = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var userLoginRequest model.UserLogin
	var userLoginModel model.UserLogin
	var err error

	if err = json.NewDecoder(r.Body).Decode(&userLoginRequest); err != nil {
		payload, _ := json.Marshal(struct {
			StatusCode	int 	`json:"statusCode"`
			Message 	string	`json:"message"`
			Errors		string	`json:"errors"`
		}{
			http.StatusBadRequest, "invalid", fmt.Sprintf("%s", err),
		})
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(payload))
		return
	}

	var isExist bool = false
	if isExist, err = userLoginModel.IsUserExistByEmail(userLoginRequest.Email); err != nil {
		payload, _ := json.Marshal(struct {
			StatusCode	int	`json:"statusCode"`
			Message		string	`json:"message"`
			Errors		string	`json:"errors"`
		}{
			http.StatusNotFound, "User can't be found", fmt.Sprintf("%s", err),
		})
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(payload))
		return
	}

	if isExist {
		var isValid bool = false
		if isValid, err = userLoginRequest.IsUserValid(userLoginRequest.Email, userLoginRequest.Password); err != nil {
			payload, _ := json.Marshal(struct {
				StatusCode	int	`json:"statusCode"`
				Message		string	`json:"message"`
				Errors		string	`json:"errors"`
			}{
				http.StatusUnauthorized, "Make sure the email and password is correct.", fmt.Sprintf("%s", err),
			})
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte(payload))
			return
		}

		if isValid {
			// Create the JWT string
			tokenString, err := auth.CreateToken(userLoginRequest.Email)
			if err != nil {
				payload, _ := json.Marshal(struct {
					StatusCode	int 	`json:"statusCode"`
					Message 	string 	`json:"message"`
					Errors		string	`json:"errors"`
				}{
					http.StatusInternalServerError, "Somethings wrong!", fmt.Sprintf("%s", err),
				})
				w.WriteHeader(http.StatusInternalServerError)
				w.Write([]byte(payload))
				return
			}
			payload, _ := json.Marshal(struct {
				StatusCode	int 	`json:"statusCode"`
				Token		string 	`json:"token"`
			}{
				http.StatusOK, tokenString,
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
