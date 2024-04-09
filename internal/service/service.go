package service

import (
	"github.com/nekvil/cars-go-api/internal/model"
	"github.com/nekvil/cars-go-api/internal/repository"
)

type Car interface {
	Create(car model.Car) error
	GetAll(filter string, page int) ([]model.Car, error)
	GetById(carId int) (model.Car, error)
	Delete(carId int) error
	Update(carId int, carData *model.Car) error
	GetByRegNum(regNum string) (model.Car, error)
}

type People interface {
	Create(people model.People) (model.People, error)
	GetAll(peopleId int) ([]model.People, error)
	GetById(peopleId int) (model.People, error)
	Delete(peopleId int) error
	Update(peopleId int, peopleData *model.People) error
	GetByFullName(name, surname, patronymic string) (model.People, error)
}

type ClientApi interface {
	GetByRegNum(regNum []string) ([]model.Car, error)
}

type Service struct {
	Car
	People
	ClientApi
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Car:       NewCarService(repos.Car, repos.People),
		People:    NewPeopleService(repos.People),
		ClientApi: NewClientApiService(repos.Car, repos.ClientApi, repos.People),
	}
}
