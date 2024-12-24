package service

import (
	"context"
	"goblog/model"
	"goblog/repository"
)

type UserService interface {
	Create(ctx context.Context, user model.User) (model.User, error)
	Signup(ctx context.Context, user model.User) (model.User, error)
	Signin(ctx context.Context, user model.User) (model.User, error)
	Google(ctx context.Context, user model.User) (model.User, error)
}

type UserServiceImpl struct {
	Repo repository.UserRepository
}

func (s *UserServiceImpl) Create(ctx context.Context, user model.User) (model.User, error) {
	return s.Repo.Create(ctx, user)
}

func (s *UserServiceImpl) Signup(ctx context.Context, user model.User) (model.User, error) {
	return s.Repo.Signup(ctx, user)
}

func (s *UserServiceImpl) Signin(ctx context.Context, user model.User) (model.User, error) {
	return s.Repo.Signin(ctx, user)
}

func (s *UserServiceImpl) Google(ctx context.Context, user model.User) (model.User, error) {
	return s.Repo.Google(ctx, user)
}
