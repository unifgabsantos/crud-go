package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/unifgabsantos/crud-go/src/controller"
)

// r é um *gin.RouterGroup, ex.: api := router.Group("/api")
func InitRoutes(r *gin.RouterGroup) {
	users := r.Group("/users")
	{
		users.GET("/id/:id", controller.FindUserById)
		users.GET("/email/:email", controller.FindUserByEmail)

		users.POST("", controller.CreateUser)
		users.PUT("/id/:id", controller.UpdateUser)
		users.DELETE("/id/:id", controller.DeleteUser)
	}
}
