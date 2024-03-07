package middlewares

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/golang-jwt/jwt/v5"
	"github.com/lai0xn/cr-dermasuim/app/utils"
)

func IsAuthenticated(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			fmt.Println("auth header not found")
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(H{
				"message": "invalid token",
			})
			return
		}
		args := strings.Split(authHeader, " ")
		if len(args) < 2 {
			fmt.Println("length")
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(H{
				"message": "please provide a token",
			})
			return

		}
		token, err := utils.ParseToken(args[1])
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Println("parsing error")
			json.NewEncoder(w).Encode(H{
				"message": "invalid token",
			})
			return

		}
		if !token.Valid {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Println("token invalid")
			json.NewEncoder(w).Encode(H{
				"message": "invalid token",
			})
			return

		}
		claims := token.Claims.(jwt.MapClaims)
		id, ok := claims["userID"]
		if !ok {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(H{
				"message": "invalid token",
			})
			return

		}
		ctx := context.WithValue(r.Context(), "userID", id)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
