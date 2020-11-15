package users

//User is the domain
type User struct {
	ID          string `json:"id" bson:"_id,omitempty"`
	FirstName   string `json:"first_name" bson:"first_name"`
	LastName    string `json:"last_name" bson:"last_name"`
	Email       string `json:"email"`
	DateCreated string `json:"date_created" bson:"date_created"`
	Status      string `json:"status"`
	Password    string `json:"password"`
}

type CreateUserRequest struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	Password  string `json:"password"`
}

type PublicUser struct {
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Email       string `json:"email"`
	DateCreated string `json:"date_created"`
	Status      string `json:"status"`
}
