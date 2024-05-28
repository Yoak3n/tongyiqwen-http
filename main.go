package main

import (
	"tongyiqwen/api"
	"tongyiqwen/package/util"
)

func init() {
	util.CreatePathNotExists("preset.json")
	util.CreatePathNotExists("config.yaml")
}

func main() {
	api.RunRouter()
}
