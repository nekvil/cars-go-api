package service

import (
	"errors"
	"fmt"

	"github.com/nekvil/cars-go-api/internal/model"
	"github.com/nekvil/cars-go-api/internal/repository"
	"github.com/nekvil/cars-go-api/internal/utils"
	"gorm.io/gorm"
)

type ClientApiService struct {
	carRepo         repository.Car
	externalApiRepo repository.ClientApi
	peopleRepo      repository.People
}

func NewClientApiService(carRepo repository.Car, externalApiRepo repository.ClientApi, peopleRepo repository.People) *ClientApiService {
	return &ClientApiService{carRepo: carRepo, externalApiRepo: externalApiRepo, peopleRepo: peopleRepo}
}

func (s *ClientApiService) GetByRegNum(regNums []string) ([]model.Car, error) {
	var createdCars []model.Car

	for _, regNum := range regNums {
		utils.Logger.Debugf("Sending request to external API for car info with registration number '%s'", regNum)

		carInfo, err := s.externalApiRepo.GetByRegNum(regNum)
		if err != nil {
			utils.Logger.Errorf("Failed to get car info for registration number %s: %v", regNum, err)
			return createdCars, fmt.Errorf("failed to get car info for registration number %s: %w", regNum, err)
		}

		existingCar, err := s.carRepo.GetByRegNum(regNum)
		if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
			utils.Logger.Errorf("Failed to check car existence in the database: %v", err)
			return createdCars, fmt.Errorf("failed to check car existence in the database: %w", err)
		}

		if existingCar.ID != 0 {
			carInfo.ID = existingCar.ID
			if err := s.carRepo.Update(carInfo.ID, &carInfo); err != nil {
				utils.Logger.Errorf("Failed to update car info in the database: %v", err)
				return createdCars, fmt.Errorf("failed to update car info in the database: %w", err)
			}
			utils.Logger.Debugf("Car info updated in the database: %+v", carInfo)
		} else {
			existingOwner, err := s.peopleRepo.GetByFullName(carInfo.Owner.Name, carInfo.Owner.Surname, carInfo.Owner.Patronymic)
			if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
				utils.Logger.Errorf("Failed to check owner existence in the database: %v", err)
				return createdCars, fmt.Errorf("failed to check owner existence in the database: %w", err)
			}
			fmt.Println(existingOwner.ID)
			if existingOwner.ID != 0 {
				carInfo.Owner.ID = existingOwner.ID
			}

			if err := s.carRepo.Create(carInfo); err != nil {
				utils.Logger.Errorf("Failed to add car info to database: %v", err)
				return createdCars, fmt.Errorf("failed to add car info to database: %w", err)
			}
			utils.Logger.Debugf("Car added to database: %+v", carInfo)
		}

		createdCars = append(createdCars, carInfo)
	}

	utils.Logger.Infof("All cars added to the database successfully")
	return createdCars, nil
}
