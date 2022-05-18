package user

import (
	"advanced_rest_api/internal/handlers"
	"advanced_rest_api/pkg/logging"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"net/http"
)

var _ handlers.Handler = &handler{}

const (
	usersURL   = "/users"
	userURL    = "/users/:uuid"
	swaggerURL = "/swagger/*any"
)

type handler struct {
	logger *logging.Logger
}

func NewHandler(logger *logging.Logger) handlers.Handler {
	return &handler{
		logger: logger,
	}
}

func (h *handler) Register(router *gin.Engine) {
	router.GET(swaggerURL, ginSwagger.WrapHandler(swaggerFiles.Handler))
	router.GET(usersURL, h.GetList)
	router.POST(usersURL, h.CreateUser)
	router.GET(userURL, h.GetUserByUUID)
	router.PUT(userURL, h.UpdateUser)
	router.PATCH(userURL, h.PartiallyUpdateUser)
	router.DELETE(userURL, h.DeleteUser)
}

// GetList godoc
// @Summary ping example
// @Schemes
// @Description do ping
// @Tags example
// @Accept json
// @Produce json
// @Success 200 {string} this is list of users
// @Router /users [get]
func (h *handler) GetList(c *gin.Context) {
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
