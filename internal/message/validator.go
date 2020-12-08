package message

import service "github.com/micromaniacs/gcloud-grpc-boilerplate"

// Validator messages input validator.
type Validator interface {
	ValidateGetByID(*service.GetByIDRequest) (GetByIDInput, error)
	ValidateCreate(*service.CreateRequest) (CreateInput, error)
}
