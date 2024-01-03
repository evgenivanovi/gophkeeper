package grpc

import (
	"context"

	"google.golang.org/grpc"
)

/* __________________________________________________ */

func UnaryClientProtected(
	methods map[string]bool,
	interceptor grpc.UnaryClientInterceptor,
) grpc.UnaryClientInterceptor {
	return func(
		ctx context.Context,
		method string,
		req, reply any,
		cc *grpc.ClientConn,
		invoker grpc.UnaryInvoker,
		opts ...grpc.CallOption,
	) error {
		if methods[method] == true {
			return interceptor(ctx, method, req, reply, cc, invoker, opts...)
		}
		return invoker(ctx, method, req, reply, cc, opts...)
	}
}

func UnaryServerProtected(
	methods map[string]bool,
	interceptor grpc.UnaryServerInterceptor,
) grpc.UnaryServerInterceptor {
	return func(
		ctx context.Context,
		req any,
		info *grpc.UnaryServerInfo,
		handler grpc.UnaryHandler,
	) (resp any, err error) {
		if methods[info.FullMethod] == true {
			return interceptor(ctx, req, info, handler)
		}
		return handler(ctx, req)
	}
}

/* __________________________________________________ */
