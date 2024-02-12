package main

import (
	"tongyiqwen/api"
	"tongyiqwen/qianwen"
)

func init() {
	qianwen.CreateToken()
}

func main() {
	api.RunRouter()
}
