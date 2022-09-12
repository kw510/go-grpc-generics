package interceptor

import (
	"context"

	"google.golang.org/grpc/metadata"
)

// Context keys
type beforeInterceptor struct{}
type beforeHandlerFunc struct{}

// Implement a generic interceptor
type interceptor struct {
	beforeHandler func(ctx context.Context) context.Context
	afterHandler  func(ctx context.Context, err error)
}

func (i interceptor) BeforeHandler(ctx context.Context) context.Context {
	return i.beforeHandler(ctx)
}

func (i interceptor) AfterHandler(ctx context.Context, err error) {
	i.afterHandler(ctx, err)
}

// Implement grpc.ServerStream
type serverStream struct {
	ctx context.Context
}

func (s serverStream) Context() context.Context {
	return s.ctx
}

func (s serverStream) RecvMsg(m interface{}) error {
	return nil
}

func (s serverStream) SendHeader(metadata.MD) error {
	return nil
}

func (s serverStream) SendMsg(m interface{}) error {
	return nil
}

func (s serverStream) SetHeader(metadata.MD) error {
	return nil
}

func (s serverStream) SetTrailer(metadata.MD) {}
