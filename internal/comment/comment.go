package comment

import (
	"context"
	"errors"
	"fmt"
)

var (
	ErrFetchingComment = errors.New("failed to fetch comment by ID")
	ErrNotImplemented  = errors.New("not Implemented")
)

// Comment - a structure for our comment objects
// for our service
type Comment struct {
	ID     string
	Slug   string
	Body   string
	Author string
}

// Store - this interface defines all of the methods
// that our service needs in order to operate
type Store interface {
	GetComment(context.Context, string) (Comment, error)
	PostComment(context.Context, Comment) (Comment, error)
	DeleteComment(context.Context, string) error
	UpdateComment(context.Context, string, Comment) (Comment, error)
}

// Service - struct on which all our logic will be build on top of
type Service struct {
	Store Store
}

// NewService - returns a pointer to a new
//service
func NewService(store Store) *Service {
	return &Service{
		Store: store,
	}
}

func (s *Service) GetComment(ctx context.Context, id string) (Comment, error) {
	fmt.Println("retrieving a comment")

	cmt, err := s.Store.GetComment(ctx, id)
	if err != nil {
		fmt.Println(err)
		return Comment{}, err
	}
	return cmt, nil
}

func (s *Service) UpdateComment(
	ctx context.Context,
	ID string,
	newComment Comment,
) (Comment, error) {
	cmt, err := s.Store.UpdateComment(ctx, ID, newComment)
	if err != nil {
		fmt.Println("Error updating comment")
		return Comment{}, err
	}
	return cmt, nil
}

func (s *Service) DeleteComment(ctx context.Context, id string) error {
	return s.Store.DeleteComment(ctx, id)
}

func (s *Service) CreateComment(ctx context.Context, cmt Comment) (Comment, error) {
	return Comment{}, ErrNotImplemented
}

func (s *Service) PostComment(ctx context.Context, cmt Comment) (Comment, error) {
	insertedCmt, err := s.Store.PostComment(ctx, cmt)
	if err != nil {
		return Comment{}, err
	}
	return insertedCmt, nil
}
