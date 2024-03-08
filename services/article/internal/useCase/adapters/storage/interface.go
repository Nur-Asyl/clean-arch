package storage

import (
	"architecture_go/services/article/internal/domain/article"
	"architecture_go/services/article/internal/domain/comment"
	"context"
)

type Article interface {
	CreateArticle(ctx context.Context, article *article.Article) error
	ReadArticle(ctx context.Context, name string) (*article.Article, error)
	UpdateArticle(ctx context.Context, contact *article.Article) error
	DeleteArticle(ctx context.Context, name string) error
}

type Comment interface {
	CreateComment(ctx context.Context, comment *comment.Comment) error
	ReadComment(ctx context.Context, userID, articleID int) (*comment.Comment, error)
	//AddCommentToArticle(ctx context.Context, userID, articleID int) error
}
