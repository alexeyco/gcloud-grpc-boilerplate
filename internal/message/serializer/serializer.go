package serializer

import (
	service "github.com/micromaniacs/gcloud-grpc-boilerplate"
	"github.com/micromaniacs/gcloud-grpc-boilerplate/internal/message"
	"github.com/micromaniacs/gcloud-grpc-boilerplate/internal/models"
	"go.uber.org/fx"
)

type serializer struct {
}

func (s *serializer) ToMessage(m *models.Message) *service.Message {
	panic("implement me")
}

func (s *serializer) ToMessages(messages []*models.Message) []*service.Message {
	panic("implement me")
}

// nolint:stylecheck,golint,gochecknoglobals
var Module = fx.Provide(
	New,
)

// New returns new serializer.
func New() message.Serializer {
	return &serializer{}
}
