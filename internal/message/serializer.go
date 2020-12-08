package message

import (
	service "github.com/micromaniacs/gcloud-grpc-boilerplate"
	"github.com/micromaniacs/gcloud-grpc-boilerplate/internal/models"
)

// Serializer messages serializer.
type Serializer interface {
	ToMessage(*models.Message) *service.Message
	ToMessages([]*models.Message) []*service.Message
}
