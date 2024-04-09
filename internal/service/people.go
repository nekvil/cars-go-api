package service

import (
	"fmt"

	"github.com/nekvil/cars-go-api/internal/model"
	"github.com/nekvil/cars-go-api/internal/repository"
	"github.com/nekvil/cars-go-api/internal/utils"
)

type PeopleService struct {
	peopleRepo repository.People
}

func NewPeopleService(peopleRepo repository.People) *PeopleService {
	return &PeopleService{peopleRepo: peopleRepo}
}

func (s *PeopleService) Create(people model.People) (model.People, error) {
	utils.Logger.Debugf("Creating owner: %+v", people)
	createdOwner, err := s.peopleRepo.Create(people)
	if err != nil {
		utils.Logger.Debugf("Failed to create owner: %v", err)
		return createdOwner, fmt.Errorf("failed to create owner: %w", err)
	}
	utils.Logger.Debugf("Owner created successfully: %+v", people)
	return createdOwner, nil
}

func (s *PeopleService) GetAll(peopleId int) ([]model.People, error) {
	utils.Logger.Debugf("Getting all owners")
	peopleList, err := s.peopleRepo.GetAll(peopleId)
	if err != nil {
		utils.Logger.Debugf("Failed to get all owners: %v", err)
		return nil, fmt.Errorf("failed to get all owners: %w", err)
	}
	utils.Logger.Debugf("Retrieved all owners successfully")
	return peopleList, nil
}

func (s *PeopleService) GetById(peopleId int) (model.People, error) {
	utils.Logger.Debugf("Getting owner by ID: %d", peopleId)
	owner, err := s.peopleRepo.GetById(peopleId)
	if err != nil {
		return model.People{}, fmt.Errorf("failed to get owner by ID: %w", err)
	}
	utils.Logger.Debugf("Retrieved owner by ID %d successfully: %+v", peopleId, owner)
	return owner, nil
}

func (s *PeopleService) Delete(peopleId int) error {
	utils.Logger.Debugf("Deleting owner with ID: %d", peopleId)
	err := s.peopleRepo.Delete(peopleId)
	if err != nil {
		return fmt.Errorf("failed to delete owner with ID %d: %w", peopleId, err)
	}
	utils.Logger.Debugf("Owner with ID %d deleted successfully", peopleId)
	return nil
}

func (s *PeopleService) Update(peopleId int, peopleData *model.People) error {
	utils.Logger.Debugf("Updating owner with ID: %d", peopleId)
	err := s.peopleRepo.Update(peopleId, peopleData)
	if err != nil {
		return fmt.Errorf("failed to update owner with ID %d: %w", peopleId, err)
	}
	utils.Logger.Debugf("Owner with ID %d updated successfully: %+v", peopleId, peopleData)
	return nil
}

func (s *PeopleService) GetByFullName(name, surname, patronymic string) (model.People, error) {
	utils.Logger.Debugf("Getting owner by name '%s', surname '%s', and patronymic '%s'", name, surname, patronymic)

	owner, err := s.peopleRepo.GetByFullName(name, surname, patronymic)
	if err != nil {
		return owner, fmt.Errorf("failed to get owner by name: %w", err)
	}

	utils.Logger.Debugf("Retrieved owner by name '%s', surname '%s', and patronymic '%s'", name, surname, patronymic)

	return owner, nil
}
