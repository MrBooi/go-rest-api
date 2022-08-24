package comment

import (
	"context"
	"errors"
	"fmt"
)

var (
	ErrCreatingComment = errors.New("failed to create comment")
	ErrUpdatingComment = errors.New("could not update comment")
	ErrNoCommentFound  = errors.New("no comment found")
	ErrDeletingComment = errors.New("could not delete comment")
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
type StoreComment interface { // repository layer
	GetComment(ctx context.Context, id string) (Comment, error)
	UpdateComment(ctx context.Context, id string, cmt Comment) (Comment, error)
	DeleteComment(ctx context.Context, id string) error
	PostComment(ctx context.Context, cmt Comment) (Comment, error)
}

// service -  is the struct on which all our
// logic will be applied/ be build on top of
type Service struct {
	Store StoreComment
}

// NewService - returns a pointer to a new
// service
func NewService(store StoreComment) *Service {
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

func (s *Service) UpdateComment(ctx context.Context, id string, cmt Comment) (Comment, error) {
	cmt, err := s.Store.UpdateComment(ctx, id, cmt)
	if err != nil {
		fmt.Println(err)
		return Comment{}, ErrUpdatingComment
	}
	return cmt, nil
}

func (s *Service) DeleteComment(ctx context.Context, id string) error {
	return s.Store.DeleteComment(ctx, id)
}

func (s *Service) PostComment(ctx context.Context, cmt Comment) (Comment, error) {
	fmt.Println("posting comment.....")
	cmtResponse, err := s.Store.PostComment(ctx, cmt)
	if err != nil {
		fmt.Println(err)
		return Comment{}, ErrCreatingComment
	}
	fmt.Printf("comment posted %v", cmtResponse.ID)
	return cmtResponse, nil
}
