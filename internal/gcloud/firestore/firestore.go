package firestore

import (
	"context"
	"fmt"

	"cloud.google.com/go/firestore"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

// New returns firestore client.
func New(lc fx.Lifecycle, log *zap.Logger) (*firestore.Client, error) {
	l := log.Named("firestore")

	client, err := firestore.NewClient(context.Background(), "")
	if err != nil {
		return nil, fmt.Errorf("couldn't connect firestore: %w", err)
	}

	lc.Append(fx.Hook{
		OnStart: func(_ context.Context) error {
			l.Debug("connected")

			return nil
		},
		OnStop: func(_ context.Context) error {
			l.Debug("disconnected")

			return client.Close()
		},
	})

	return client, nil
}
