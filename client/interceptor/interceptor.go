package interceptor

import (
	"context"

	"google.golang.org/grpc"
)

type Interceptor interface {
	BeforeHandler(ctx context.Context) context.Context
	AfterHandler(ctx context.Context, err error)
}

// Creates a client gRPC unary interceptor from a generic interceptor
func UnaryInterceptor(interceptor Interceptor) grpc.UnaryClientInterceptor {
	return func(ctx context.Context, method string, req, reply interface{}, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
		ctx = interceptor.BeforeHandler(ctx)
		err := invoker(ctx, method, req, reply, cc, opts...)
		interceptor.AfterHandler(ctx, err)
		return err
	}
}

// Creates a client gRPC streaming interceptor from a generic interceptor
func StreamInterceptor(interceptor Interceptor) grpc.StreamClientInterceptor {
	return func(ctx context.Context, desc *grpc.StreamDesc, cc *grpc.ClientConn, method string, streamer grpc.Streamer, opts ...grpc.CallOption) (grpc.ClientStream, error) {
		ctx = interceptor.BeforeHandler(ctx)
		cs, err := streamer(ctx, desc, cc, method, opts...)
		interceptor.AfterHandler(ctx, err)
		return cs, err
	}
}
