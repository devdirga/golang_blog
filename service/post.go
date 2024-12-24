package service

import (
	"context"
	"encoding/json"
	"fmt"
	"goblog/model"
	"goblog/repository"

	"github.com/go-redis/redis/v8"
)

type PostService interface {
	CreatePostWithUpdateUser(ctx context.Context, post model.Post) (model.Post, error)
	// CreatePost(ctx context.Context, post model.Post) (model.Post, error)
	GetAllPosts(ctx context.Context) ([]model.Post, error)
	GetPostByID(ctx context.Context, id uint) (model.Post, error)
	UpdatePost(ctx context.Context, post model.Post) (model.Post, error)
	DeletePost(ctx context.Context, id uint) error
}

type PostServiceImpl struct {
	Repo        repository.PostRepository
	RedisClient *redis.Client
}

func (s *PostServiceImpl) CreatePostWithUpdateUser(ctx context.Context, post model.Post) (model.Post, error) {
	tx := s.Repo.(*repository.PostgresPostRespository).Db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()
	createPost, err := s.Repo.Create(ctx, post)
	if err != nil {
		tx.Rollback()
		return model.Post{}, err
	}

	if err := s.Repo.UpdateUserPostCount(ctx, 1); err != nil {
		tx.Rollback()
		return model.Post{}, err
	}

	if err := tx.Commit().Error; err != nil {
		return model.Post{}, err
	}

	return createPost, nil
}

// func (s *PostServiceImpl) CreatePost(ctx context.Context, post model.Post) (model.Post, error) {
// 	return s.Repo.Create(ctx, post)
// }

func (s *PostServiceImpl) GetAllPosts(ctx context.Context) ([]model.Post, error) {
	return s.Repo.GetAll(ctx)
}

func (s *PostServiceImpl) GetPostByID(ctx context.Context, id uint) (model.Post, error) {
	cacheKey := fmt.Sprintf("blogpost:%d", id)
	cachedPost, err := s.RedisClient.Get(ctx, cacheKey).Result()
	if err == nil {
		fmt.Print("From cache....")
		var post model.Post
		if json.Unmarshal([]byte(cachedPost), &post) == nil {
			return post, nil
		}
	}

	post, err := s.Repo.GetByID(ctx, id)
	if err != nil {
		return model.Post{}, err
	}
	fmt.Print("From Database....")

	// Cache the post
	postBytes, _ := json.Marshal(post)
	s.RedisClient.Set(ctx, cacheKey, postBytes, 0)

	return post, nil
}

func (s *PostServiceImpl) UpdatePost(ctx context.Context, post model.Post) (model.Post, error) {
	return s.Repo.Update(ctx, post)
}

func (s *PostServiceImpl) DeletePost(ctx context.Context, id uint) error {
	return s.Repo.Delete(ctx, id)
}
