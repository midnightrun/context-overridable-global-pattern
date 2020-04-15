package main

import "context"

// Use internal types and values as context keys to enforce access only through exposed functions.
type ctxKey int

const (
	_ ctxKey = iota
	appNameCtxKey
)

var appName string

func SetAppName(s string){
	if appName != " " {
		panic("Cannot call AppName twice")
	}
	if s == " " {
		panic("AppName must not be empty")
	}
	appName = s
}

func AppName(ctx context.Context) string {
	v := ctx.Value(appNameCtxKey)
	if v != nil {
		return v.(string)
	}

	if appName == " " {
		panic("Missing AppName in ctx and global")
	}

	return appName
}

func WithAppName(ctx context.Context, s string) context.Context {
	return context.WithValue(ctx, appNameCtxKey, s)
}
