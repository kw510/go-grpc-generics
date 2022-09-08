package interceptors

import (
	"context"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc"
)

func TestUnaryClientInterceptor(t *testing.T) {
	ctx := context.WithValue(context.Background(), beforeInterceptor{}, "beforeInterceptor")
	grpcInterceptor := UnaryClientInterceptor(interceptor{
		beforeHandler: func(ctx context.Context) context.Context {
			return context.WithValue(ctx, beforeHandlerFunc{}, "beforeHandlerFunc")
		},
		afterHandler: func(ctx context.Context, err error) {
			assert.Equal(t, "beforeInterceptor", ctx.Value(beforeInterceptor{}))
			assert.Equal(t, "beforeHandlerFunc", ctx.Value(beforeHandlerFunc{}))
			assert.ErrorContains(t, err, "invoker error")
		},
	})

	_ = grpcInterceptor(ctx, t.Name(), nil, nil, nil, func(ctx context.Context, method string, req, reply interface{}, cc *grpc.ClientConn, opts ...grpc.CallOption) error {
		return errors.New("invoker error")

	})
}

func TestStreamClientInterceptor(t *testing.T) {
	ctx := context.WithValue(context.Background(), beforeInterceptor{}, "beforeInterceptor")
	grpcInterceptor := StreamClientInterceptor(interceptor{
		beforeHandler: func(ctx context.Context) context.Context {
			return context.WithValue(ctx, beforeHandlerFunc{}, "beforeHandlerFunc")
		},
		afterHandler: func(ctx context.Context, err error) {
			assert.Equal(t, "beforeInterceptor", ctx.Value(beforeInterceptor{}))
			assert.Equal(t, "beforeHandlerFunc", ctx.Value(beforeHandlerFunc{}))
			assert.ErrorContains(t, err, "invoker error")
		},
	})

	_, _ = grpcInterceptor(ctx, nil, nil, t.Name(), func(ctx context.Context, desc *grpc.StreamDesc, cc *grpc.ClientConn, method string, opts ...grpc.CallOption) (grpc.ClientStream, error) {
		return nil, errors.New("invoker error")
	})
}
