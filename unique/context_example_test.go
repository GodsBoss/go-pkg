package unique_test

import (
	"github.com/GodsBoss/go-pkg/unique"

	"context"
	"fmt"
)

var key = unique.NewKey()

func NewContext(ctx context.Context, value string) context.Context {
	return context.WithValue(ctx, key, value)
}

func FromContext(ctx context.Context) (string, bool) {
	val := ctx.Value(key)
	s, ok := val.(string)
	if !ok {
		return "", false
	}
	return s, true
}

func ExampleNewKey_context() {
	ctx := NewContext(context.Background(), "world")
	val, ok := FromContext(ctx)
	if ok {
		fmt.Printf("Hello, %s!", val)
	}

	// Output:
	// Hello, world!
}
