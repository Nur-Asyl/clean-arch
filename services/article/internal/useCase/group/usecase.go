package group

import (
	"architecture_go/services/article/internal/domain/comment"
	"architecture_go/services/article/internal/useCase"
	"architecture_go/services/article/internal/useCase/adapters/storage"
	"context"
)

type CommentUseCase struct {
	commentRepo storage.Comment
	//articleRepo   storage.Article
}

func NewCommentUseCase(commentRepo storage.Comment) useCase.CommentUseCase {
	return &CommentUseCase{commentRepo: commentRepo}
}

func (uc *CommentUseCase) CreateComment(ctx context.Context, userID, articleID int, text string) (*comment.Comment, error) {
	newComment := comment.NewComment(userID, articleID, text)

	err := uc.commentRepo.CreateComment(ctx, newComment)
	if err != nil {
		return nil, err
	}

	return newComment, nil
}

func (uc *CommentUseCase) ReadComment(ctx context.Context, userID int, articleID int) (*comment.Comment, error) {
	return uc.commentRepo.ReadComment(ctx, userID, articleID)
}

//
//func (uc *CommentUseCase) AddContactToGroup(ctx context.Context, groupID, contactID int) error {
//	return uc.groupRepo.AddContactToGroup(ctx, groupID, contactID)
//}
