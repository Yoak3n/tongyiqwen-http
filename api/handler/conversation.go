package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"time"
	"tongyiqwen/api/model"
	"tongyiqwen/package/openai_model"
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
	//all, err := io.ReadAll(c.Request.Body)
	//if err != nil {
	//	return
	//}
	//defer c.Request.Body.Close()
	//fmt.Println("body1:", string(all))
	req := &openai_model.RequestBody{}
	err := c.BindJSON(req)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
	}
	//b, err := json.Marshal(req)
	//fmt.Println("body2:", string(b))
	res, err := qianwen.AskWithOpenAIStyle(req)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
	}
	c.JSON(200, res)

}
