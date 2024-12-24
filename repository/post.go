package repository

import (
	"context"
	"goblog/model"

	"gorm.io/gorm"
)

type PostRepository interface {
	Create(ctx context.Context, post model.Post) (model.Post, error)
	GetAll(ctx context.Context) ([]model.Post, error)
	GetByID(ctx context.Context, id uint) (model.Post, error)
	Update(ctx context.Context, post model.Post) (model.Post, error)
	UpdateUserPostCount(ctx context.Context, id uint) error
	Delete(ctx context.Context, id uint) error
}

type PostgresPostRespository struct {
	Db *gorm.DB
}

func (r *PostgresPostRespository) Create(ctx context.Context, post model.Post) (model.Post, error) {
	if err := r.Db.WithContext(ctx).Create(&post).Error; err != nil {
		return model.Post{}, err
	}
	return post, nil
}

func (r *PostgresPostRespository) GetAll(ctx context.Context) ([]model.Post, error) {
	var posts []model.Post
	if err := r.Db.WithContext(ctx).Find(&posts).Error; err != nil {
		return nil, err
	}
	return posts, nil
}

func (r *PostgresPostRespository) GetByID(ctx context.Context, id uint) (model.Post, error) {
	var post model.Post
	if err := r.Db.WithContext(ctx).First(&post, id).Error; err != nil {
		return model.Post{}, err
	}
	return post, nil
}

func (r *PostgresPostRespository) Update(ctx context.Context, post model.Post) (model.Post, error) {
	if err := r.Db.WithContext(ctx).Save(&post).Error; err != nil {
		return model.Post{}, err
	}
	return post, nil
}

func (r *PostgresPostRespository) Delete(ctx context.Context, id uint) error {
	if err := r.Db.WithContext(ctx).Delete(&model.Post{}, id).Error; err != nil {
		return err
	}
	return nil
}

func (r *PostgresPostRespository) UpdateUserPostCount(ctx context.Context, userID uint) error {
	return r.Db.WithContext(ctx).Model(&model.User{}).
		Where("id = ?", userID).
		Update("post", gorm.Expr("post + ?", 1)).Error
}
