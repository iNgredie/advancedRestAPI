package posts

import (
	"advanced_rest_api/internal/handlers"
	"advanced_rest_api/pkg/logging"
	"github.com/gin-gonic/gin"
	"net/http"
)

var _ handlers.Handler = &handler{}

const (
	postsURL = "/posts"
)

type handler struct {
	logger *logging.Logger
}

func NewPostsHandler(logger *logging.Logger) handlers.Handler {
	return &handler{
		logger: logger,
	}
}

func (h *handler) Register(router *gin.Engine) {
	router.GET(postsURL, h.GetPostsList)
}

func (h *handler) GetPostsList(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, "this is list of posts")
}
