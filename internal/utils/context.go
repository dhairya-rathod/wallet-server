package utils

import (
	"context"

	"github.com/golang-jwt/jwt"
)

// type contextKey string

func AddToContext(ctx context.Context, contextKey string, value interface{}) context.Context {
	return context.WithValue(ctx, contextKey, value)
}

func GetFromContext(contextKey string, ctx context.Context) interface{} {
	return ctx.Value(contextKey)
}

// this retrieves the JwtClaims from the context
func GetUserFromContext(ctx context.Context) JwtClaims {
	userClaims := ctx.Value("user").(jwt.MapClaims)
	userIDFloat64 := userClaims["id"].(float64)
	return JwtClaims{
		UserID: int(userIDFloat64),
		Name:   userClaims["name"].(string),
		Email:  userClaims["email"].(string),
	}
}
