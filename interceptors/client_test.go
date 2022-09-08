package interceptors

import (
	"context"
	"testing"

	"google.golang.org/grpc"
)

func TestUnaryClientInterceptor(t *testing.T) {
	ctx := context.Background()
	grpcInterceptor := UnaryClientInterceptor(interceptor{
		beforeHandler: func(ctx context.Context) context.Context {
			return ctx
		},
		afterHandler: func(ctx context.Context, err error) {},
	})

	_ = grpcInterceptor(ctx, t.Name(), nil, nil, nil, func(ctx context.Context, method string, req, reply interface{}, cc *grpc.ClientConn, opts ...grpc.CallOption) error {
		return nil
	})
}

func TestStreamClientInterceptor(t *testing.T) {
	ctx := context.Background()
	grpcInterceptor := StreamClientInterceptor(interceptor{
		beforeHandler: func(ctx context.Context) context.Context {
			return ctx
		},
		afterHandler: func(ctx context.Context, err error) {},
	})

	_, _ = grpcInterceptor(ctx, nil, nil, t.Name(), func(ctx context.Context, desc *grpc.StreamDesc, cc *grpc.ClientConn, method string, opts ...grpc.CallOption) (grpc.ClientStream, error) {
		return nil, nil
	})
}
