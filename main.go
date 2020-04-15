package main

import (
	"context"
	"fmt"
)

func main() {
	app.SetAppName("Global Value")

	// The background context will use the global value.
	fmt.Println("Background context:", app.Appname(context.Background()))

	// A context can override the global value.
	ctx := app.WithAppName(context.Background(), "Context Value")
	fmt.Println("WithAppName context:", app.Appname(ctx))

	// Verifying that the global value is unchanged
	fmt.Println("Background context:", app.AppName(context.Background()))
}
