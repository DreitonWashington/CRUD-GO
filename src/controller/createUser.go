package controller

import (
	"fmt"

	errorhandler "github.com/DreitonWashington/CRUD-GO/src/configuration/errorHandler"
	"github.com/DreitonWashington/CRUD-GO/src/controller/model/request"
	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context) {
	var userRequest request.UserRequest

	if err := c.ShouldBindJSON(&userRequest); err != nil {
		erro := errorhandler.NewBadRequestError(fmt.Sprintf("There are some incorrect fields, error=%s", err.Error()))
		c.JSON(erro.Code, erro)
		return
	}
	fmt.Println(userRequest)
}
