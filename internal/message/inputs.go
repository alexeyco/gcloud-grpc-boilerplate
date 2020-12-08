package message

// GetByIDInput input arguments to get a message by ID.
type GetByIDInput struct {
	ID string
}

// CreateInput input arguments to create a message.
type CreateInput struct {
	Text string
}
