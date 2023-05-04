package services

import (
	"errors"
	"log"

	model "rest-api/modules/v1/utilities/drivers/models"
	repo "rest-api/modules/v1/utilities/drivers/repository"
)

type IDriverService interface {
	FindAll(page int, limit int) ([]model.Driver, *int64, error)
	FindByID(id int) (model.Driver, error)
	Create(driverRequest model.DriverRequest) (*model.Driver, error)
	Update(id int, driverRequest model.DriverRequest) (*model.Driver, error, int)
	Delete(id int) (error, int)
}

type driverService struct {
	repository repo.IDriverRepository
}

func NewDriverService(repository repo.IDriverRepository) *driverService {
	return &driverService{repository}
}

func (s *driverService) FindAll(page int, limit int) ([]model.Driver, *int64, error) {
	pageA := page - 1
	if page == 0 {
		pageA = 0
	}
	skip := pageA * limit
	driver, count, err := s.repository.FindAll(limit, skip)

	return driver, count, err
}

func (s *driverService) FindByID(id int) (model.Driver, error) {
	driver, err := s.repository.FindByID(id)
	return driver, err
}

func (s *driverService) Create(DriverRequest model.DriverRequest) (*model.Driver, error) {

	Driver := model.Driver{
		DriverName:  DriverRequest.DriverName,
		Nik:         DriverRequest.Nik,
		PhoneNumber: DriverRequest.PhoneNumber,
		DailyCost:   DriverRequest.DailyCost,
	}
	newDriver, err := s.repository.Create(Driver)
	return &newDriver, err
}

func (s *driverService) Update(id int, DriverRequest model.DriverRequest) (*model.Driver, error, int) {

	Driver, errGet := s.repository.FindByID(id)
	if errGet != nil {
		return nil, errGet, 500
	}

	if Driver.ID == 0 {
		return nil, errors.New("ID Not Found"), 404
	}

	Driver.DriverName = DriverRequest.DriverName
	Driver.Nik = DriverRequest.Nik
	Driver.PhoneNumber = DriverRequest.PhoneNumber
	Driver.DailyCost = DriverRequest.DailyCost

	newDriver, err := s.repository.Update(Driver)
	if err != nil {
		return nil, nil, 500
	}
	return &newDriver, nil, 200
}

func (s *driverService) Delete(id int) (error, int) {

	Driver, errGet := s.repository.FindByID(id)
	if errGet != nil {
		return errGet, 500
	}
	if Driver.ID == 0 {
		return nil, 404
	}
	log.Print("Driver", Driver)

	_, err := s.repository.Delete(Driver)
	if err != nil {
		return err, 500
	}
	return nil, 200
}
