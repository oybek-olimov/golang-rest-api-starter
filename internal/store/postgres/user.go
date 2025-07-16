package postgres

import (
	"context"
	"database/sql"
	"errors"

	"go-crud-api/internal/model"
)

type UserStore struct {
	DB *sql.DB
}

func NewUserStore(db *sql.DB) *UserStore {
	return &UserStore{DB: db}
}

func (s *UserStore) Create(ctx context.Context, u *model.User) error {
	query := `
		INSERT INTO users (username, email, password)
		VALUES ($1, $2, $3)
		RETURNING id, created_at;
	`
	return s.DB.QueryRowContext(ctx, query, u.Username, u.Email, u.Password).
		Scan(&u.ID, &u.CreatedAt)
}

func (s *UserStore) GetByEmail(ctx context.Context, email string) (*model.User, error) {
	query := `
		SELECT id, username, email, password, created_at
		FROM users
		WHERE email = $1;
	`

	var u model.User
	err := s.DB.QueryRowContext(ctx, query, email).
		Scan(&u.ID, &u.Username, &u.Email, &u.Password, &u.CreatedAt)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}

	return &u, nil
}
