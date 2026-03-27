package middleware

import (
	"context"
	"net/http"
	"os" //auth
	"strings" //auth

	"github.com/golang-jwt/jwt/v5" //auth
)

// AuthMiddleware checks for a valid Bearer token
func AuthMiddleware(next http.Handler) http.Handler { //auth
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) { //auth
		// 1. Get the Authorization header
		authHeader := r.Header.Get("Authorization") //auth
		if authHeader == "" { //auth
			http.Error(w, "Authorization header missing", http.StatusUnauthorized) //auth
			return //auth
		} //auth

		// 2. Parse the "Bearer <token>" string
		tokenStr := strings.TrimPrefix(authHeader, "Bearer ") //auth
		if tokenStr == authHeader { //auth
			http.Error(w, "Invalid token format", http.StatusUnauthorized) //auth
			return //auth
		} //auth

		// 3. Validate the token
		secretKey := os.Getenv("JWT_SECRET") //auth
		claims := jwt.MapClaims{} //auth

		token, err := jwt.ParseWithClaims(tokenStr, claims, func(token *jwt.Token) (interface{}, error) { //auth
			return []byte(secretKey), nil //auth
		}) //auth

		if err != nil || !token.Valid { //auth
			http.Error(w, "Invalid or expired token", http.StatusUnauthorized) //auth
			return //auth
		} //auth

		// 4. Attach the user_id to the request context so handlers can use it
		ctx := context.WithValue(r.Context(), "user_id", claims["user_id"]) //auth
		next.ServeHTTP(w, r.WithContext(ctx)) //auth
	}) //auth
}