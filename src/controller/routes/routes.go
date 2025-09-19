package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/unifgabsantos/crud-go/src/controller"
)

func InitRoutes(r *gin.RouterGroup) {
	r.GET("/User/:id", controller.FindUserById)
	r.GET("/User/:email", controller.FindUserByEmail)
	r.POST("/User", controller.CreateUser)
	r.PUT("/User/:id", controller.UpdateUser)
	r.DELETE("/User/:id", controller.DeleteUser)
}
