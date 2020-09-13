package users

import (
	"context"
	"fmt"

	"github.com/mendezdev/expenses_users-api/db/mongodb"
	"github.com/mendezdev/expenses_users-api/utils/api_errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

const (
	//DB_NAME is the db name
	DB_NAME = "mydb"
	//USERS_COLLECTION_NAME is the name for the users collection
	USERS_COLLECTION_NAME = "users"
)

//Save save a new user to the users collections
func (user *User) Save() api_errors.RestErr {
	collection := getUserCollection()
	insertResult, insertErr := collection.InsertOne(context.TODO(), user)

	if insertErr != nil {
		fmt.Println("error when trying to insert User", insertErr)
		return api_errors.NewInternalServerError("database error", insertErr)
	}

	if insertedID, ok := insertResult.InsertedID.(primitive.ObjectID); ok {
		user.ID = insertedID.Hex()
	}
	return nil
}

//Get get the user by the ID given
func (user *User) Get() api_errors.RestErr {
	userID, userIDErr := getUserID(user.ID)
	if userIDErr != nil {
		return userIDErr
	}

	findErr := user.findOneByFilter("_id", userID)
	if findErr != nil {
		if findErr == mongo.ErrNoDocuments {
			return api_errors.NewNotFoundError(fmt.Sprintf("user not found with given id: %s", user.ID))
		}
		fmt.Println("error trying to find document", findErr)
		return api_errors.NewInternalServerError("database error", findErr)
	}
	return nil
}

//Delete delete the document from the collection by the userID
func (user *User) Delete() api_errors.RestErr {
	userID, userIDErr := getUserID(user.ID)
	if userIDErr != nil {
		return userIDErr
	}

	filter := bson.D{{"_id", userID}}
	_, err := getUserCollection().DeleteOne(context.TODO(), filter)
	if err != nil {
		return api_errors.NewInternalServerError(fmt.Sprintf("error trying to delete user with id %s", user.ID), err)
	}
	return nil
}

// FindByEmailAndPassword try to get a User by the email and password
func (user *User) FindByEmailAndPassword() api_errors.RestErr {
	collection := getUserCollection()
	filter := bson.D{{"email", user.Email}, {"password", user.Password}}

	userGetErr := collection.FindOne(context.TODO(), filter).Decode(&user)
	if userGetErr != nil {
		if userGetErr == mongo.ErrNoDocuments {
			return api_errors.NewNotFoundError("user not found with given values")
		}
		return api_errors.NewInternalServerError("database error", userGetErr)
	}

	return nil
}

// FindByEmail fill the &user with the data user if the filter founds one
// otherwise returns an RestErr
func (user *User) FindByEmail() api_errors.RestErr {
	return user.searchByEmail(false)
}

// IsAvailableEmail returns an RestErr if the email already exist
func (user *User) IsAvailableEmail() api_errors.RestErr {
	return user.searchByEmail(true)
}

func (user *User) findOneByFilter(fieldName string, fieldValue interface{}) error {
	collection := getUserCollection()
	filter := bson.D{{fieldName, fieldValue}}

	userGetErr := collection.FindOne(context.TODO(), filter).Decode(&user)
	if userGetErr != nil {
		return userGetErr
	}

	return nil
}

func (user *User) searchByEmail(isForEmailAvailable bool) api_errors.RestErr {
	findErr := user.findOneByFilter("email", user.Email)
	if findErr != nil {
		if findErr == mongo.ErrNoDocuments {
			if isForEmailAvailable {
				return nil
			}
			return api_errors.NewNotFoundError("no user found with given email")
		}
		return api_errors.NewInternalServerError("database error", findErr)
	}

	if isForEmailAvailable {
		return api_errors.NewBadRequestError("email is not available")
	}
	return nil
}

func getUserCollection() *mongo.Collection {
	return mongodb.MongoClient.Database(DB_NAME).Collection(USERS_COLLECTION_NAME)
}

func getUserID(userID string) (*primitive.ObjectID, api_errors.RestErr) {
	objectUserID, userIDErr := primitive.ObjectIDFromHex(userID)
	if userIDErr != nil {
		fmt.Println("error when trying to parse ID to get user in db", userIDErr)
		return nil, api_errors.NewBadRequestError("invalid id to get user")
	}

	return &objectUserID, nil
}
