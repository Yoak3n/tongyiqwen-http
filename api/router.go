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
	chat := v1.Group("/chat")
	chat.POST("/completions", api.Ask)
	err := r.Run(fmt.Sprintf(":%d", config.GetConfig().Port))
	if err != nil {
		return
	}
}
