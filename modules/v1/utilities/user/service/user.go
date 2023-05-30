package service

import (
	"errors"
	"fmt"
	model "rest-api/modules/v1/utilities/user/models"
	repo "rest-api/modules/v1/utilities/user/repository"
	helper "rest-api/pkg/helpers"
)

type IUserService interface {
	GetById(id int) (model.Users, error)
	Create(userRequest model.UserRequest) (*model.Users, error)
	Login(loginRequest model.UserLogin, refresh_token string) (*model.Users, error)
	Refresh(email string, cookie string) (*model.Users, error, int)
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

	userRequest.Password, _ = helper.Hash(userRequest.Password)
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

func (service *userService) Login(loginRequest model.UserLogin, refresh_token string) (*model.Users, error) {

	Login := model.UserLogin{
		Email:    loginRequest.Email,
		Password: loginRequest.Password,
	}

	NewLogin, err := service.repository.Login(Login, refresh_token)

	return NewLogin, err

}

func (service *userService) Refresh(email string, cookie string) (*model.Users, error, int) {

	NewRefresh, errRefresh := service.repository.Refresh(email, cookie)
	if errRefresh != nil {
		return nil, errRefresh, 500
	}

	if NewRefresh.Refresh_token != cookie {
		return nil, errors.New("Invalid Refresh Token"), 400
	}

	return NewRefresh, nil, 200
}
