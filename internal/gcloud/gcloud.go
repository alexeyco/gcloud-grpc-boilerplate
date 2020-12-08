package gcloud

import (
	"github.com/micromaniacs/gcloud-grpc-boilerplate/internal/gcloud/firestore"
	"github.com/micromaniacs/gcloud-grpc-boilerplate/internal/gcloud/logger"
	"github.com/micromaniacs/gcloud-grpc-boilerplate/internal/gcloud/server"
	"go.uber.org/fx"
)

// nolint:stylecheck,golint,gochecknoglobals
var Module = fx.Provide(
	firestore.New,
	logger.New,
	server.New,
)
