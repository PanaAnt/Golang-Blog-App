package services

import (
	"blogApp/middleware"
	"blogApp/models"
	"blogApp/repository"
	"blogApp/utils"
	"errors"
	"fmt"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type UserService struct {
	Repo repository.UserRepository
}

func (s *UserService) RegisterUser(req *models.User) error {

	//check if they exist
	existingUser, err := s.Repo.GetUserByUsername(req.Username)
	if err == nil && existingUser != nil {
		return fmt.Errorf("user already exists")
	}

	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
    return err 
	} 

	//hashing the password
	hashedPass, err := utils.HashPassword(req.Password)
	if err != nil {
		return err
	}
	req.Password = hashedPass

	myuuid := uuid.New()
	req.ID = myuuid

	//put the users into the database
	err = s.Repo.CreateUser(req)
	if err != nil {
		return err
	}
	return nil
}

func (s *UserService) Login(req *models.User) (string, *models.User, error) {
	// check if user exists
	user, err := s.Repo.GetUserByUsername(req.Username)
	if err != nil {
		return "", &models.User{}, err 
	}

	//compare passwords
	err = utils.ComparePassword(user.Password, req.Password)
	if err != nil {
		return "", &models.User{}, err
	}

	//generate token
	token, err := middleware.GenerateJWT(user.ID.String())
	if err != nil {
		return "", &models.User{}, err
	}
	return token, user, nil
}