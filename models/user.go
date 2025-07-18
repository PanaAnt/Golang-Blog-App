package models

import "github.com/google/uuid"

type User struct {
	ID       uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	Username string    `json:"username"`
	Email    string    `gorm:"unique"`
	Password string    `json:"password"`
	Posts    []Post    `gorm:"many2many:user_post;"`
}
