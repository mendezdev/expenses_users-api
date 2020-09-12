package services

import (
	"github.com/mendezdev/golang_mongo-example/domain/users"
	"github.com/mendezdev/golang_mongo-example/utils/api_errors"
	"github.com/mendezdev/golang_mongo-example/utils/date_utils"
)

var (
	UsersService usersServiceInterface = &usersService{}
)

type usersService struct{}

type usersServiceInterface interface {
	GetUser(string) (*users.User, api_errors.RestErr)
	CreateUser(users.User) (*users.User, api_errors.RestErr)
	DeleteUser(string) api_errors.RestErr
}

func (s *usersService) GetUser(userID string) (*users.User, api_errors.RestErr) {
	result := &users.User{ID: userID}
	if err := result.Get(); err != nil {
		return nil, err
	}
	return result, nil
}

func (s *usersService) CreateUser(user users.User) (*users.User, api_errors.RestErr) {
	if err := user.Validate(); err != nil {
		return nil, err
	}

	user.Status = "active" //TODO change this
	user.DateCreated = date_utils.GetNowString()
	if err := user.Save(); err != nil {
		return nil, err
	}

	return &user, nil
}

func (s *usersService) DeleteUser(userId string) api_errors.RestErr {
	var user = &users.User{ID: userId}
	return user.Delete()
}
