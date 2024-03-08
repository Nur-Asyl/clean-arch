package article

import (
	"architecture_go/services/article/internal/domain/article"
	"context"
	"database/sql"
	"errors"
)

type ArticleRepository struct {
	db *sql.DB
}

func NewArticleRepository(db *sql.DB) *ArticleRepository {
	return &ArticleRepository{db: db}
}

func (r *ArticleRepository) CreateArticle(ctx context.Context, article *article.Article) error {
	_, err := r.db.ExecContext(ctx, "INSERT INTO articles (name, text) VALUES ($1, $2)", article.Name, article.Text)
	if err != nil {
		return err
	}
	return nil
}

func (r *ArticleRepository) ReadArticle(ctx context.Context, name string) (*article.Article, error) {
	var article article.Article
	err := r.db.QueryRowContext(ctx, "SELECT id, name, text FROM articles WHERE name = $1", name).Scan(&article.ID, &article.Name, &article.Text)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errors.New("article not found")
		}
		return nil, err
	}
	return &article, nil
}

func (r *ArticleRepository) UpdateArticle(ctx context.Context, newArticle *article.Article) error {
	_, err := r.db.ExecContext(ctx, "UPDATE articles SET name = $1, text = $2 WHERE id = $3", newArticle.Name, newArticle.Text, newArticle.ID)
	if err != nil {
		return err
	}
	return nil
}

func (r *ArticleRepository) DeleteArticle(ctx context.Context, name string) error {
	_, err := r.db.ExecContext(ctx, "DELETE FROM articles WHERE name = $1", name)
	if err != nil {
		return err
	}
	return nil
}
