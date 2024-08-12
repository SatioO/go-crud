package usecase

import (
	"errors"

	"github.com/satioO/basics/v2/models"
)

type UserService interface {
	FindUsers() []models.User
	FindUserById(userId int) (*models.User, error)
	CreateUser(body *models.CreateUserRequest) (*models.User, error)
	UpdateUser(userId int) (*models.User, error)
	DeleteUser(userId int) error
}

type userServiceImpl struct{}

func NewUserService() UserService {
	return &userServiceImpl{}
}

func (u *userServiceImpl) FindUsers() []models.User {
	users := []models.User{}
	return users
}

func (u *userServiceImpl) FindUserById(userId int) (*models.User, error) {
	return nil, errors.New("Not Implemented")
}

func (u *userServiceImpl) CreateUser(body *models.CreateUserRequest) (*models.User, error) {
	return nil, errors.New("Not Implemented")
}

func (u *userServiceImpl) UpdateUser(userId int) (*models.User, error) {
	return nil, errors.New("Not Implemented")
}

func (u *userServiceImpl) DeleteUser(userId int) error {
	return errors.New("Not Implemented")
}
