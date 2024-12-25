package repository

import (
	"context"
	"errors"
	"goblog/model"

	"gorm.io/gorm"
)

type PostRepository interface {
	Create(ctx context.Context, tx *gorm.DB, post model.Post) (model.Post, error)
	GetAll(ctx context.Context) ([]model.Post, error)
	GetAllAdmin(ctx context.Context, author int) ([]model.Post, error)
	GetByID(ctx context.Context, id int) (model.Post, error)
	Update(ctx context.Context, post model.Post) (model.Post, error)
	UpdateUserPostCount(ctx context.Context, tx *gorm.DB, id int) error
	Delete(ctx context.Context, id int, author int) error
}

type PostgresPostRespository struct {
	Db *gorm.DB
}

func (r *PostgresPostRespository) Create(ctx context.Context, tx *gorm.DB, post model.Post) (model.Post, error) {
	if err := tx.WithContext(ctx).Create(&post).Error; err != nil {
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

func (r *PostgresPostRespository) GetAllAdmin(ctx context.Context, author int) ([]model.Post, error) {
	var posts []model.Post
	if err := r.Db.WithContext(ctx).Where("author=?", author).Find(&posts).Error; err != nil {
		return nil, err
	}
	return posts, nil
}

func (r *PostgresPostRespository) GetByID(ctx context.Context, id int) (model.Post, error) {
	var post model.Post
	if err := r.Db.WithContext(ctx).First(&post, id).Error; err != nil {
		return model.Post{}, err
	}
	return post, nil
}

func (r *PostgresPostRespository) Update(ctx context.Context, post model.Post) (model.Post, error) {
	var p model.Post
	if err := r.Db.WithContext(ctx).Where("id=?", post.ID).Find(&p).Error; err != nil {
		return model.Post{}, err
	}
	if p.ID == 0 {
		return model.Post{}, errors.New("post not fount")
	}
	if p.Author != post.Author {
		return model.Post{}, errors.New("no authorization")
	}
	if err := r.Db.WithContext(ctx).Save(&post).Error; err != nil {
		return model.Post{}, err
	}
	return post, nil
}

func (r *PostgresPostRespository) Delete(ctx context.Context, id, author int) error {
	var p model.Post
	if err := r.Db.WithContext(ctx).Where("id=?", id).Find(&p).Error; err != nil {
		return err
	}
	if p.ID == 0 {
		return errors.New("post not fount")
	}
	if author != p.Author {
		return errors.New("no authorization")
	}
	if err := r.Db.WithContext(ctx).Delete(&model.Post{}, id).Error; err != nil {
		return err
	}
	return nil
}

func (r *PostgresPostRespository) UpdateUserPostCount(ctx context.Context, tx *gorm.DB, userID int) error {
	return tx.WithContext(ctx).Model(&model.User{}).Where("id=?", userID).Update("post_count", gorm.Expr("post_count + ?", 1)).Error
}
