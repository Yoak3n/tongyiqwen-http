package main

import (
	"tongyiqwen/api"
	"tongyiqwen/package/util"
	"tongyiqwen/qianwen"
)

func init() {
	qianwen.CreateToken()
	util.CreatePathNotExists("preset.json")
	util.CreatePathNotExists("config.yaml")
}

func main() {
	api.RunRouter()
}
