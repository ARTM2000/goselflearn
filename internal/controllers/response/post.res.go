package response

import "github.com/ARTM2000/goselflearn/internal/models"

type newPostDataWrapper struct {
	Post models.Post `json:"post"`
}
type CreatePost struct {
	baseRes
	Data newPostDataWrapper `json:"data"`
}

type allPostsDataWrapper struct {
	Posts []models.Post `json:"posts"`
}
type GetUserPosts struct {
	baseRes
	Data allPostsDataWrapper `json:"data"`
}
