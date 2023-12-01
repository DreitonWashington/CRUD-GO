package service

import (
	"fmt"

	errorhandler "github.com/DreitonWashington/CRUD-GO/src/configuration/errorHandler"
	"github.com/DreitonWashington/CRUD-GO/src/configuration/logger"
	"github.com/DreitonWashington/CRUD-GO/src/model"
	"go.uber.org/zap"
)

func (ud *userDomainService) CreateUser(userDomain model.UserDomainInterface) *errorhandler.Erro {
	logger.Info("Init createUser model", zap.String("journey", "createUser"))
	userDomain.EncryptPassword()
	fmt.Println(userDomain)
	return nil
}
