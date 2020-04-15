package main

import (
	"context"
	"fmt"
)

func main() {
	SetAppName("Global Value")

	// The background context will use the global value.
	fmt.Println("Background context:", AppName(context.Background()))

	// A context can override the global value.
	ctx := WithAppName(context.Background(), "Context Value")
	fmt.Println("WithAppName context:", AppName(ctx))

	// Verifying that the global value is unchanged
	fmt.Println("Background context:", AppName(context.Background()))
}
