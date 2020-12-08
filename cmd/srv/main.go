package main

import (
	"context"
	"fmt"
	"net"

	service "github.com/micromaniacs/gcloud-grpc-boilerplate"
	"github.com/micromaniacs/gcloud-grpc-boilerplate/internal/gcloud"
	"github.com/micromaniacs/gcloud-grpc-boilerplate/internal/handler"
	"github.com/micromaniacs/gcloud-grpc-boilerplate/internal/message/repository"
	"github.com/micromaniacs/gcloud-grpc-boilerplate/internal/message/serializer"
	"github.com/micromaniacs/gcloud-grpc-boilerplate/internal/message/validator"
	"go.uber.org/fx"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

func run(lc fx.Lifecycle, log *zap.Logger, s *grpc.Server, srv service.MessagesServer) (err error) {
	port := 8080
	l := log.Named("server")

	// TODO: get app port

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		return fmt.Errorf("can't listen port %d: %w", port, err)
	}

	service.RegisterMessagesServer(s, srv)

	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			l.Debug("listening", zap.Int("port", port))

			return nil
		},
		OnStop: func(ctx context.Context) error {
			s.GracefulStop()
			l.Debug("stopped listening", zap.Int("port", port))

			return nil
		},
	})

	go func() {
		err = s.Serve(lis)
	}()

	return nil
}

func main() {
	fx.New(
		gcloud.Module,
		repository.Module,
		serializer.Module,
		validator.Module,
		handler.Module,
		fx.Invoke(run),
	).Run()
}
