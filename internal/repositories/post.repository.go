package repositories

import (
	"goselflearn/internal/initializers"
	"goselflearn/internal/models"

	"gorm.io/gorm"
)

func NewPostRepository() PostRepository {
	return PostRepository{
		DB: *initializers.DB,
	}
}

type PostRepository struct {
	DB gorm.DB
}

func (pr *PostRepository) CreatePost(post *models.Post) error {
	dbResult := pr.DB.Model(&models.Post{}).Create(post)
	if dbResult.Error != nil {
		return dbResult.Error
	}
	return nil
}

func (pr *PostRepository) FindPostsByUserId(userId uint) (*[]models.Post, error) {
	var posts []models.Post
	dbResult := pr.DB.Model(&models.Post{}).Where(models.Post{UserID: userId}).Find(&posts)
	if dbResult.Error != nil {
		return nil, dbResult.Error
	}
	return &posts, nil
}
