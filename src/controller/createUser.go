package controller

import (
	"fmt"
	"log"

	"github.com/DreitonWashington/CRUD-GO/src/configuration/validation"
	"github.com/DreitonWashington/CRUD-GO/src/controller/model/request"
	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context) {
	var userRequest request.UserRequest

	if err := c.ShouldBindJSON(&userRequest); err != nil {
		log.Printf("There are some incorrect fields, error=%s", err.Error())
		erro := validation.ValidateUserError(err)
		c.JSON(erro.Code, erro)
		return
	}
	fmt.Println(userRequest)
}
