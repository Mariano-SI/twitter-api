package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Mariano-SI/twitter-api/internal/config"
	"github.com/Mariano-SI/twitter-api/pkg/internalsql"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	config, err := config.LoadConfig()

	if err != nil {
		log.Fatal(err)
	}

	_, err = internalsql.ConnectMySQL(config.DBUser, config.DBPassword, config.DBHost, config.DBPort, config.DBName)

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

	server := fmt.Sprintf("127.0.0.1:%s", config.Port)
	r.Run(server)
}
