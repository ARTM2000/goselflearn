package services

import (
	"fmt"

	"github.com/ARTM2000/goselflearn/internal/common"
	"github.com/ARTM2000/goselflearn/internal/controllers/dto"
	"github.com/ARTM2000/goselflearn/internal/models"
	"github.com/ARTM2000/goselflearn/internal/repositories"
)

func NewPostService() postService {
	postRepo := repositories.NewPostRepository()
	return postService{
		postRepository: postRepo,
	}
}

type postService struct {
	postRepository repositories.PostRepository
}

func (ps *postService) CreatePost(post *dto.CreatePost, userId uint) (*models.Post, error) {
	newPost := models.Post{
		Title:       post.Title,
		Description: post.Description,
		UserID:      userId,
	}

	err := ps.postRepository.CreatePost(&newPost)
	if err != nil {
		fmt.Println("create post error:", err.Error())
		return nil, fmt.Errorf(common.MessageInternalServerError)
	}
	return &newPost, nil
}

func (ps *postService) FindUserPosts(userId uint) (*[]models.Post, error) {
	posts, err := ps.postRepository.FindPostsByUserId(userId)
	if err != nil {
		fmt.Println("find user posts error:", err.Error())
		return nil, fmt.Errorf(common.MessageInternalServerError)
	}
	return posts, nil
}
