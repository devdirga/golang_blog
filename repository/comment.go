package repository

import (
	"context"
	"errors"
	"goblog/model"

	"gorm.io/gorm"
)

type CommentRepository interface {
	Create(ctx context.Context, comment model.Comment) (model.Comment, error)
	GetAll(ctx context.Context) ([]model.Comment, error)
	GetsByPostID(ctx context.Context, id int) ([]model.Comment, error)
	Update(ctx context.Context, comment model.Comment) (model.Comment, error)
	Delete(ctx context.Context, comment model.Comment) error
}

type PostgresCommentRespository struct {
	Db *gorm.DB
}

func (r *PostgresCommentRespository) Create(ctx context.Context, comment model.Comment) (model.Comment, error) {
	if err := r.Db.WithContext(ctx).Create(&comment).Error; err != nil {
		return model.Comment{}, err
	}
	return comment, nil
}

func (r *PostgresCommentRespository) GetAll(ctx context.Context) ([]model.Comment, error) {
	var comments []model.Comment
	if err := r.Db.WithContext(ctx).Find(&comments).Error; err != nil {
		return nil, err
	}
	return comments, nil
}

func (r *PostgresCommentRespository) GetsByPostID(ctx context.Context, id int) ([]model.Comment, error) {
	var comments []model.Comment
	if err := r.Db.WithContext(ctx).Where("post_id=?", id).Find(&comments).Error; err != nil {
		return nil, err
	}
	return comments, nil
}

func (r *PostgresCommentRespository) Update(ctx context.Context, comment model.Comment) (model.Comment, error) {
	var c model.Comment
	if err := r.Db.WithContext(ctx).Where("id=?", comment.ID).Find(&c).Error; err != nil {
		return model.Comment{}, err
	}
	if c.ID == 0 {
		return model.Comment{}, errors.New("post not fount")
	}
	if c.UserID != comment.UserID {
		return model.Comment{}, errors.New("no authorization")
	}
	if err := r.Db.WithContext(ctx).Save(&comment).Error; err != nil {
		return model.Comment{}, err
	}
	return comment, nil
}

func (r *PostgresCommentRespository) Delete(ctx context.Context, comment model.Comment) error {
	var c model.Comment
	if err := r.Db.WithContext(ctx).Where("id=?", comment.ID).Find(&c).Error; err != nil {
		return err
	}
	if c.ID == 0 {
		return errors.New("post not fount")
	}
	if c.UserID != comment.UserID {
		return errors.New("no authorization")
	}
	comment.IsActive = false
	if err := r.Db.WithContext(ctx).Save(&comment).Error; err != nil {
		return err
	}
	return nil
}
