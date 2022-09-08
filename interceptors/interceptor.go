package interceptors

import "context"

type Interceptor interface {
	BeforeHandler(ctx context.Context) context.Context
	AfterHandler(ctx context.Context, err error)
}
