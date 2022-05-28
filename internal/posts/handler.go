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
	postURL  = "/posts:id"
)

type handler struct {
	logger *logging.Logger
}

func NewPostsHandler(logger *logging.Logger) handlers.Handler {
	return &handler{
		logger: logger,
	}
}

// TODO gets from database
var posts = []Post{
	{ID: "1", Article: 123, Slug: "post_1", Title: "Первый пост", Author: "1", Summary: "text", Content: "text", ViewsCount: 0, Favorites: false, Vote: 1, Rating: 1},
	{ID: "2", Article: 123, Slug: "post_2", Title: "Второй пост", Author: "1", Summary: "text", Content: "text", ViewsCount: 0, Favorites: false, Vote: 1, Rating: 1},
	{ID: "3", Article: 123, Slug: "post_3", Title: "Третий пост", Author: "1", Summary: "text", Content: "text", ViewsCount: 0, Favorites: false, Vote: 1, Rating: 1},
}

func (h *handler) Register(router *gin.Engine) {
	router.GET(postsURL, h.GetPostsList)
	router.POST(postURL, h.CreatePost)
	router.GET(postURL, h.GetPostByID)
	router.PUT(postURL, h.UpdatePost)
	router.PATCH(postURL, h.PartiallyUpdatePost)
	router.DELETE(postURL, h.DeletePost)
}

func (h *handler) GetPostsList(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, posts)
}

func (h *handler) CreatePost(c *gin.Context) {
	var newPost Post
	err := c.BindJSON(&newPost)
	if err != nil {
		return
	}

	posts = append(posts, newPost)
	c.IndentedJSON(http.StatusCreated, newPost)
}

func (h *handler) GetPostByID(c *gin.Context) {
	id := c.Param("id")

	for _, post := range posts {
		if post.ID == id {
			c.IndentedJSON(http.StatusOK, post)
			return
		}
	}
	c.IndentedJSON(http.StatusOK, gin.H{"message": "post not found"})
}

func (h *handler) UpdatePost(c *gin.Context) {
	var newPost Post
	err := c.BindJSON(&newPost)
	if err != nil {
		return
	}

	id := c.Param("id")

	for _, post := range posts {
		if post.ID == id {
			post = newPost
			c.IndentedJSON(http.StatusOK, post)
			return
		}
	}

	c.IndentedJSON(http.StatusOK, gin.H{"message": "post not found"})
}

func (h *handler) PartiallyUpdatePost(c *gin.Context) {
	c.IndentedJSON(http.StatusNoContent, "this is partially update post")
}

func (h *handler) DeletePost(c *gin.Context) {
	id := c.Param("id")

	for _, post := range posts {
		if post.ID != id {
			posts = append(posts, post)
			c.IndentedJSON(http.StatusOK, post)
			return
		}
	}

	c.IndentedJSON(http.StatusOK, gin.H{"message": "post not found"})
}
