package middleware

import (
	"fmt"
	"net/http"
	"wallet-server/internal/database"
	"wallet-server/internal/utils"
)

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		tokenString := r.Header.Get("Authorization")
		if tokenString == "" {
			w.WriteHeader(http.StatusUnauthorized)
			fmt.Fprint(w, "Missing authorization header")
			return
		}
		tokenString = tokenString[len("Bearer "):]

		token, err := utils.VerifyToken(tokenString)
		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			fmt.Fprint(w, "Invalid authentication token")
			return
		}
		claims := token.Claims.(*utils.JwtClaims)

		// tokenRecord, err := db.Service.GetRecordByToken(tokenString)
		var db *database.Service
		tokenRecord, err := db.GetRecordByToken(tokenString)
		if err != nil || tokenRecord.IsRevoked {
			w.WriteHeader(http.StatusUnauthorized)
			fmt.Fprint(w, "Authentication token expired")
			return
		}

		ctx := utils.AddToContext(r.Context(), "user", claims)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
