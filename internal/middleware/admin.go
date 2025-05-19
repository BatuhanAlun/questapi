package middleware

import (
	"net/http"
	"questapi/internal/utils"

	"github.com/golang-jwt/jwt/v5"
)

func AuthAdmin(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		tokenString := r.Header.Get("Authorization")

		if tokenString == "" {
			http.Error(w, "Authorization Header is required", http.StatusUnauthorized)
			return
		}

		if len(tokenString) > 7 && tokenString[:7] == "Bearer " {
			tokenString = tokenString[7:]
		}

		token, err := jwt.Parse(tokenString, utils.JWTKeyFunc)

		if err != nil {
			http.Error(w, "Invalid Token", http.StatusUnauthorized)
			return
		}

		claims, ok := token.Claims.(jwt.MapClaims)

		if !ok || !token.Valid {
			http.Error(w, "Invalid Token", http.StatusUnauthorized)
			return
		}

		role, ok := claims["role"].(string)

		if !ok {
			http.Error(w, "Invalid Token", http.StatusUnauthorized)
			return
		}
		if role != "admin" {
			http.Error(w, "Unauthorized User", http.StatusUnauthorized)
			return
		}

		next.ServeHTTP(w, r)

	})
}
