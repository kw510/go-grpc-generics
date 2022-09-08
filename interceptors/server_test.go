package interceptors

import (
	"context"
	"testing"

	"google.golang.org/grpc"
)

func TestUnaryServerInterceptor(t *testing.T) {
	ctx := context.Background()
	grpcInterceptor := UnaryServerInterceptor(interceptor{
		beforeHandler: func(ctx context.Context) context.Context {
			return ctx
		},
		afterHandler: func(ctx context.Context, err error) {},
	})

	_, _ = grpcInterceptor(ctx, nil, nil, func(ctx context.Context, req interface{}) (interface{}, error) {
		return nil, nil
	})
}

func TestStreamServerInterceptor(t *testing.T) {
	ctx := context.Background()
	grpcInterceptor := StreamServerInterceptor(interceptor{
		beforeHandler: func(ctx context.Context) context.Context {
			return ctx
		},
		afterHandler: func(ctx context.Context, err error) {},
	})

	_ = grpcInterceptor(nil, serverStream{ctx: ctx}, nil, func(srv interface{}, stream grpc.ServerStream) error {
		return nil
	})
}
