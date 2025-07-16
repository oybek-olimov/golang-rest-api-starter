package service

import (
	"context"
	"go-crud-api/internal/model"
)

type PostStore interface {
	Create(context.Context, *model.Post) error
	GetAll(context.Context) ([]model.Post, error)
	GetByID(context.Context, int) (*model.Post, error)
	Update(context.Context, *model.Post) error
	Delete(context.Context, int, int) error
}

type PostService struct {
	store PostStore
}

func NewPostService(store PostStore) *PostService {
	return &PostService{store: store}
}

func (s *PostService) Create(ctx context.Context, post *model.Post) error {
	return s.store.Create(ctx, post)
}

func (s *PostService) GetAll(ctx context.Context) ([]model.Post, error) {
	return s.store.GetAll(ctx)
}

func (s *PostService) GetByID(ctx context.Context, id int) (*model.Post, error) {
	return s.store.GetByID(ctx, id)
}

func (s *PostService) Update(ctx context.Context, post *model.Post) error {
	return s.store.Update(ctx, post)
}

func (s *PostService) Delete(ctx context.Context, id int, userID int) error {
	return s.store.Delete(ctx, id, userID)
}
