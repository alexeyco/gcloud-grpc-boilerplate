package message

import (
	"context"

	"github.com/micromaniacs/gcloud-grpc-boilerplate/internal/models"
)

// Repository messages repository.
type Repository interface {
	GetByID(ctx context.Context, input GetByIDInput) (*models.Message, error)
	Create(ctx context.Context, input CreateInput) (*models.Message, error)
}
