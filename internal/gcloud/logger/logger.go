package logger

import (
	"context"
	"fmt"

	"cloud.google.com/go/logging"
	"github.com/jonstaryuk/gcloudzap"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

// New returns new logger.
func New(lc fx.Lifecycle) (*zap.Logger, error) {
	c := zap.NewProductionConfig()

	// TODO: use Google Cloud Project ID
	client, err := logging.NewClient(context.Background(), "")
	if err != nil {
		return nil, fmt.Errorf("couldn't initialize google cloud logging client: %w", err)
	}

	// nolint:exhaustivestruct
	lc.Append(fx.Hook{
		OnStop: func(_ context.Context) error {
			return client.Close()
		},
	})

	// TODO: use Google Cloud Service Name
	return gcloudzap.New(c, client, "")
}
