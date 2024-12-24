package repository

import (
	"context"
	"goblog/model"

	"gorm.io/gorm"
)

type UserRepository interface {
	Create(ctx context.Context, post model.User) (model.User, error)
}

type PostgresUserRespository struct {
	Db *gorm.DB
}

func (r *PostgresUserRespository) Create(ctx context.Context, user model.User) (model.User, error) {
	if err := r.Db.WithContext(ctx).Create(&user).Error; err != nil {
		return model.User{}, err
	}
	return user, nil
}
