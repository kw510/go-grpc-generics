package interceptor

import (
	"context"

	"google.golang.org/grpc"
)

type serverWrapper struct {
	grpc.ServerStream
	Ctx context.Context
}

func (w serverWrapper) Context() context.Context {
	return w.Ctx
}
