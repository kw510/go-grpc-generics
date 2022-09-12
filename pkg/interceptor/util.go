package interceptor

import (
	"context"

	"google.golang.org/grpc"
)

type ServerWrapper struct {
	grpc.ServerStream
	Ctx context.Context
}

func (w ServerWrapper) Context() context.Context {
	return w.Ctx
}
