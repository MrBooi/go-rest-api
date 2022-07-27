package comment

import (
	"context"
	"errors"
	"fmt"
)

var (
	ErrFetchingComment = errors.New("error fetching comment by id")
	ErrNotImplemented  = errors.New("not implemented")
)

// comment - a representation of a comment
// structure for our service
type Comment struct {
	ID     string
	Slug   string
	Body   string
	Author string
}

// store - this interface defines all of the methods
// that our service needs in order to operate
type Store interface { // repository layer
	GetComment(ctx context.Context, id string) (Comment, error)
	UpdateComment(ctx context.Context, id string) error
	DeleteComment(ctx context.Context, id string) error
	CreateComment(ctx context.Context, cmt Comment) (Comment, error)
}

// service -  is the struct on which all our
// logic will be applied/ be build on top of
type Service struct {
	Store Store
}

// NewService - returns a pointer to a new
// service
func NewService(store Store) *Service {
	return &Service{
		Store: store,
	}
}

func (s *Service) GetComment(ctx context.Context, id string) (Comment, error) {
	fmt.Println("getting comment")

	cmt, err := s.Store.GetComment(ctx, id)
	if err != nil {
		fmt.Println(err)
		return Comment{}, ErrFetchingComment
	}

	return cmt, nil
}

func (s *Service) UpdateComment(ctx context.Context, cmt Comment) error {
	return ErrNotImplemented
}

func (s *Service) DeleteComment(ctx context.Context, cmt Comment) error {
	return ErrNotImplemented
}

func (s *Service) CreateComment(ctx context.Context, cmt Comment) (Comment, error) {
	return Comment{}, ErrNotImplemented
}
