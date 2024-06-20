package middleware

import (
	// "context"
	"net/http"
	"wallet-server/internal/database"
	"wallet-server/internal/utils"
)

func AuthMiddleware(db *database.Service) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			tokenString := r.Header.Get("Authorization")
			if tokenString == "" {
				utils.SendErrorResponse(w, "Missing authorization header", http.StatusUnauthorized)
				return
			}
			tokenString = tokenString[len("Bearer "):]

			token, err := utils.VerifyToken(tokenString)
			if err != nil {
				utils.SendErrorResponse(w, "Invalid authentication token", http.StatusUnauthorized)
				return
			}
			claims := token.Claims

			tokenRecord, err := db.GetRecordByToken(tokenString)
			if err != nil || tokenRecord.IsRevoked {
				utils.SendErrorResponse(w, "Authentication token expired", http.StatusUnauthorized)
				return
			}

			ctx := utils.AddToContext(r.Context(), utils.UserContextKey, claims)

			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}
