package model

import (
	"errors"

	database "github.com/sinulingga23/go-jwt/db"
)

type UserLogin struct {
	Email		string 	`json:"username"`
	Password	string 	`json:"password"`
}

func (uL *UserLogin) IsUserExistByEmail() (bool, error) {
	db, err := database.ConnectDB()
	if err != nil {
		return false, err
	}
	defer db.Close()

	var check int64
	err = db.QueryRow("SELECT count(user_id) FROM user WHERE email = ?", uL.Email).Scan(&check)
	if err != nil {
		return false, err
	}

	if check != 1 {
		return false, errors.New("Somethings wrong!")
	}

	return true, nil
}
