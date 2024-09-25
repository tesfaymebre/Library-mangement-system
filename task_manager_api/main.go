package main

import (
	"task_manager/router"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	router.SetupRouter(r)
	r.Run("localhost:8080")
}
