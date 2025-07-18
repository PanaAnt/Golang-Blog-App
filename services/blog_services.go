package services

import (
	"blogApp/models"
	"blogApp/repository"

	"github.com/google/uuid"
)

type PostService struct {
	Repo repository.PostRepository
}

func (s *PostService) CreatePost(req *models.Post) error {
	myuuid := uuid.New()
	req.ID = myuuid

	err := s.Repo.CreatePost(req)
	if err != nil {
		return err
	}
	return nil
}

func (s *PostService) UpdatePost(id string, title string, content string) (*models.Post, error) {
	post, err := s.Repo.GetPostByID(id)
	if err != nil {
		return nil, err
	}

	post.Title = title
	post.Content = content

	err = s.Repo.UpdatePost(post)
	if err != nil {
		return nil, err
	}

	return post, nil
}

func (s *PostService) ViewPostByID(id string) (*models.Post, error) {
	
	post, err := s.Repo.GetPostByID(id)
	if err != nil {
		return nil, err
	}
	return post, nil
}

func (s *PostService) GetAllPosts() ([]models.Post, error) {
	post, err := s.Repo.GetAllPosts()
	if err != nil {
		return nil, err
	}
	return post, nil
}

func (s *PostService) DeletePostByID(id string) error {
	err := s.Repo.DeletePostByID(id)
	if err != nil {
		return err
	}
	return nil
}