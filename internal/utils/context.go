package utils

import (
	"context"
)

type contextKey string

func AddToContext(ctx context.Context, contextKey contextKey, value interface{}) context.Context {
	return context.WithValue(ctx, contextKey, value)
}

func GetFromContext(contextKey string, ctx context.Context) interface{} {
	return ctx.Value(contextKey)
}

// this retrieves the JwtClaims from the context
func GetUserFromContext(ctx context.Context) *JwtClaims {
	return ctx.Value("user").(*JwtClaims)
}
