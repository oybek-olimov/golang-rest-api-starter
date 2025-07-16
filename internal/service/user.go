package service

import (
	"context"
	"errors"

	"go-crud-api/internal/model"
	"go-crud-api/internal/store/postgres"
	"go-crud-api/pkg/hash"
)

type UserStore interface {
	Create(user *model.User) error
	GetByEmail(email string) (*model.User, error)
}

type Hasher interface {
	Hash(password string) (string, error)
	Compare(hash, password string) bool
}


type UserService struct {
	UserStore *postgres.UserStore
	Hasher    *hash.Hasher
}

func NewUserService(store *postgres.UserStore, hasher *hash.Hasher) *UserService {
	return &UserService{
		UserStore: store,
		Hasher:    hasher,
	}
}

func (s *UserService) Register(ctx context.Context, username, email, password string) (*model.User, error) {
	hashedPassword, err := s.Hasher.HashPassword(password)
	if err != nil {
		return nil, err
	}

	user := &model.User{
		Username: username,
		Email:    email,
		Password: hashedPassword,
	}

	if err := s.UserStore.Create(ctx, user); err != nil {
		return nil, err
	}

	return user, nil
}

func (s *UserService) Login(ctx context.Context, email, password string) (*model.User, error) {
	user, err := s.UserStore.GetByEmail(ctx, email)
	if err != nil {
		return nil, err
	}
	if user == nil {
		return nil, errors.New("user not found")
	}

	if !s.Hasher.CheckPasswordHash(password, user.Password) {
		return nil, errors.New("invalid credentials")
	}

	return user, nil
}
