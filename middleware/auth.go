package middleware

import (
	"context"
	"net/http"

	"booknest-api/utils"
)

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		tokenString := r.Header.Get("Authorization")
		if tokenString == "" {
			http.Error(w, "Missing authorization token", http.StatusUnauthorized)
			return
		}

		tokenString = tokenString[len("Bearer "):] // Remove "Bearer " prefix

		claims, err := utils.ParseJWT(tokenString)
		if err != nil {
			http.Error(w, "Invalid authorization token", http.StatusUnauthorized)
			return
		}

		// Store user ID in context for downstream handlers
		ctx := context.WithValue(r.Context(), "userID", claims["user_id"])
		ctx = context.WithValue(ctx, "role", claims["role"])
		next.ServeHTTP(w, r.WithContext(ctx))
	})
} 