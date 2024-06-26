package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"tongyiqwen/api/handler"
	"tongyiqwen/config"
)

var r *gin.Engine

func RunRouter() {
	r = gin.Default()
	v1 := r.Group("/v1")
	v1.POST("/chat", api.Ask)
	v1.GET("/chat/reset/:id", api.Reset)
	v1.POST("/chat/completions", api.AskCompletions)
	v1.POST("/preset", api.Upload)
	v1.GET("/preset", api.GetPreset)
	err := r.Run(fmt.Sprintf(":%d", config.GetConfig().Port))
	if err != nil {
		return
	}
}
