package controller

import (
	"net/http"

	"github.com/DreitonWashington/CRUD-GO/src/configuration/logger"
	"github.com/DreitonWashington/CRUD-GO/src/configuration/validation"
	"github.com/DreitonWashington/CRUD-GO/src/controller/model/request"
	"github.com/DreitonWashington/CRUD-GO/src/model"
	"github.com/DreitonWashington/CRUD-GO/src/model/service"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

var (
	UserDomainInterface model.UserDomainInterface
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

	service := service.NewUserDomainService()
	domain := model.NewUserDomain(userRequest.Email, userRequest.Password, userRequest.Name, userRequest.Age)
	if err := service.CreateUser(domain); err != nil {
		c.JSON(err.Code, err)
		return
	}
	logger.Info("User created successfully", zap.String("journey", "createUser"))
	c.JSON(http.StatusOK, "")
}
