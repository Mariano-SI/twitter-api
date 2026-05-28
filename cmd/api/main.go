package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Mariano-SI/twitter-api/internal/config"
	userHandler "github.com/Mariano-SI/twitter-api/internal/handler/user"
	refreshToken "github.com/Mariano-SI/twitter-api/internal/repository/refresh_token"
	userRepository "github.com/Mariano-SI/twitter-api/internal/repository/user"
	userService "github.com/Mariano-SI/twitter-api/internal/service/user"
	"github.com/Mariano-SI/twitter-api/pkg/internalSql"
	"github.com/gin-gonic/gin"
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

	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	v1 := r.Group("/v1")

	v1.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	userRepository := userRepository.NewRepository(db)
	refreshTokenRepository := refreshToken.NewRepository(db)
	userService := userService.NewService(config, userRepository, refreshTokenRepository)
	userHandler := userHandler.NewHandler(v1, userService)

	userHandler.RouteList(config.JwtSecret)

	server := fmt.Sprintf("127.0.0.1:%s", config.Port)
	r.Run(server)
}
