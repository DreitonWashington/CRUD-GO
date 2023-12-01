package controller

import (
	"fmt"

	"github.com/DreitonWashington/CRUD-GO/src/configuration/logger"
	"github.com/DreitonWashington/CRUD-GO/src/configuration/validation"
	"github.com/DreitonWashington/CRUD-GO/src/controller/model/request"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func CreateUser(c *gin.Context) {
	logger.Info("Init CreateUser controller", zap.String("journey", "createUser"))
	var userRequest request.UserRequest

	if err := c.ShouldBindJSON(&userRequest); err != nil {
		logger.Error("Error trying to validate user info", err, zap.String("journey", "createUser"))
		erro := validation.ValidateUserError(err)
		c.JSON(erro.Code, erro)
		return
	}
	fmt.Println(userRequest)
	logger.Info("User created successfully", zap.String("journey", "createUser"))
}
