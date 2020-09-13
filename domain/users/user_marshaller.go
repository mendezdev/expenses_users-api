package users

import "encoding/json"

type PublicUser struct {
	ID          string `json:"id"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Email       string `json:"email"`
	DateCreated string `json:"date_created"`
	Status      string `json:"status"`
}

func (user *User) Marshal() interface{} {
	userJson, _ := json.Marshal(user)
	var publicUser PublicUser
	json.Unmarshal(userJson, &publicUser)
	return publicUser
}
