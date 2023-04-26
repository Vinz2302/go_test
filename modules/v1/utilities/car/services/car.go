package services

import (
	"errors"
	"log"
	model "rest-api/modules/v1/utilities/car/models"
	repo "rest-api/modules/v1/utilities/car/repository"
)

type ICarService interface {
	FindAll(page int, limit int) ([]model.Car, *int64, error)
	FindByID(id int) (model.Car, error)
	Create(carRequest model.CarRequest) (*model.Car, error)
	Update(id int, carRequest model.CarRequest) (*model.Car, error, int)
	Delete(id int) (error, int)
}

type carService struct {
	repository repo.ICarRepository
}

func NewCarService(repository repo.ICarRepository) *carService {
	return &carService{repository}
}

func (s *carService) FindAll(page int, limit int) ([]model.Car, *int64, error) {
	pageA := page - 1
	if page == 0 {
		pageA = 0
	}
	skip := pageA * limit
	car, count, err := s.repository.FindAll(limit, skip)

	return car, count, err
}

func (s *carService) FindByID(id int) (model.Car, error) {
	car, err := s.repository.FindByID(id)
	return car, err
}

func (s *carService) Create(CarRequest model.CarRequest) (*model.Car, error) {

	Car := model.Car{
		CarName:        CarRequest.CarName,
		RentDailyPrice: CarRequest.RentDailyPrice,
		Stock:          *CarRequest.Stock,
	}
	newCar, err := s.repository.Create(Car)
	return &newCar, err
}

func (s *carService) Update(id int, CarRequest model.CarRequest) (*model.Car, error, int) {

	Car, errGet := s.repository.FindByID(id)
	if errGet != nil {
		return nil, errGet, 500
	}
	if Car.ID == 0 {
		return nil, errors.New("ID Not Found"), 404
	}

	Car.CarName = CarRequest.CarName
	Car.RentDailyPrice = CarRequest.RentDailyPrice
	Car.Stock = *CarRequest.Stock

	newCar, err := s.repository.Update(Car)
	if err != nil {
		return nil, nil, 500
	}
	return &newCar, nil, 200
}

func (s *carService) Delete(id int) (error, int) {

	Car, errGet := s.repository.FindByID(id)
	if errGet != nil {
		return errGet, 500
	}
	if Car.ID == 0 {
		return nil, 404
	}
	log.Print("Car", Car)

	_, err := s.repository.Delete(Car)
	if err != nil {
		return err, 500
	}
	return nil, 200
}
