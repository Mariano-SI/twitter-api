package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/Mariano-SI/twitter-api/internal/config"
	userHandler "github.com/Mariano-SI/twitter-api/internal/handler/user"
	refreshToken "github.com/Mariano-SI/twitter-api/internal/repository/refresh_token"
	userRepository "github.com/Mariano-SI/twitter-api/internal/repository/user"
	userService "github.com/Mariano-SI/twitter-api/internal/service/user"
	"github.com/Mariano-SI/twitter-api/pkg/internalSql"
	"github.com/Mariano-SI/twitter-api/pkg/r2"
	"github.com/gin-gonic/gin"

	postHandler "github.com/Mariano-SI/twitter-api/internal/handler/post"
	r2storage "github.com/Mariano-SI/twitter-api/internal/infra/storage/r2"
	postRepository "github.com/Mariano-SI/twitter-api/internal/repository/post"
	postImageRepository "github.com/Mariano-SI/twitter-api/internal/repository/post_image"
	postService "github.com/Mariano-SI/twitter-api/internal/service/post"
)

func main() {
	r := gin.Default()

	config, err := config.LoadConfig()

	if err != nil {
		log.Fatal(err)
	}

	db, err := internalSql.ConnectMySQL(config.DBUser, config.DBPassword, config.DBHost, config.DBPort, config.DBName)

	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	r2Client, err := r2.NewClient(
		context.Background(),
		config.R2AccountID,
		config.R2AccessKeyID,
		config.R2SecretAccessKey,
	)
	
	if err != nil {
		log.Fatal(err)
	}

	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	v1 := r.Group("/v1")

	v1.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	userRepository := userRepository.NewRepository(db)
	postRepository := postRepository.NewRepository(db)
	postImageRepository := postImageRepository.NewRepository(db)
	refreshTokenRepository := refreshToken.NewRepository(db)

	imageStorage := r2storage.NewStorage(r2Client, config.R2Bucket, config.R2PublicURL)
	transactor := internalSql.NewTransactor(db)

	postService := postService.NewService(transactor, postRepository, postImageRepository, imageStorage)
	userService := userService.NewService(config, userRepository, refreshTokenRepository)

	postHandler := postHandler.NewHandler(v1, postService)
	userHandler := userHandler.NewHandler(v1, userService)

	postHandler.RouteList(config.JwtSecret)
	userHandler.RouteList(config.JwtSecret)

	server := fmt.Sprintf("127.0.0.1:%s", config.Port)
	r.Run(server)
}
