package interceptors

import (
	"context"

	"google.golang.org/grpc"
)

// Creates a gRPC unary interceptor from a generic interceptor
func UnaryServerInterceptor(interceptor Interceptor) grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (any, error) {
		ctx = interceptor.BeforeHandler(ctx)
		resp, err := handler(ctx, req)
		interceptor.AfterHandler(ctx, err)
		return resp, err
	}
}

// Creates a gRPC stream interceptor from a generic interceptor
func StreamServerInterceptor(interceptor Interceptor) grpc.StreamServerInterceptor {
	return func(srv any, ss grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) error {
		ctx := interceptor.BeforeHandler(ss.Context())
		err := handler(srv, &serverWrapper{Ctx: ctx, ServerStream: ss})
		interceptor.AfterHandler(ctx, err)
		return err
	}
}
