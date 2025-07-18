package repository

import (
	"blogApp/db"
	"blogApp/models"
)

type PostRepository interface {
	CreatePost(post *models.Post) error
	UpdatePost(post *models.Post) error
	GetPostByID(id string) (*models.Post, error)
	GetAllPosts() ([]models.Post, error)
	DeletePostByID(id string) error
}


type PostRepo struct{}

func (r *PostRepo) CreatePost(post *models.Post) error {
	err := db.DB.Create(&post).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *PostRepo) UpdatePost(post *models.Post) error {
	err := db.DB.Save(&post).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *PostRepo) GetPostByID(id string) (*models.Post, error) {
	var post models.Post
	if err := db.DB.First(&post, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return &post, nil
}

func (r *PostRepo) GetAllPosts() ([]models.Post, error) {
	var posts []models.Post
	err := db.DB.Find(&posts).Error
	if err != nil {
		return nil, err
	}
	return posts, nil
}

func (r *PostRepo) DeletePostByID(id string) error {
	err := db.DB.Delete(&models.Post{}, "id = ?", id).Error
	if err != nil {
		return err
	}
	return nil
}