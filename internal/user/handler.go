package user

import (
	"advanced_rest_api/internal/handlers"
	"advanced_rest_api/pkg/logging"
	"github.com/gin-gonic/gin"
	"net/http"
)

var _ handlers.Handler = &handler{}

const (
	usersURL = "/users"
	userURL  = "/users/:uuid"
)

type handler struct {
	logger *logging.Logger
}

func NewUserHandler(logger *logging.Logger) handlers.Handler {
	return &handler{
		logger: logger,
	}
}

func (h *handler) Register(router *gin.Engine) {
	router.GET(usersURL, h.GetUsersList)
	router.POST(usersURL, h.CreateUser)
	router.GET(userURL, h.GetUserByUUID)
	router.PUT(userURL, h.UpdateUser)
	router.PATCH(userURL, h.PartiallyUpdateUser)
	router.DELETE(userURL, h.DeleteUser)
}

func (h *handler) GetUsersList(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, "this is list of users")
}

func (h *handler) CreateUser(c *gin.Context) {
	c.IndentedJSON(http.StatusCreated, "this is create user")
}

func (h *handler) GetUserByUUID(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, "this is user by uuid")
}

func (h *handler) UpdateUser(c *gin.Context) {
	c.IndentedJSON(http.StatusNoContent, "this is update user")
}

func (h *handler) PartiallyUpdateUser(c *gin.Context) {
	c.IndentedJSON(http.StatusNoContent, "this is partially update user")
}

func (h *handler) DeleteUser(c *gin.Context) {
	c.IndentedJSON(http.StatusNoContent, "his is delete user")
}
