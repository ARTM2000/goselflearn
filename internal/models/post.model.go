package models

import (
	"fmt"
	"time"

	"gorm.io/gorm"
)

type Post struct {
	ID          uint           `gorm:"primaryKey" json:"id"`
	Title       string         `gorm:"type:string" json:"title"`
	Description string         `gorm:"type:string" json:"description"`
	UserID      uint           `json:"user_id"`
	CreatedAt   time.Time      `gorm:"autoUpdateTime:milli" json:"created_at"`
	UpdatedAt   time.Time      `gorm:"autoUpdateTime:milli" json:"updated_at"`
	DeletedAt   gorm.DeletedAt `json:"-"`
}

func (p *Post) String() string {
	return fmt.Sprintf(
		"postId: %d, title: %s, description: %s",
		p.ID,
		p.Title,
		p.Description,
	)
}
