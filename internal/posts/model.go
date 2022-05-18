package posts

type Post struct {
	ID         string `json:"id" bson:"_id,omitempty"`
	Article    int64  `json:"article" bson:"article"`
	Slug       string `json:"slug" bson:"slug"`
	Title      string `json:"title" bson:"title"`
	Author     string `json:"author" bson:"author"`
	Summary    string `json:"summary" bson:"summary"`
	Content    string `json:"content" bson:"content"`
	ViewsCount int64  `json:"views_count" bson:"views_count"`
	Favorites  bool   `json:"favorites" bson:"favorites"`
	Vote       int8   `json:"vote" bson:"vote"`
	Rating     int64  `json:"rating" bson:"rating"`
}

type CreatePostDTO struct {
	Article    int64  `json:"article"`
	Slug       string `json:"slug"`
	Title      string `json:"title"`
	Author     string `json:"author"`
	Summary    string `json:"summary"`
	Content    string `json:"content"`
	ViewsCount int64  `json:"views_count"`
	Favorites  bool   `json:"favorites"`
	Vote       int8   `json:"vote"`
	Rating     int64  `json:"rating"`
}
