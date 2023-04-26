package service

import (
	model "rest-api/modules/v1/utilities/user/model"
	repo "rest-api/modules/v1/utilities/user/repository"
)

type IUserService interface {
	FindByID(id int) (model.Membership, error)
}

type userService struct {
	repository repo.IUserRepository
}

func NewMembershipService(repository repo.IUserRepository) *userService {
	return &userService{repository}
}

func (s *userService) FindByID(id int) (model.Membership, error) {
	membership, err := s.repository.FindByID(id)
	return membership, err
}
