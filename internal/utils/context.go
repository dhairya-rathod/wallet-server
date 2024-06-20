package utils

import (
	"context"
	"fmt"

	"github.com/golang-jwt/jwt"
)

type contextKey string

var UserContextKey = contextKey("user")

func AddToContext(ctx context.Context, contextKey contextKey, value interface{}) context.Context {
	return context.WithValue(ctx, contextKey, value)
}

func GetFromContext(contextKey contextKey, ctx context.Context) interface{} {
	return ctx.Value(contextKey)
}

// this retrieves the JwtClaims from the context
func GetUserFromContext(ctx context.Context) (JwtClaims, error) {
	user := ctx.Value(UserContextKey)
	userClaims, ok := user.(jwt.MapClaims)
	if !ok {
		return JwtClaims{}, fmt.Errorf("user is not of type jwt.MapClaims")
	}
	return JwtClaims{
		UserID: int(userClaims["userId"].(float64)),
		Name:   userClaims["name"].(string),
		Email:  userClaims["email"].(string),
	}, nil
}
