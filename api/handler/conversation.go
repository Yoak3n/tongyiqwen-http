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

func Reset(c *gin.Context) {
	id := c.Param("id")
	if id != "" {
		qianwen.Reset(id)
		c.String(200, "Conversation(id:%s) reset!", id)
	} else {
		c.String(400, "Conversation id not found!")
	}
}

// AskCompletions TODO
func AskCompletions(c *gin.Context) {

}
