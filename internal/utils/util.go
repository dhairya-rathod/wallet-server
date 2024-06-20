package utils

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
)

type JwtClaims struct {
	UserID int    `json:"userId"`
	Name   string `json:"name"`
	Email  string `json:"email"`
	jwt.StandardClaims
}

type ErrorResponse struct {
	Message string `json:"message"`
	Status  int    `json:"status"`
}

func GetEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}

// * token handler
func GenerateToken(jwtClaims JwtClaims) (string, error) {
	secretKey := GetEnv("TOKEN_SECRET", "")

	// * set token expiration time to two months from now
	jwtClaims.StandardClaims = jwt.StandardClaims{
		ExpiresAt: time.Now().AddDate(0, 2, 0).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwtClaims)

	tokenString, err := token.SignedString([]byte(secretKey))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func VerifyToken(tokenString string) (*jwt.Token, error) {
	secretKey := GetEnv("TOKEN_SECRET", "")

	return jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(secretKey), nil
	})
}

// * password handler
func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

// * helper function to send error response
func SendErrorResponse(w http.ResponseWriter, message string, statusCode int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(ErrorResponse{
		Message: message,
		Status:  statusCode,
	})
}

// * helper function to get integer url parameter
func GetIntegerURLParam(r *http.Request, param string) (int, error) {
	idStr := chi.URLParam(r, param)
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return 0, err
	}
	return id, nil
}
