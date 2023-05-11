package service

import (
	"fmt"
	model "rest-api/modules/v1/utilities/user/models"
	repo "rest-api/modules/v1/utilities/user/repository"
)

type IUserService interface {
	GetById(id int) (model.Users, error)
	Create(userRequest model.UserRequest) (*model.Users, error)
}

type userService struct {
	repository repo.IUserRepository
}

func NewUserService(repository repo.IUserRepository) *userService {
	return &userService{repository}
}

func (service *userService) GetById(id int) (model.Users, error) {
	user, err := service.repository.GetById(id)

	return user, err
}

func (service *userService) Create(userRequest model.UserRequest) (*model.Users, error) {

	fmt.Println("serv", userRequest)

	User := model.Users{
		Name:     userRequest.Name,
		Email:    userRequest.Email,
		Password: userRequest.Password,
		RoleId:   userRequest.RoleId,
	}

	fmt.Println("serv2", User)

	newUser, err := service.repository.Create(User)
	return &newUser, err

}
