package main

import (
	"bytes"
	"encoding/json"
	"goblog/config"
	"goblog/handler"
	"goblog/model"
	"goblog/repository"
	"goblog/route"
	"goblog/service"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-redis/redis/v8"
	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func setupTestApp() *fiber.App {
	config.Init()
	db, err := gorm.Open(postgres.Open(config.GetConf().DB), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		log.Fatalf("Failed to connect to PostgreSQL: %v", err)
	}
	rdb := redis.NewClient(&redis.Options{
		Addr:     config.GetConf().Redis,
		Password: "",
		DB:       0,
	})
	defer rdb.Close()
	// migration
	db.AutoMigrate(&model.Post{})
	db.AutoMigrate(&model.User{})
	// repo
	postRepo := &repository.PostgresPostRespository{Db: db}
	postService := &service.PostServiceImpl{Repo: postRepo, RedisClient: rdb}
	postHandler := handler.NewPostHandler(postService)
	userRepo := &repository.PostgresUserRespository{Db: db}
	userService := &service.UserServiceImpl{Repo: userRepo}
	userHandler := handler.NewUserHandler(userService)

	app := fiber.New()
	route.PostRoute(app, postHandler)
	route.UserRoute(app, userHandler)

	return app
}

func TestUserCreate(t *testing.T) {
	app := setupTestApp()
	user := model.User{
		Name: "ahmad20",
		Post: 0,
	}
	body, _ := json.Marshal(user)
	req := httptest.NewRequest(http.MethodPost, "/user", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	resp, _ := app.Test(req)
	assert.Equal(t, http.StatusCreated, resp.StatusCode)
}

func TestPostCreate(t *testing.T) {
	app := setupTestApp()
	post := model.Post{
		Title:   "new post",
		Content: "New post content",
	}
	body, _ := json.Marshal(post)
	req := httptest.NewRequest(http.MethodPost, "/post", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	resp, _ := app.Test(req)
	assert.Equal(t, http.StatusCreated, resp.StatusCode)
}
