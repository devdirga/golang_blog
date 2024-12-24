package main

import (
	"goblog/config"
	"goblog/handler"
	"goblog/model"
	"goblog/repository"
	"goblog/route"
	"goblog/service"
	"log"
	"net/http"

	"github.com/go-redis/redis/v8"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"

	"github.com/gofiber/adaptor/v2"
	jwtware "github.com/gofiber/contrib/jwt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

/*

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

	db.AutoMigrate(&model.Post{})
	db.AutoMigrate(&model.User{})

	postRepo := &repository.PostgresPostRespository{Db: db}
	postService := &service.PostServiceImpl{Repo: postRepo, RedisClient: rdb}
	postHandler := handler.NewPostHandler(postService)
	userRepo := &repository.PostgresUserRespository{Db: db}
	userService := &service.UserServiceImpl{Repo: userRepo}
	userHandler := handler.NewUserHandler(userService)

	app := fiber.New()
	app.Use(cors.New(cors.Config{AllowOrigins: "*"}))
	route.UserRoute(app, userHandler)
	route.PublicRoute(app, postHandler)
	app.Use(jwtware.New(jwtware.Config{
		SigningKey: jwtware.SigningKey{Key: []byte(config.GetConf().Secret)},
	}))
	route.PostRoute(app, postHandler)
	route.InfoRoute(app, userHandler)
	log.Fatal(app.Listen(":5000"))
}

*/

// Handler is the exported function required by Vercel
func Handler(w http.ResponseWriter, r *http.Request) {
	// Create the app
	app := setupApp()
	// Convert Fiber to net/http
	adaptor.FiberApp(app)(w, r)
}

func setupApp() *fiber.App {
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

	db.AutoMigrate(&model.Post{})
	db.AutoMigrate(&model.User{})

	postRepo := &repository.PostgresPostRespository{Db: db}
	postService := &service.PostServiceImpl{Repo: postRepo, RedisClient: rdb}
	postHandler := handler.NewPostHandler(postService)
	userRepo := &repository.PostgresUserRespository{Db: db}
	userService := &service.UserServiceImpl{Repo: userRepo}
	userHandler := handler.NewUserHandler(userService)

	app := fiber.New()
	app.Use(cors.New(cors.Config{AllowOrigins: "*"}))
	route.UserRoute(app, userHandler)
	route.PublicRoute(app, postHandler)
	app.Use(jwtware.New(jwtware.Config{
		SigningKey: jwtware.SigningKey{Key: []byte(config.GetConf().Secret)},
	}))
	route.PostRoute(app, postHandler)
	route.InfoRoute(app, userHandler)

	return app
}
