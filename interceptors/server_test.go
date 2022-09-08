package interceptors

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc"
)

type beforeInterceptor struct{}
type beforeHandlerFunc struct{}

func TestUnaryServerInterceptor(t *testing.T) {
	ctx := context.WithValue(context.Background(), beforeInterceptor{}, "beforeInterceptor")
	grpcInterceptor := UnaryServerInterceptor(interceptor{
		beforeHandler: func(ctx context.Context) context.Context {
			return context.WithValue(ctx, beforeHandlerFunc{}, "beforeHandlerFunc")
		},
		afterHandler: func(ctx context.Context, err error) {},
	})

	_, _ = grpcInterceptor(ctx, nil, nil, func(ctx context.Context, req interface{}) (interface{}, error) {
		assert.Equal(t, "beforeInterceptor", ctx.Value(beforeInterceptor{}))
		assert.Equal(t, "beforeHandlerFunc", ctx.Value(beforeHandlerFunc{}))
		return nil, nil
	})
}

func TestStreamServerInterceptor(t *testing.T) {
	ctx := context.WithValue(context.Background(), beforeInterceptor{}, "beforeInterceptor")
	grpcInterceptor := StreamServerInterceptor(interceptor{
		beforeHandler: func(ctx context.Context) context.Context {
			return context.WithValue(ctx, beforeHandlerFunc{}, "beforeHandlerFunc")
		},
		afterHandler: func(ctx context.Context, err error) {},
	})

	_ = grpcInterceptor(nil, serverStream{ctx: ctx}, nil, func(srv interface{}, stream grpc.ServerStream) error {
		assert.Equal(t, "beforeInterceptor", stream.Context().Value(beforeInterceptor{}))
		assert.Equal(t, "beforeHandlerFunc", stream.Context().Value(beforeHandlerFunc{}))
		return nil
	})
}
