package contact

import (
	"architecture_go/services/article/internal/domain/comment"
	"context"
	"database/sql"
	"errors"
)

type CommentRepository struct {
	db *sql.DB
}

func NewCommentRepository(db *sql.DB) *CommentRepository {
	return &CommentRepository{db: db}
}

func (r *CommentRepository) CreateComment(ctx context.Context, comment *comment.Comment) error {
	_, err := r.db.ExecContext(ctx, "INSERT INTO comments (userID, articleID, text) VALUES ($1, $2, $3)", comment.UserID, comment.ArticleID, comment.Text)
	if err != nil {
		return err
	}
	return nil
}

func (r *CommentRepository) ReadComment(ctx context.Context, userID, articleID int) (*comment.Comment, error) {
	var id int
	var text string

	err := r.db.QueryRowContext(ctx, "SELECT id, userID, articleID, text FROM comments WHERE userID = $1 AND articleID = $2", userID, articleID).Scan(&id, &userID, &articleID, &text)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errors.New("comment not found")
		}
		return nil, err
	}
	return comment.NewComment(userID, articleID, text), nil
}
