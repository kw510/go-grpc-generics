package server

import (
	"context"

	"google.golang.org/grpc"
)

type Interceptor interface {
	BeforeHandler(ctx context.Context) context.Context
	AfterHandler(ctx context.Context, err error)
}

// Creates a gRPC unary interceptor from a generic interceptor
func UnaryInterceptor(interceptor Interceptor) grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (any, error) {
		ctx = interceptor.BeforeHandler(ctx)
		resp, err := handler(ctx, req)
		interceptor.AfterHandler(ctx, err)
		return resp, err
	}
}

// Creates a gRPC stream interceptor from a generic interceptor
func StreamInterceptor(interceptor Interceptor) grpc.StreamServerInterceptor {
	return func(srv any, ss grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) error {
		ctx := interceptor.BeforeHandler(ss.Context())
		err := handler(srv, &wrapper{Ctx: ctx, ServerStream: ss})
		interceptor.AfterHandler(ctx, err)
		return err
	}
}
