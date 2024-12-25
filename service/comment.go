package service

import (
	"context"
	"goblog/model"
	"goblog/repository"
)

type CommentService interface {
	Create(ctx context.Context, comment model.Comment) (model.Comment, error)
	GetAll(ctx context.Context) ([]model.Comment, error)
	GetsByPostID(ctx context.Context, postID int) ([]model.Comment, error)
	Update(ctx context.Context, comment model.Comment) (model.Comment, error)
	Delete(ctx context.Context, comment model.Comment) error
}

type CommentServiceImpl struct {
	Repo repository.CommentRepository
}

func (s *CommentServiceImpl) Create(ctx context.Context, comment model.Comment) (model.Comment, error) {
	return s.Repo.Create(ctx, comment)
}

func (s *CommentServiceImpl) GetAll(ctx context.Context) ([]model.Comment, error) {
	return s.Repo.GetAll(ctx)
}

func (s *CommentServiceImpl) GetsByPostID(ctx context.Context, postID int) ([]model.Comment, error) {
	return s.Repo.GetsByPostID(ctx, postID)
}

func (s *CommentServiceImpl) Update(ctx context.Context, comment model.Comment) (model.Comment, error) {
	return s.Repo.Update(ctx, comment)
}

func (s *CommentServiceImpl) Delete(ctx context.Context, comment model.Comment) error {
	return s.Repo.Delete(ctx, comment)
}
