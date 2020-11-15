package users

import (
	"context"
	"fmt"
	"log"

	users "github.com/mendezdev/expenses_users-api/users/models"
	"github.com/mendezdev/expenses_users-api/utils/api_errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

const (
	// DbName is the db name
	DbName = "mydb"
	// UsersCollectionName is the name for the users collection
	UsersCollectionName = "users"
)

// UserStorage represents the actions that impacts in the user database
type UserStorage interface {
	createUser(*users.User) api_errors.RestErr
	getUserByID(string) (*users.User, api_errors.RestErr)
	deleteUser(string) api_errors.RestErr
}

// UserService implements the storage actions
type UserService struct {
	db *mongo.Client
}

// NewUserStorageGateway returns the struct that implements the UserStorage the interacts with the Users db
func NewUserStorageGateway(db *mongo.Client) UserStorage {
	return &UserService{db: db}
}

func (s *UserService) createUser(u *users.User) api_errors.RestErr {
	collection := getUserCollection(s.db)
	insertResult, err := collection.InsertOne(context.TODO(), u)
	if err != nil {
		log.Println("err when trying to insert User", err)
		return api_errors.NewInternalServerError("database error creating user", err)
	}
	if insertedID, ok := insertResult.InsertedID.(primitive.ObjectID); ok {
		u.ID = insertedID.Hex()
	}
	return nil
}

func (s *UserService) getUserByID(ID string) (*users.User, api_errors.RestErr) {
	userID, err := getUserID(ID)
	if err != nil {
		return nil, err
	}

	collection := getUserCollection(s.db)
	filter := bson.D{{"_id", userID}}
	user := &users.User{}
	userGetErr := collection.FindOne(context.TODO(), filter).Decode(&user)
	if userGetErr != nil {
		if userGetErr == mongo.ErrNoDocuments {
			return nil, api_errors.NewNotFoundError(fmt.Sprintf("user not found with given id: %s", ID))
		}
		log.Println("error trying to get user by id from db", userGetErr)
		return nil, api_errors.NewInternalServerError("database error", userGetErr)
	}
	return user, nil
}

func (s *UserService) deleteUser(ID string) api_errors.RestErr {
	userID, err := getUserID(ID)
	if err != nil {
		return err
	}
	filter := bson.D{{"_id", userID}}
	_, deleteErr := getUserCollection(s.db).DeleteOne(context.TODO(), filter)
	if deleteErr != nil {
		return api_errors.NewInternalServerError(fmt.Sprintf("error trying to delete user with given id: %s", ID), deleteErr)
	}
	return nil
}

func getUserCollection(db *mongo.Client) *mongo.Collection {
	return db.Database(DbName).Collection(UsersCollectionName)
}

func getUserID(userID string) (*primitive.ObjectID, api_errors.RestErr) {
	objectUserID, userIDErr := primitive.ObjectIDFromHex(userID)
	if userIDErr != nil {
		log.Println("error when trying to parse ID to get user in db", userIDErr)
		return nil, api_errors.NewBadRequestError("invalid id to get user")
	}

	return &objectUserID, nil
}
