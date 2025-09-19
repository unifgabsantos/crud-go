package routes

import (
	"github.com/gin-gonic/gin"
)

func InitializeRoutes(r *gin.RouterGroup) {
	r.GET("/User/:id")
	r.GET("/User/:email")
	r.POST("/User")
	r.PUT("/User/:id")
	r.DELETE("/User/:id")
}
