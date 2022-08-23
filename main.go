package main

import (
	"labplatform/model"
	"labplatform/router"
)

func main() {
	model.InitDb()
	model.InitRedis()
	router.InitRouter()
}
