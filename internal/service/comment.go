package service

import (
	"context"
	"go-crud-api/internal/model"
)

type CommentStore interface {
	Create(context.Context, *model.Comment) error
	GetByPostID(context.Context, int) ([]model.Comment, error)
	Delete(context.Context, int, int) error
}

type CommentService struct {
	store CommentStore
}

func NewCommentService(store CommentStore) *CommentService {
	return &CommentService{store: store}
}

func (s *CommentService) Create(ctx context.Context, c *model.Comment) error {
	return s.store.Create(ctx, c)
}

func (s *CommentService) GetByPostID(ctx context.Context, postID int) ([]model.Comment, error) {
	return s.store.GetByPostID(ctx, postID)
}

func (s *CommentService) Delete(ctx context.Context, id, userID int) error {
	return s.store.Delete(ctx, id, userID)
}
