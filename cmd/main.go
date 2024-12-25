package main

import (
	"goblog/config"
	"goblog/handler"
	"goblog/model"
	"goblog/repository"
	"goblog/route"
	"goblog/service"
	"log"

	"github.com/go-redis/redis/v8"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"

	jwtware "github.com/gofiber/contrib/jwt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func main() {
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

	log.Fatal(app.Listen(":5000"))
}
