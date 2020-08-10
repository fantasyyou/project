package main

import (
	orm "gin/api/dao"
	"gin/api/router"
)

func main() {
	defer orm.Eloquent.Close()
	initRouter := router.InitRouter()
	_ = initRouter.Run(":8000")
}