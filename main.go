package main

import (
	"avito_test_task/db"
	"avito_test_task/routers"
	"github.com/gin-gonic/gin"
)

func main() {
	if err := db.InitDatabase(); err != nil {
		panic("init/connect to database error: " + err.Error())
	}
	r := gin.Default()
	routers.InitRoutes(r)
	if err := r.Run(":8080"); err != nil {
		panic("failed to start server: " + err.Error())
	}
}
