package main

import (
	"avito_test_task/internal/db"
	"avito_test_task/internal/routers"
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	if err := db.InitDatabase(); err != nil {
		log.Fatal("init/connect to database error: " + err.Error())
	}
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	routers.InitRoutes(r)
	if err := r.Run(":8080"); err != nil {
		log.Fatal("failed to start server: " + err.Error())
	}
}
