package utils

import (
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt"
)

type JwtClaims struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
	jwt.StandardClaims
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

	// token, err := jwt.ParseWithClaims(tokenString, &JwtClaims{}, func(token *jwt.Token) (interface{}, error) {
	// 	if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
	// 		return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
	// 	}
	// 	return []byte(secretKey), nil
	// })

	// if err != nil {
	// 	return nil, err
	// }

	// if claims, ok := token.Claims.(*JwtClaims); ok && token.Valid {
	// 	return claims, nil
	// }

	// return nil, fmt.Errorf("invalid token")
}
