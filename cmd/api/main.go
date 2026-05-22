package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Mariano-SI/twitter-api/internal/config"
	userHandler "github.com/Mariano-SI/twitter-api/internal/handler/user"
	userRepository "github.com/Mariano-SI/twitter-api/internal/repository/user"
	userService "github.com/Mariano-SI/twitter-api/internal/service/user"
	"github.com/Mariano-SI/twitter-api/pkg/internalsql"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	config, err := config.LoadConfig()

	if err != nil {
		log.Fatal(err)
	}

	db, err := internalsql.ConnectMySQL(config.DBUser, config.DBPassword, config.DBHost, config.DBPort, config.DBName)

	if err != nil {
		log.Fatal(err)
	}

	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	r.GET("/hello", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "it's works",
		})
	})

	userRepository := userRepository.NewRepository(db)
	userService := userService.NewService(config, userRepository)
	userHandler := userHandler.NewHandler(r, userService)

	userHandler.RouteList()

	server := fmt.Sprintf("127.0.0.1:%s", config.Port)
	r.Run(server)
}
