package auth

import (
	"context"
	"fmt"
	"github.com/go-kratos/kratos/v2/middleware"
)

func Server() middleware.Middleware {
	return func(handler middleware.Handler) middleware.Handler {
		return func(ctx context.Context, req interface{}) (interface{}, error) {

			fmt.Println("Enter auth middleware")

			return handler(ctx, req)
		}
	}
}
