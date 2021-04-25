package auth

import (
	"os"
	"fmt"
	"log"
	"time"
	"strings"
	"net/http"
	"encoding/json"

	"github.com/sinulingga23/go-jwt/model"
	"github.com/dgrijalva/jwt-go"
)

func CreateToken(email string) (string, error) {
	claims := &model.Claims {
		Email: email,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 1).Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	var tokenString string
	var err error
	if tokenString, err = token.SignedString([]byte(os.Getenv("JWT_SECRET_KEY"))); err != nil {
		return "", err
	}
	return tokenString, nil
}

func ExtractToken(r *http.Request) string {
	keys := r.URL.Query()
	token := keys.Get("token")
	if token != "" {
		return token
	}
	bearerToken := r.Header.Get("Authorization")
	if len(strings.Split(bearerToken, " ")) == 2 {
		return strings.Split(bearerToken, " ")[1]
	}
	return ""
}

func TokenValid(r *http.Request) error {
	tokenString := ExtractToken(r)
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("JWT_SECRET_KEY")), nil
	})
	if err != nil {
		return err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		b, err := json.MarshalIndent(claims, "", " ")
		if err != nil {
			log.Printf("%v", err)
		}
		fmt.Printf("%v", string(b))
	}
	return nil
}
