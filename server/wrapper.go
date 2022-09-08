package server

import (
	"context"

	"google.golang.org/grpc"
)

type Wrapper struct {
	grpc.ServerStream
	Ctx context.Context
}

func (w Wrapper) Context() context.Context {
	return w.Ctx
}
