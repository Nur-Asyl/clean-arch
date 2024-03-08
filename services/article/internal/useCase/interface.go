package useCase

import (
	"architecture_go/services/article/internal/domain/article"
	"architecture_go/services/article/internal/domain/comment"
	"context"
)

type ArticleUseCase interface {
	CreateArticle(ctx context.Context, name, text string) (*article.Article, error)
	ReadArticle(ctx context.Context, name string) (*article.Article, error)
	UpdateArticle(ctx context.Context, name string, text string) error
	DeleteArticle(ctx context.Context, name string) error
}

type CommentUseCase interface {
	CreateComment(ctx context.Context, userID, articleID int, text string) (*comment.Comment, error)
	ReadComment(ctx context.Context, userID, articleID int) (*comment.Comment, error)
	//AddContactToGroup(ctx context.Context, groupID, contactID int) error
}
