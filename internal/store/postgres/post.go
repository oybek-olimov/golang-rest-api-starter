package postgres

import (
	"context"
	"database/sql"
	"go-crud-api/internal/model"
)

type PostStore struct {
	DB *sql.DB
}

func NewPostStore(db *sql.DB) *PostStore {
	return &PostStore{DB: db}
}

func (s *PostStore) Create(ctx context.Context, post *model.Post) error {
	query := `INSERT INTO posts (user_id, title, content, created_at) VALUES ($1, $2, $3, NOW()) RETURNING id, created_at`
	return s.DB.QueryRowContext(ctx, query, post.UserID, post.Title, post.Content).Scan(&post.ID, &post.CreatedAt)
}

func (s *PostStore) GetAll(ctx context.Context) ([]model.Post, error) {
	rows, err := s.DB.QueryContext(ctx, "SELECT id, user_id, title, content, created_at FROM posts ORDER BY created_at DESC")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var posts []model.Post
	for rows.Next() {
		var post model.Post
		if err := rows.Scan(&post.ID, &post.UserID, &post.Title, &post.Content, &post.CreatedAt); err != nil {
			return nil, err
		}
		posts = append(posts, post)
	}
	return posts, nil
}

func (s *PostStore) GetByID(ctx context.Context, id int) (*model.Post, error) {
	var post model.Post
	query := `SELECT id, user_id, title, content, created_at FROM posts WHERE id = $1`
	err := s.DB.QueryRowContext(ctx, query, id).Scan(&post.ID, &post.UserID, &post.Title, &post.Content, &post.CreatedAt)
	if err != nil {
		return nil, err
	}
	return &post, nil
}

func (s *PostStore) Update(ctx context.Context, post *model.Post) error {
	query := `UPDATE posts SET title = $1, content = $2 WHERE id = $3 AND user_id = $4`
	_, err := s.DB.ExecContext(ctx, query, post.Title, post.Content, post.ID, post.UserID)
	return err
}

func (s *PostStore) Delete(ctx context.Context, id int, userID int) error {
	_, err := s.DB.ExecContext(ctx, "DELETE FROM posts WHERE id = $1 AND user_id = $2", id, userID)
	return err
}
