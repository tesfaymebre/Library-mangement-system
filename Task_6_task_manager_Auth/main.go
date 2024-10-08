package main

import (
	"task_manager/data"
	"task_manager/router"

	"github.com/gin-gonic/gin"
)

func main() {
	data.InitializeMongoDB()
	r := gin.Default()
	router.SetupRouter(r)
	r.Run("localhost:8080")
}
