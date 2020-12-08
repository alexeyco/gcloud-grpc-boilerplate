package validator

import (
	service "github.com/micromaniacs/gcloud-grpc-boilerplate"
	"github.com/micromaniacs/gcloud-grpc-boilerplate/internal/message"
	"go.uber.org/fx"
)

type validator struct {
}

func (v *validator) ValidateGetByID(request *service.GetByIDRequest) (message.GetByIDInput, error) {
	panic("implement me")
}

func (v *validator) ValidateCreate(request *service.CreateRequest) (message.CreateInput, error) {
	panic("implement me")
}

// nolint:stylecheck,golint,gochecknoglobals
var Module = fx.Provide(
	New,
)

// New returns new validator.
func New() message.Validator {
	return &validator{}
}
