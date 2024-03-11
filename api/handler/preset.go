package api

import (
	"github.com/gin-gonic/gin"
	api "tongyiqwen/api/model"
	"tongyiqwen/plugin"
)

func Upload(c *gin.Context) {
	preset := &api.UploadPreset{}
	err := c.BindJSON(preset)
	if err != nil {
		c.String(400, "Bad Request")
	}

	if preset.Type == "text" {
		err = plugin.PushNewTextPreset(preset.Name, preset.Text)
		if err != nil {
			c.String(500, "Internal Server Error:%v", err)
		}
		c.String(200, "text")
	} else if preset.Type == "map" {
		err = plugin.PushNewMapPreset(preset.Name, preset.Map)
		c.String(200, "map")
	} else {
		c.String(400, "Bad Request")
	}

}

func GetPreset(c *gin.Context) {
	preset, err := plugin.GetAllPreset()
	if err != nil {
		c.String(400, err.Error())
	}
	c.JSON(200, gin.H{
		"data": preset,
	})
}
