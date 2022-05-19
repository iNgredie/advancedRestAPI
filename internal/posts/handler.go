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
	postURL  = "/posts:uuid"
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
	router.POST(postURL, h.CreatePost)
	router.GET(postURL, h.GetPostByUUID)
	router.PUT(postURL, h.UpdatePost)
	router.PATCH(postURL, h.PartiallyUpdatePost)
	router.DELETE(postURL, h.DeletePost)
}

func (h *handler) GetPostsList(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, "this is list of posts")
}

func (h *handler) CreatePost(c *gin.Context) {
	c.IndentedJSON(http.StatusCreated, "this is create post")
}

func (h *handler) GetPostByUUID(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, "this is post by uuid")
}

func (h *handler) UpdatePost(c *gin.Context) {
	c.IndentedJSON(http.StatusNoContent, "this is update post")
}

func (h *handler) PartiallyUpdatePost(c *gin.Context) {
	c.IndentedJSON(http.StatusNoContent, "this is partially update post")
}

func (h *handler) DeletePost(c *gin.Context) {
	c.IndentedJSON(http.StatusNoContent, "his is delete post")
}
