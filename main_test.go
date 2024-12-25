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
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/stretchr/testify/assert"

	jwtware "github.com/gofiber/contrib/jwt"
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
		Password: config.GetConf().RedisPassword,
		DB:       0,
	})
	defer rdb.Close()

	db.AutoMigrate(&model.User{})
	db.AutoMigrate(&model.Post{})
	db.AutoMigrate(&model.Comment{})

	userRepo := &repository.PostgresUserRespository{Db: db}
	userService := &service.UserServiceImpl{Repo: userRepo}
	userHandler := handler.NewUserHandler(userService)

	postRepo := &repository.PostgresPostRespository{Db: db}
	postService := &service.PostServiceImpl{Repo: postRepo, RedisClient: rdb}
	postHandler := handler.NewPostHandler(postService)

	commentRepo := &repository.PostgresCommentRespository{Db: db}
	commentService := &service.CommentServiceImpl{Repo: commentRepo}
	commentHandler := handler.NewCommentHandler(commentService)

	app := fiber.New()
	app.Use(cors.New(cors.Config{AllowOrigins: "*"}))

	route.UserRoute(app, userHandler)
	route.PublicRoute(app, postHandler)

	app.Use(jwtware.New(jwtware.Config{
		SigningKey: jwtware.SigningKey{Key: []byte(config.GetConf().Secret)},
	}))

	route.InfoRoute(app, userHandler)
	route.PostRoute(app, postHandler)
	route.CommentRoute(app, commentHandler)

	return app
}

func TestSignup(t *testing.T) {
	app := setupTestApp()
	user := model.User{
		Name:     "ahmad",
		Email:    "ahmad@gmail.com",
		Password: "admin123",
	}
	body, _ := json.Marshal(user)
	req := httptest.NewRequest(http.MethodPost, "/signup", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	resp, _ := app.Test(req)
	assert.Equal(t, http.StatusCreated, resp.StatusCode)
}

func TestSignin(t *testing.T) {
	app := setupTestApp()
	post := model.User{
		Email:    "ahmad@gmail.com",
		Password: "admin123",
	}
	body, _ := json.Marshal(post)
	req := httptest.NewRequest(http.MethodPost, "/signin", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	resp, _ := app.Test(req)
	assert.Equal(t, http.StatusCreated, resp.StatusCode)
}

func TestMe(t *testing.T) {
	app := setupTestApp()
	req := httptest.NewRequest(http.MethodGet, "/me", nil)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6ImRpcmdhQGdtYWlsLmNvbSIsImV4cCI6MTc2NjE1Mzk1MiwiaWQiOiIxIiwibmFtZSI6IkRpcmdhIn0.CKkmy5UvlVEOtE-xh9UgfeWVNosBlTf8YX5z8M76VgI")
	resp, _ := app.Test(req)
	assert.Equal(t, http.StatusOK, resp.StatusCode)
}

func TestPostGets(t *testing.T) {
	app := setupTestApp()
	req := httptest.NewRequest(http.MethodGet, "/post", nil)
	req.Header.Set("Content-Type", "application/json")
	resp, _ := app.Test(req)
	assert.Equal(t, http.StatusOK, resp.StatusCode)
}

func TestPostGet(t *testing.T) {
	app := setupTestApp()
	req := httptest.NewRequest(http.MethodPatch, "/post/10", nil)
	req.Header.Set("Content-Type", "application/json")
	resp, _ := app.Test(req)
	assert.Equal(t, http.StatusOK, resp.StatusCode)
}

func TestPostGetsAdmin(t *testing.T) {
	app := setupTestApp()
	req := httptest.NewRequest(http.MethodGet, "/admin/post", nil)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6ImRpcmdhQGdtYWlsLmNvbSIsImV4cCI6MTc2NjE1Mzk1MiwiaWQiOiIxIiwibmFtZSI6IkRpcmdhIn0.CKkmy5UvlVEOtE-xh9UgfeWVNosBlTf8YX5z8M76VgI")
	resp, _ := app.Test(req)
	assert.Equal(t, http.StatusOK, resp.StatusCode)
}

func TestPostCreate(t *testing.T) {
	app := setupTestApp()
	post := model.Post{
		Title:   "Post title",
		Content: "Post Content",
	}
	body, _ := json.Marshal(post)
	req := httptest.NewRequest(http.MethodPost, "/post", bytes.NewReader(body))
	req.Header.Set("Authorization", "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6ImRpcmdhQGdtYWlsLmNvbSIsImV4cCI6MTc2NjE1Mzk1MiwiaWQiOiIxIiwibmFtZSI6IkRpcmdhIn0.CKkmy5UvlVEOtE-xh9UgfeWVNosBlTf8YX5z8M76VgI")
	req.Header.Set("Content-Type", "application/json")
	resp, _ := app.Test(req)
	assert.Equal(t, http.StatusCreated, resp.StatusCode)
}

func TestPostUpdate(t *testing.T) {
	app := setupTestApp()
	post := model.Post{
		Title:   "Post title",
		Content: "Post ContentX",
	}
	body, _ := json.Marshal(post)
	req := httptest.NewRequest(http.MethodPut, "/post/10", bytes.NewReader(body))
	req.Header.Set("Authorization", "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6ImRpcmdhQGdtYWlsLmNvbSIsImV4cCI6MTc2NjE1Mzk1MiwiaWQiOiIxIiwibmFtZSI6IkRpcmdhIn0.CKkmy5UvlVEOtE-xh9UgfeWVNosBlTf8YX5z8M76VgI")
	req.Header.Set("Content-Type", "application/json")
	resp, _ := app.Test(req)
	assert.Equal(t, http.StatusOK, resp.StatusCode)
}

func TestPostDelete(t *testing.T) {
	app := setupTestApp()
	req := httptest.NewRequest(http.MethodDelete, "/post/8", nil)
	req.Header.Set("Authorization", "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6ImRpcmdhQGdtYWlsLmNvbSIsImV4cCI6MTc2NjE1Mzk1MiwiaWQiOiIxIiwibmFtZSI6IkRpcmdhIn0.CKkmy5UvlVEOtE-xh9UgfeWVNosBlTf8YX5z8M76VgI")
	req.Header.Set("Content-Type", "application/json")
	resp, _ := app.Test(req)
	assert.Equal(t, http.StatusNoContent, resp.StatusCode)
}

func TestCommentCreate(t *testing.T) {
	app := setupTestApp()
	post := model.Comment{
		Content: "Comment 1",
		PostID:  2,
	}
	body, _ := json.Marshal(post)
	req := httptest.NewRequest(http.MethodPost, "/comment", bytes.NewReader(body))
	req.Header.Set("Authorization", "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6ImRpcmdhQGdtYWlsLmNvbSIsImV4cCI6MTc2NjIyMTMzMSwiaWQiOiIyIiwibmFtZSI6IkRpcmdhbnRhcmEifQ.4G1O1g63P7OOogx3Mu9OQdmUezTO2x-zUX6PbYA4g10")
	req.Header.Set("Content-Type", "application/json")
	resp, _ := app.Test(req)
	assert.Equal(t, http.StatusCreated, resp.StatusCode)
}

func TestCommentUpdate(t *testing.T) {
	app := setupTestApp()
	post := model.Comment{
		Content: "Comment 1333333",
		PostID:  2,
	}
	body, _ := json.Marshal(post)
	req := httptest.NewRequest(http.MethodPut, "/comment/2", bytes.NewReader(body))
	req.Header.Set("Authorization", "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6ImRpcmdhQGdtYWlsLmNvbSIsImV4cCI6MTc2NjIyMTMzMSwiaWQiOiIyIiwibmFtZSI6IkRpcmdhbnRhcmEifQ.4G1O1g63P7OOogx3Mu9OQdmUezTO2x-zUX6PbYA4g10")
	req.Header.Set("Content-Type", "application/json")
	resp, _ := app.Test(req)
	assert.Equal(t, http.StatusOK, resp.StatusCode)
}
