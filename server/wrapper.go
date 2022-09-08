package server

import (
	"context"

	"google.golang.org/grpc"
)

type wrapper struct {
	grpc.ServerStream
	Ctx context.Context
}

func (w wrapper) Context() context.Context {
	return w.Ctx
}
