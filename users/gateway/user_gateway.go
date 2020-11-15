package users

import (
	"encoding/json"
	"strings"
	"time"

	users "github.com/mendezdev/expenses_users-api/users/models"
	"github.com/mendezdev/expenses_users-api/utils/api_errors"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserGateway interface {
	CreateUser(u *users.CreateUserRequest) (*users.PublicUser, api_errors.RestErr)
	GetUserByID(string) (*users.PublicUser, api_errors.RestErr)
	DeleteUserByID(string) api_errors.RestErr
}

type UserDB struct {
	UserStorage
}

// NewUserGateway return an instance of UserGateway to interact with Users db
func NewUserGateway(db *mongo.Client) UserGateway {
	return &UserDB{NewUserStorageGateway(db)}
}

// CreateUser gets the user request struct and create the User
func (userDB *UserDB) CreateUser(u *users.CreateUserRequest) (*users.PublicUser, api_errors.RestErr) {
	err := Validate(u)
	if err != nil {
		return nil, err
	}

	var user = &users.User{
		FirstName:   u.FirstName,
		LastName:    u.LastName,
		Email:       u.Email,
		Password:    u.Password,
		DateCreated: time.Now().UTC().String(),
		Status:      "active",
	}
	if err := userDB.createUser(user); err != nil {
		return nil, err
	}
	return marshal(user), nil
}

func (userDB *UserDB) GetUserByID(userID string) (*users.PublicUser, api_errors.RestErr) {
	user, err := userDB.getUserByID(userID)
	if err != nil {
		return nil, err
	}
	return marshal(user), nil
}

func (userDB *UserDB) DeleteUserByID(userID string) api_errors.RestErr {
	return userDB.deleteUser(userID)
}

//Validate validates FirstName, LastName, Email and Password (only trimspace)
func Validate(userReq *users.CreateUserRequest) api_errors.RestErr {
	userReq.FirstName = strings.TrimSpace(userReq.FirstName)
	userReq.LastName = strings.TrimSpace(userReq.LastName)
	userReq.Email = strings.TrimSpace(userReq.Email)
	userReq.Password = strings.TrimSpace(userReq.Password)

	if userReq.Email == "" {
		return api_errors.NewBadRequestError("invalid value for email address field")
	}

	if userReq.Password == "" {
		return api_errors.NewBadRequestError("invalid value for password field")
	}

	return nil
}

func marshal(user *users.User) *users.PublicUser {
	userJSON, _ := json.Marshal(user)
	var publicUser users.PublicUser
	json.Unmarshal(userJSON, &publicUser)
	return &publicUser
}
