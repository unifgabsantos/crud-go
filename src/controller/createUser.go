package controller

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/unifgabsantos/crud-go/src/configuration/rest_err"
	"github.com/unifgabsantos/crud-go/src/controller/model/request"
)

func CreateUser(c *gin.Context) {
	var userRequest request.UserRequest
	if err := c.ShouldBindJSON(&userRequest); err != nil {
		restErr := rest_err.NewBadRequestError(
			fmt.Sprintf("invalid json body: %s", err.Error()))
		c.JSON(int(restErr.Code), restErr)
		return
	}
	fmt.Println(userRequest)
}
