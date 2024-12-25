package service

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"goblog/model"
	"goblog/repository"

	"github.com/go-redis/redis/v8"
)

type PostService interface {
	CreatePostWithUpdateUser(ctx context.Context, post model.Post) (model.Post, error)
	GetAll(ctx context.Context) ([]model.Post, error)
	GetAllAdmin(ctx context.Context, author int) ([]model.Post, error)
	GetByID(ctx context.Context, id int) (model.Post, error)
	Update(ctx context.Context, post model.Post) (model.Post, error)
	Delete(ctx context.Context, id int, author int) error
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
	createPost, err := s.Repo.Create(ctx, tx, post)
	if err != nil {
		tx.Rollback()
		return model.Post{}, err
	}
	if err := s.Repo.UpdateUserPostCount(ctx, tx, post.Author); err != nil {
		tx.Rollback()
		return model.Post{}, err
	}
	if err := tx.Commit().Error; err != nil {
		return model.Post{}, err
	}

	return createPost, nil
}

func (s *PostServiceImpl) GetAll(ctx context.Context) ([]model.Post, error) {
	return s.Repo.GetAll(ctx)
}

func (s *PostServiceImpl) GetAllAdmin(ctx context.Context, author int) ([]model.Post, error) {
	return s.Repo.GetAllAdmin(ctx, author)
}

func (s *PostServiceImpl) GetByID(ctx context.Context, id int) (model.Post, error) {
	cacheKey := fmt.Sprintf("blogpost:%d", id)
	cachedPost, err := s.RedisClient.Get(ctx, cacheKey).Result()
	if err == nil {
		fmt.Println("From cache....")
		var post model.Post
		if json.Unmarshal([]byte(cachedPost), &post) == nil {
			return post, nil
		}
	}
	post, err := s.Repo.GetByID(ctx, id)
	if err != nil {
		return model.Post{}, err
	}
	fmt.Println("From Database....")
	postBytes, _ := json.Marshal(post)
	s.RedisClient.Set(ctx, cacheKey, postBytes, 0)
	return post, nil
}

func (s *PostServiceImpl) Update(ctx context.Context, post model.Post) (model.Post, error) {

	cacheKey := fmt.Sprintf("blogpost:%d", post.ID)
	_, err := s.RedisClient.Get(ctx, cacheKey).Result()
	if err == nil {
		if delErr := s.RedisClient.Del(ctx, cacheKey).Err(); delErr != nil {
			return model.Post{}, errors.New("error delete cache")
		}
	}

	return s.Repo.Update(ctx, post)
}

func (s *PostServiceImpl) Delete(ctx context.Context, id, author int) error {

	cacheKey := fmt.Sprintf("blogpost:%d", id)
	_, err := s.RedisClient.Get(ctx, cacheKey).Result()
	if err == nil {
		if delErr := s.RedisClient.Del(ctx, cacheKey).Err(); delErr != nil {
			return errors.New("error delete cache")
		}
	}

	return s.Repo.Delete(ctx, id, author)
}
