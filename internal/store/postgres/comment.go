package postgres

import (
	"context"
	"database/sql"
	"go-crud-api/internal/model"
)

type CommentStore struct {
	DB *sql.DB
}

func NewCommentStore(db *sql.DB) *CommentStore {
	return &CommentStore{DB: db}
}

func (s *CommentStore) Create(ctx context.Context, c *model.Comment) error {
	query := `INSERT INTO comments (post_id, user_id, content, created_at) VALUES ($1, $2, $3, NOW()) RETURNING id, created_at`
	return s.DB.QueryRowContext(ctx, query, c.PostID, c.UserID, c.Content).Scan(&c.ID, &c.CreatedAt)
}

func (s *CommentStore) GetByPostID(ctx context.Context, postID int) ([]model.Comment, error) {
	rows, err := s.DB.QueryContext(ctx, `SELECT id, post_id, user_id, content, created_at FROM comments WHERE post_id = $1 ORDER BY created_at`, postID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var comments []model.Comment
	for rows.Next() {
		var c model.Comment
		if err := rows.Scan(&c.ID, &c.PostID, &c.UserID, &c.Content, &c.CreatedAt); err != nil {
			return nil, err
		}
		comments = append(comments, c)
	}
	return comments, nil
}

func (s *CommentStore) Delete(ctx context.Context, id, userID int) error {
	_, err := s.DB.ExecContext(ctx, "DELETE FROM comments WHERE id = $1 AND user_id = $2", id, userID)
	return err
}
