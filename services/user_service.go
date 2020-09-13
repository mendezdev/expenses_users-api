package services

import (
	"github.com/mendezdev/expenses_users-api/domain/users"
	"github.com/mendezdev/expenses_users-api/utils/api_errors"
	"github.com/mendezdev/expenses_users-api/utils/crypto_utils"
	"github.com/mendezdev/expenses_users-api/utils/date_utils"
)

var (
	UsersService usersServiceInterface = &usersService{}
)

type usersService struct{}

type usersServiceInterface interface {
	GetUser(string) (*users.User, api_errors.RestErr)
	CreateUser(users.User) (*users.User, api_errors.RestErr)
	DeleteUser(string) api_errors.RestErr
	LoginUser(users.UserLoginRequest) (*users.User, api_errors.RestErr)
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
	hashedPassowrd, err := crypto_utils.Hash(user.Password)
	if err != nil {
		return nil, api_errors.NewInternalServerError("error trying to hash the password", err)
	}
	user.Password = hashedPassowrd
	if err := user.Save(); err != nil {
		return nil, err
	}

	return &user, nil
}

// DeleteUser by ID
func (s *usersService) DeleteUser(userID string) api_errors.RestErr {
	var user = &users.User{ID: userID}
	return user.Delete()
}

// LoginUser takes the credentials inside of UserLoginRequest and try to authenticate
func (s *usersService) LoginUser(userRequest users.UserLoginRequest) (*users.User, api_errors.RestErr) {
	userDao := &users.User{
		Email:    userRequest.Email,
		Password: userRequest.Password,
	}

	if err := userDao.FindByEmail(); err != nil {
		return nil, err
	}

	isValid := crypto_utils.CheckHash(userRequest.Password, userDao.Password)
	if !isValid {
		return nil, api_errors.NewBadRequestError("invalid credentials")
	}

	return userDao, nil
}
