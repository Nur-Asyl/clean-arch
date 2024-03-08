package article

import (
	"architecture_go/services/article/internal/domain/article"
	"architecture_go/services/article/internal/useCase/adapters/storage"
	"context"
	"errors"
)

type ArticleUseCase struct {
	articleRepo storage.Article
}

func NewArticleUseCase(articleRepo storage.Article) *ArticleUseCase {
	return &ArticleUseCase{
		articleRepo: articleRepo,
	}
}

func (uc *ArticleUseCase) CreateArticle(ctx context.Context, name, text string) (*article.Article, error) {
	newArticle, err := article.NewArticle(name, text)

	err = uc.articleRepo.CreateArticle(ctx, newArticle)
	if err != nil {
		return nil, err
	}

	return newArticle, nil
}

func (uc *ArticleUseCase) ReadArticle(ctx context.Context, name string) (*article.Article, error) {
	return uc.articleRepo.ReadArticle(ctx, name)
}

func (uc *ArticleUseCase) UpdateArticle(ctx context.Context, name, text string) error {
	existingArticle, err := uc.articleRepo.ReadArticle(ctx, name)
	if err != nil {
		return err
	}

	if existingArticle == nil {
		return errors.New("article not found")
	}

	existingArticle.Name = name
	existingArticle.Text = text

	return uc.articleRepo.UpdateArticle(ctx, existingArticle)
}

func (uc *ArticleUseCase) DeleteArticle(ctx context.Context, name string) error {
	return uc.articleRepo.DeleteArticle(ctx, name)
}
