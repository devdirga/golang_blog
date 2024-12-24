package service

import (
	"context"
	"goblog/model"
	"goblog/repository"
)

type UserService interface {
	Create(ctx context.Context, user model.User) (model.User, error)
}

type UserServiceImpl struct {
	Repo repository.UserRepository
}

func (s *UserServiceImpl) Create(ctx context.Context, user model.User) (model.User, error) {
	return s.Repo.Create(ctx, user)
}
