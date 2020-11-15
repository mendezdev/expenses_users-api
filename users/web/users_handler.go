package users

import (
	"net/http"

	"github.com/gin-gonic/gin"
	usersGtw "github.com/mendezdev/expenses_users-api/users/gateway"
	usersModels "github.com/mendezdev/expenses_users-api/users/models"
	"github.com/mendezdev/expenses_users-api/utils/api_errors"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserHTTPService interface {
	CreateUserHandler(c *gin.Context)
	GetUserByIDHandler(c *gin.Context)
	DeleteUserByIDHandler(c *gin.Context)
}

type userHTTPServiceImpl struct {
	gtw usersGtw.UserGateway
}

func NewUserHTTPService(db *mongo.Client) UserHTTPService {
	return &userHTTPServiceImpl{
		gtw: usersGtw.NewUserGateway(db),
	}
}

func (s *userHTTPServiceImpl) CreateUserHandler(c *gin.Context) {
	var userRequest usersModels.CreateUserRequest
	if err := c.ShouldBindJSON(&userRequest); err != nil {
		restErr := api_errors.NewBadRequestError("invalid json body")
		c.JSON(restErr.Status(), restErr)
		return
	}

	result, saveErr := s.gtw.CreateUser(&userRequest)
	if saveErr != nil {
		c.JSON(saveErr.Status(), saveErr)
		return
	}
	c.JSON(http.StatusCreated, result)
}

func (s *userHTTPServiceImpl) GetUserByIDHandler(c *gin.Context) {
	userID, userIDErr := getUserID(c)
	if userIDErr != nil {
		c.JSON(userIDErr.Status(), userIDErr)
		return
	}

	user, getErr := s.gtw.GetUserByID(*userID)
	if getErr != nil {
		c.JSON(getErr.Status(), getErr)
		return
	}
	c.JSON(http.StatusOK, user)
}

func (s *userHTTPServiceImpl) DeleteUserByIDHandler(c *gin.Context) {
	userID, userIDErr := getUserID(c)
	if userIDErr != nil {
		c.JSON(userIDErr.Status(), userIDErr)
		return
	}

	deleteErr := s.gtw.DeleteUserByID(*userID)
	if deleteErr != nil {
		c.JSON(deleteErr.Status(), deleteErr)
		return
	}
	c.JSON(http.StatusOK, map[string]string{"status": "deleted"})
}

func getUserID(c *gin.Context) (*string, api_errors.RestErr) {
	userID := c.Param("user_id")
	if userID == "" {
		return nil, api_errors.NewBadRequestError("user id should be a number")
	}

	return &userID, nil
}
