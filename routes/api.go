package routes

import (
	"gin-app/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	r.GET("/users", controllers.GetAllUsers)
	r.GET("/hello", controllers.Hello)
	r.GET("/users/:id", controllers.GetUserById)
	r.POST("/users", controllers.CreateUser)
	r.PUT("/users/:id", controllers.UpdateUser)
	r.DELETE("/users/:id", controllers.DeleteUser)

	return r
}
