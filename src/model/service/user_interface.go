package service

import (
	errorhandler "github.com/DreitonWashington/CRUD-GO/src/configuration/errorHandler"
	"github.com/DreitonWashington/CRUD-GO/src/model"
)

type UserDomainService interface {
	CreateUser(model.UserDomainInterface) *errorhandler.Erro
	UpdateUser(string, model.UserDomainInterface) *errorhandler.Erro
	FindUser(string) (*model.UserDomainInterface, *errorhandler.Erro)
	DeleteUser(string) *errorhandler.Erro
}

type userDomainService struct {
}

func NewUserDomainService() UserDomainService {
	return &userDomainService{}
}
