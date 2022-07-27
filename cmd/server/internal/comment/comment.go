package comment

import (
	"context"
	"fmt"
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
		return Comment{}, err
	}

	return cmt, nil
}
