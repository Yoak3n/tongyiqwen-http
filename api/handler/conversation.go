package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"time"
	api "tongyiqwen/api/model"
	"tongyiqwen/qianwen"
)

func Ask(c *gin.Context) {
	req := &api.RequestBody{}
	err := c.BindJSON(req)
	if err != nil {
		return
	}
	// 处理请求
	receive := time.Now()
	answer := qianwen.Ask(req.Id, req.Preset, req.Content)
	c.JSON(200, gin.H{
		"answer": answer,
		"id":     req.Id,
		"cost":   fmt.Sprintf("%.2fs", time.Since(receive).Seconds()),
	})

}

// AskCompletions TODO
func AskCompletions(c *gin.Context) {

}
