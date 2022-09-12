# Go gRPC Generics
[![CI](https://github.com/kw510/go-grpc-generics/actions/workflows/ci.yml/badge.svg?branch=main)](https://github.com/kw510/go-grpc-generics/actions/workflows/ci.yml)
[![codecov](https://codecov.io/gh/kw510/go-grpc-generics/branch/main/graph/badge.svg?token=OD8ANI3KDK)](https://codecov.io/gh/kw510/go-grpc-generics)
[![Go Reference](https://pkg.go.dev/badge/github.com/kw510/go-grpc-generics.svg)](https://pkg.go.dev/github.com/kw510/go-grpc-generics)
[![Go Report Card](https://goreportcard.com/badge/github.com/kw510/go-grpc-generics)](https://goreportcard.com/report/github.com/kw510/go-grpc-generics)


A generic and uniform interceptor, combining unary and stream gRPC interceptors into a single interceptor.

Just define the interceptor once, then covert into the type gRPC interceptor that you need! 🪄

Insipred by https://github.com/grpc-ecosystem/go-grpc-middleware.

## Usage
To use this generic interceptor, first define it, then convert it into a grpc interceptor.

### Define an generic Interceptor

We define what the interceptor perfoms before and after the call. This implements the generic interceptor interface.

```go
type YourInterceptor struct {}

func (i YourInterceptor) BeforeHandler(ctx context.Context) context.Context {
  // Performed before the handler is called
  ...
  return ctx // Context is passed into the handler
}

func (i YourInterceptor) AfterHandler(ctx context.Context, err error) {
  // Performed after the handler is called
  ...
}
```

### Using an Interceptor

Pass in the struct to the converter functions, which will return the respective interceptor.

```go
import "github.com/kw510/grpc-interceptor/interceptors"

grpcServer := grpc.NewServer(
  grpc.StreamInterceptor(
    interceptors.StreamServerInterceptor(YourInterceptor{})
  )
  grpc.UnaryInterceptor(
    interceptors.UnaryServerInterceptor(YourInterceptor{})
  )
)
```
