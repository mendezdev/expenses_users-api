package users

import (
	"strings"

	"github.com/mendezdev/golang_mongo-example/utils/api_errors"
)

//User is the domain
type User struct {
	ID          string `bson:"_id,omitempty"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Email       string `json:"email"`
	DateCreated string `json:"date_created"`
	Status      string `json:"status"`
	Password    string `json:"password"`
}

//Validate validates FirstName, LastName, Email and Password (only for trimspace)
func (user *User) Validate() api_errors.RestErr {
	user.FirstName = strings.TrimSpace(user.FirstName)
	user.LastName = strings.TrimSpace(user.LastName)
	user.Email = strings.TrimSpace(user.Email)
	user.Password = strings.TrimSpace(user.Password)

	if user.Email == "" {
		return api_errors.NewBadRequestError("invalid email address")
	}

	if user.Password == "" {
		return api_errors.NewBadRequestError("invalid password")
	}

	emailAvailableErr := user.IsAvailableEmail()
	if emailAvailableErr != nil {
		return emailAvailableErr
	}

	return nil
}
