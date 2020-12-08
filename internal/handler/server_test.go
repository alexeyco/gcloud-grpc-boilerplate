package handler_test

import (
	"context"
	"testing"

	"github.com/golang/mock/gomock"
	service "github.com/micromaniacs/gcloud-grpc-boilerplate"
	"github.com/micromaniacs/gcloud-grpc-boilerplate/internal/handler"
	"github.com/micromaniacs/gcloud-grpc-boilerplate/internal/message"
	"github.com/micromaniacs/gcloud-grpc-boilerplate/internal/message/mocks"
	"github.com/micromaniacs/gcloud-grpc-boilerplate/internal/models"
	"github.com/stretchr/testify/assert"
)

func TestMessagesServer_Create(t *testing.T) {
	ctrl := gomock.NewController(t)
	t.Cleanup(func() {
		ctrl.Finish()
	})

	t.Parallel()

	repository := mocks.NewMockRepository(ctrl)
	serializer := mocks.NewMockSerializer(ctrl)
	validator := mocks.NewMockValidator(ctrl)

	t.Run("Success", func(t *testing.T) {
		t.Parallel()

		request := &service.CreateRequest{
			Text: "text",
		}

		repository.EXPECT().Create(gomock.Any(), gomock.Any()).Return(&models.Message{}, nil)
		serializer.EXPECT().ToMessage(gomock.Any()).Return(&service.Message{})
		validator.EXPECT().ValidateCreate(gomock.Any()).Return(message.CreateInput{}, nil)

		server := handler.New(repository, serializer, validator)
		msg, err := server.Create(context.Background(), request)

		assert.NoError(t, err)
		assert.Equal(t, &service.Message{}, msg)
	})
}

func TestMessagesServer_GetByID(t *testing.T) {
	t.Parallel()

	ctrl := gomock.NewController(t)

	repository := mocks.NewMockRepository(ctrl)
	serializer := mocks.NewMockSerializer(ctrl)
	validator := mocks.NewMockValidator(ctrl)

	t.Run("Success", func(t *testing.T) {
		t.Parallel()

		request := &service.GetByIDRequest{
			Id: "ID",
		}

		repository.EXPECT().GetByID(gomock.Any(), gomock.Any()).Return(&models.Message{}, nil)
		serializer.EXPECT().ToMessage(gomock.Any()).Return(&service.Message{})
		validator.EXPECT().ValidateGetByID(gomock.Any()).Return(message.GetByIDInput{}, nil)

		server := handler.New(repository, serializer, validator)
		msg, err := server.GetByID(context.Background(), request)

		assert.NoError(t, err)
		assert.Equal(t, &service.Message{}, msg)
	})

	ctrl.Finish()
}
