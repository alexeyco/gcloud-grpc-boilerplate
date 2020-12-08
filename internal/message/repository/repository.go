package repository

import (
	"context"

	"cloud.google.com/go/firestore"
	"github.com/micromaniacs/gcloud-grpc-boilerplate/internal/message"
	"github.com/micromaniacs/gcloud-grpc-boilerplate/internal/models"
	"go.uber.org/fx"
)

type repository struct {
	firestore *firestore.Client
}

func (r *repository) GetByID(ctx context.Context, input message.GetByIDInput) (*models.Message, error) {
	panic("implement me")
}

func (r *repository) Create(ctx context.Context, input message.CreateInput) (*models.Message, error) {
	panic("implement me")
}

// nolint:stylecheck,golint,gochecknoglobals
var Module = fx.Provide(
	New,
)

// New returns new messages repository.
func New(client *firestore.Client) message.Repository {
	return &repository{
		firestore: client,
	}
}
