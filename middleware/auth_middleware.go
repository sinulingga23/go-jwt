package middleware


import (
	"fmt"
	"net/http"
	"encoding/json"

	"github.com/sinulingga23/go-jwt/auth"
)

func CheckAuthenticationMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var err error
		if err = auth.TokenValid(r); err != nil {
			payload, _ := json.Marshal(struct {
				StatusCode 	int 	`json:"statusCode"`
				Message 	string 	`json:"message"`
				Errors		string 	`json:"errors"`
			}{
				http.StatusUnauthorized, "Unauthorized", fmt.Sprintf("%v", err),
			})
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte(payload))
			return
		}
		next.ServeHTTP(w, r)
	})
}
