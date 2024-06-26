package service

import (
	"context"

	"github.com/rahmatadlin/Go-Digitalent-2024/Final-Project/internal/model"
	"github.com/rahmatadlin/Go-Digitalent-2024/Final-Project/internal/repository"
)

type CommentService interface {
	PostComment(ctx context.Context, userId uint32, newComment model.CreateComment) (*model.CreateCommentRes, error)
	GetAllCommentsByPhotoId(ctx context.Context, photoId uint32) ([]model.CommentView, error)
	GetCommentById(ctx context.Context, commentId uint32) (*model.Comment, error)
	UpdateComment(ctx context.Context, comment model.Comment) (*model.UpdateCommentRes, error)
	DeleteComment(ctx context.Context, commentId uint32) error
}

type commentServiceImpl struct {
	repo repository.CommentRepository
}

func NewCommentService(repo repository.CommentRepository) CommentService {
	return &commentServiceImpl{repo: repo}
}

func (c *commentServiceImpl) PostComment(ctx context.Context, userId uint32, newComment model.CreateComment) (*model.CreateCommentRes, error) {
	comment := model.Comment{}
	comment.UserId = userId
	comment.Message = newComment.Message
	comment.PhotoId = newComment.PhotoId

	err := c.repo.CreateComment(ctx, &comment)
	if err != nil {
		return nil, err
	}

	commentRes := model.CreateCommentRes{}
	commentRes.ID = comment.ID
	commentRes.Message = comment.Message
	commentRes.PhotoId = comment.PhotoId
	commentRes.UserId = comment.UserId
	commentRes.CreatedAt = comment.CreatedAt

	return &commentRes, nil
}

func (c *commentServiceImpl) GetAllCommentsByPhotoId(ctx context.Context, photoId uint32) ([]model.CommentView, error) {
	comments, err := c.repo.GetAllCommentsByPhotoId(ctx, photoId)
	if err != nil {
		return nil, err
	}

	return comments, nil
}

func (c *commentServiceImpl) GetCommentById(ctx context.Context, commentId uint32) (*model.Comment, error) {
	comment, err := c.repo.GetCommentById(ctx, commentId)
	if err != nil {
		return nil, err
	}

	return comment, nil
}

func (c *commentServiceImpl) UpdateComment(ctx context.Context, comment model.Comment) (*model.UpdateCommentRes, error) {
	err := c.repo.UpdateComment(ctx, &comment)

	if err != nil {
		return nil, err
	}

	commentRes := model.UpdateCommentRes{}
	commentRes.ID = comment.ID
	commentRes.Message = comment.Message
	commentRes.PhotoId = comment.PhotoId
	commentRes.UserId = comment.UserId
	commentRes.UpdatedAt = comment.UpdatedAt

	return &commentRes, nil
}

func (c *commentServiceImpl) DeleteComment(ctx context.Context, commentId uint32) error {
	err := c.repo.DeleteComment(ctx, commentId)

	return err
}