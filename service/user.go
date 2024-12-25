package service

import (
	"context"
	"goblog/model"
	"goblog/repository"
)

type UserService interface {
	Signup(ctx context.Context, user model.User) error
	Signin(ctx context.Context, user model.User) (model.User, error)
	Google(ctx context.Context, user model.User) (model.User, error)
	Me(ctx context.Context, user model.User) (model.User, error)
	UpdateProfile(ctx context.Context, user model.User) error
}

type UserServiceImpl struct {
	Repo repository.UserRepository
}

func (s *UserServiceImpl) Signup(ctx context.Context, user model.User) error {
	return s.Repo.Signup(ctx, user)
}

func (s *UserServiceImpl) Signin(ctx context.Context, user model.User) (model.User, error) {
	return s.Repo.Signin(ctx, user)
}

func (s *UserServiceImpl) Google(ctx context.Context, user model.User) (model.User, error) {
	return s.Repo.Google(ctx, user)
}

func (s *UserServiceImpl) Me(ctx context.Context, user model.User) (model.User, error) {
	return s.Repo.Me(ctx, user)
}

func (s *UserServiceImpl) UpdateProfile(ctx context.Context, user model.User) error {
	return s.Repo.UpdateProfile(ctx, user)
}
