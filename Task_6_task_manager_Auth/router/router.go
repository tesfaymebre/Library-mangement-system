package router

import (
	"task_manager/controllers"
	"task_manager/middleware"

	"github.com/gin-gonic/gin"
)

func SetupRouter(r *gin.Engine) {
	user := r.Group("/user")
	{
		user.POST("/register", controllers.RegisterUser)
		user.POST("/verify", controllers.VerifyOTP)
		user.POST("/resendOtp", controllers.ResendOTP)
		user.POST("/login", controllers.Login)
		user.GET("/allUsers", middleware.AuthMiddleware(), middleware.AdminMiddleware(), controllers.GetAllUsers)
		user.PUT("/promote/:id", middleware.AuthMiddleware(), middleware.AdminMiddleware(), controllers.PromoteUser)
	}

	task := r.Group("/tasks")
	{
		task.GET("/", controllers.GetTasks)
		task.GET("/:id", controllers.GetTaskByID)

		task.POST("/", middleware.AuthMiddleware(), middleware.AdminMiddleware(), controllers.CreateTask)
		task.PUT("/:id", middleware.AuthMiddleware(), middleware.AdminMiddleware(), controllers.UpdateTask)
		task.DELETE("/:id", middleware.AuthMiddleware(), middleware.AdminMiddleware(), controllers.DeleteTask)
	}

	token := r.Group("/token")
	{
		token.POST("/refresh", controllers.RefreshToken)
	}

}
