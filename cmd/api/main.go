package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Mariano-SI/twitter-api/internal/app"
	commentHandler "github.com/Mariano-SI/twitter-api/internal/handler/comment"
	commentLikeHandler "github.com/Mariano-SI/twitter-api/internal/handler/comment_like"
	followHandler "github.com/Mariano-SI/twitter-api/internal/handler/follow"
	postHandler "github.com/Mariano-SI/twitter-api/internal/handler/post"
	postLikeHandler "github.com/Mariano-SI/twitter-api/internal/handler/post_like"
	userHandler "github.com/Mariano-SI/twitter-api/internal/handler/user"
	"github.com/gin-gonic/gin"
)

func main() {
	deps, err := app.New()
	if err != nil {
		log.Fatal(err)
	}
	defer deps.DB.Close()

	r := gin.Default()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	v1 := r.Group("/v1")

	v1.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	userHandler.Register(v1, deps)
	postHandler.Register(v1, deps)
	postLikeHandler.Register(v1, deps)
	commentHandler.Register(v1, deps)
	commentLikeHandler.Register(v1, deps)
	followHandler.Register(v1, deps)

	server := fmt.Sprintf("127.0.0.1:%s", deps.Config.Port)
	r.Run(server)
}
