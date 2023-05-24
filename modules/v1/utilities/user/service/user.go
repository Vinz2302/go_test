package service

import (
	"fmt"
	model "rest-api/modules/v1/utilities/user/models"
	repo "rest-api/modules/v1/utilities/user/repository"
	helper "rest-api/pkg/helpers"
)

type IUserService interface {
	GetById(id int) (model.Users, error)
	Create(userRequest model.UserRequest) (*model.Users, error)
	Login(loginRequest model.UserLogin) (*model.Users, error)
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

func (service *userService) Login(loginRequest model.UserLogin) (*model.Users, error) {

	Login := model.UserLogin{
		Email:    loginRequest.Email,
		Password: loginRequest.Password,
	}

	NewLogin, err := service.repository.Login(Login)

	/* compareErr := helper.CompareHash([]byte(NewLogin.Password), []byte(loginRequest.Password))
	if compareErr != nil {
		fmt.Println("Invalid credentials")
		return nil, err
	} */

	//jwt.GenerateToken(NewLogin.Email, NewLogin.Password, NewLogin.RoleId)

	return NewLogin, err

}
