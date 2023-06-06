package models

import (
	"fmt"
	"time"

	"gorm.io/gorm"
	// "gorm.io/gorm"
)

type User struct {
	// gorm.Model /* this is the default form of declaration */
	ID           uint           `gorm:"primaryKey;unique" json:"id"`
	Name         string         `gorm:"type:string;not null" json:"name"`
	Email        string         `gorm:"type:string;not null;unique;primaryKey" json:"email"`
	HashPassword string         `gorm:"type:string;not null" json:"-"`
	CreatedAt    time.Time      `gorm:"autoUpdateTime:milli" json:"created_at"`
	UpdatedAt    time.Time      `gorm:"autoUpdateTime:milli" json:"updated_at"`
	DeletedAt    gorm.DeletedAt `json:"-"`
	Posts        []Post         `gorm:"foreignKey:UserID" json:"posts"`
}

func (u *User) String() string {
	return fmt.Sprintf("userId: %d, userEmail: %s\n", u.ID, u.Email)
}
