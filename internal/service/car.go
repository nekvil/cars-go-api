package service

import (
	"errors"
	"fmt"

	"github.com/nekvil/cars-go-api/internal/model"
	"github.com/nekvil/cars-go-api/internal/repository"
	"github.com/nekvil/cars-go-api/internal/utils"
	"gorm.io/gorm"
)

type CarService struct {
	carRepo    repository.Car
	peopleRepo repository.People
}

func NewCarService(carRepo repository.Car, peopleRepo repository.People) *CarService {
	return &CarService{carRepo: carRepo, peopleRepo: peopleRepo}
}

func (s *CarService) Create(car model.Car) error {
	utils.Logger.Debugf("Creating car: %+v", car)

	err := s.carRepo.Create(car)
	if err != nil {
		utils.Logger.Debugf("Failed to create car: %v", err)
		return fmt.Errorf("failed to create car: %v", err)
	}

	utils.Logger.Debugf("Car created successfully: %+v", car)

	return nil
}

func (s *CarService) GetAll(filter string, page int) ([]model.Car, error) {
	utils.Logger.Debugf("Getting cars with filter '%s' and page '%d'", filter, page)

	cars, err := s.carRepo.GetAll(filter, page)
	if err != nil {
		return nil, err
	}
	utils.Logger.Debugf("Retrieved cars with filter '%s' and page '%d'", filter, page)

	return cars, nil
}

func (s *CarService) GetById(carId int) (model.Car, error) {
	utils.Logger.Debugf("Getting car by ID '%d'", carId)

	existingCar, err := s.carRepo.GetById(carId)
	if err != nil {
		return model.Car{}, fmt.Errorf("failed to get car by ID: %w", err)
	}

	utils.Logger.Debugf("Retrieved car by ID '%d'", carId)

	return existingCar, nil
}

func (s *CarService) Delete(carId int) error {
	utils.Logger.Debugf("Deleting car with ID '%d'", carId)

	_, err := s.carRepo.GetById(carId)
	if err != nil {
		return fmt.Errorf("failed to delete car: %w", err)
	}

	err = s.carRepo.Delete(carId)
	if err != nil {
		return fmt.Errorf("failed to delete car: %w", err)
	}

	utils.Logger.Infof("Car with ID '%d' deleted successfully", carId)

	return nil
}

func (s *CarService) Update(carId int, carData *model.Car) error {
	utils.Logger.Debugf("Updating car with ID '%d'", carId)

	existingCar, err := s.carRepo.GetById(carId)
	if err != nil {
		return fmt.Errorf("failed to update car: %v", err)
	}

	if existingCar.RegNum != carData.RegNum {
		existingCarByRegNum, err := s.carRepo.GetByRegNum(carData.RegNum)
		if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
			return fmt.Errorf("failed to update car: %v", err)
		}
		if existingCarByRegNum.ID != 0 {
			return errors.New("car with this registration number already exists")
		}
	}

	existingOwner, err := s.peopleRepo.GetByFullName(carData.Owner.Name, carData.Owner.Surname, carData.Owner.Patronymic)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return fmt.Errorf("failed to update car: %v", err)
	}

	if existingOwner.ID == 0 && carData.Owner.Name != "" && carData.Owner.Surname != "" && carData.Owner.Patronymic != "" {
		newOwner := model.People{
			Name:       carData.Owner.Name,
			Surname:    carData.Owner.Surname,
			Patronymic: carData.Owner.Patronymic,
		}
		createdOwner, err := s.peopleRepo.Create(newOwner)
		if err != nil {
			return fmt.Errorf("failed to create owner: %v", err)
		}
		carData.OwnerID = createdOwner.ID
	} else {
		carData.OwnerID = existingOwner.ID
	}

	if err := s.carRepo.Update(carId, carData); err != nil {
		return fmt.Errorf("failed to update car: %v", err)
	}

	utils.Logger.Infof("Car with ID '%d' updated successfully", carId)

	return nil
}

func (s *CarService) GetByRegNum(regNum string) (model.Car, error) {
	utils.Logger.Debugf("Getting car by registration number '%s'", regNum)

	car, err := s.carRepo.GetByRegNum(regNum)
	if err != nil {
		return model.Car{}, fmt.Errorf("failed to get car by registration number: %w", err)
	}

	utils.Logger.Debugf("Retrieved car by registration number '%s'", regNum)

	return car, nil
}
