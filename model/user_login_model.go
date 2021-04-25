package model

import (
	"errors"

	database "github.com/sinulingga23/go-jwt/db"
	"golang.org/x/crypto/bcrypt"
)

type UserLogin struct {
	Email		string 	`json:"email"`
	Password	string 	`json:"password"`
}

func (uL *UserLogin) IsUserExistByEmail(email string) (bool, error) {
	db, err := database.ConnectDB()
	if err != nil {
		return false, err
	}
	defer db.Close()

	var check int64
	err = db.QueryRow("SELECT count(user_id) FROM user WHERE email = ?", email).Scan(&check)
	if err != nil {
		return false, err
	}

	if check != 1 {
		return false, errors.New("Somethings wrong!")
	}

	return true, nil
}

func (ul *UserLogin) IsUserValid(email string, password string) (bool, error) {
	db, err := database.ConnectDB()
	if err != nil {
		return false, err
	}
	defer db.Close()

	var	currentEmail		string
	var	currentPassword		string
	err = db.QueryRow("SELECT email, password FROM user WHERE email = ?", email).Scan(&currentEmail, &currentPassword)
	if err != nil {
		return false, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(currentPassword), []byte(password))
	if err != nil {
		return false, errors.New("Somethings wrong!")
	}

	return true, nil
}
