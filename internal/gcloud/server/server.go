package server

import (
	grpcMiddleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpcZap "github.com/grpc-ecosystem/go-grpc-middleware/logging/zap"
	grpcCtxtags "github.com/grpc-ecosystem/go-grpc-middleware/tags"
	"go.uber.org/fx"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

// New returns new handler server.
func New(lc fx.Lifecycle, log *zap.Logger) (*grpc.Server, error) {
	l := log.Named("handler")

	grpcZap.ReplaceGrpcLoggerV2(l)

	// TODO: auth with gcp

	return grpc.NewServer(
		grpcMiddleware.WithUnaryServerChain(
			grpcCtxtags.UnaryServerInterceptor(
				grpcCtxtags.WithFieldExtractor(grpcCtxtags.CodeGenRequestFieldExtractor),
			),
			grpcZap.UnaryServerInterceptor(l),
		),
		grpcMiddleware.WithStreamServerChain(
			grpcCtxtags.StreamServerInterceptor(
				grpcCtxtags.WithFieldExtractor(grpcCtxtags.CodeGenRequestFieldExtractor),
			),
			grpcZap.StreamServerInterceptor(l),
		),
	), nil
}
