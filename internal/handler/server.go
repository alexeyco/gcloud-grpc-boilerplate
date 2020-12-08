package handler

import (
	"context"

	service "github.com/micromaniacs/gcloud-grpc-boilerplate"
	"github.com/micromaniacs/gcloud-grpc-boilerplate/internal/message"
	"go.uber.org/fx"
)

type messagesServer struct {
	repository message.Repository
	serializer message.Serializer
	validator  message.Validator
}

// nolint:wrapcheck
func (s *messagesServer) Create(ctx context.Context, request *service.CreateRequest) (*service.Message, error) {
	input, err := s.validator.ValidateCreate(request)
	if err != nil {
		return nil, err
	}

	m, err := s.repository.Create(ctx, input)
	if err != nil {
		return nil, err
	}

	return s.serializer.ToMessage(m), nil
}

// nolint:wrapcheck
func (s *messagesServer) GetByID(ctx context.Context, request *service.GetByIDRequest) (*service.Message, error) {
	input, err := s.validator.ValidateGetByID(request)
	if err != nil {
		return nil, err
	}

	m, err := s.repository.GetByID(ctx, input)
	if err != nil {
		return nil, err
	}

	return s.serializer.ToMessage(m), nil
}

// nolint:stylecheck,golint,gochecknoglobals
var Module = fx.Provide(
	New,
)

// New returns new service server.
func New(r message.Repository, s message.Serializer, v message.Validator) service.MessagesServer {
	return &messagesServer{
		repository: r,
		serializer: s,
		validator:  v,
	}
}
