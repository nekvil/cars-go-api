package repository

import (
	"github.com/nekvil/cars-go-api/internal/model"
	"gorm.io/gorm"
)

type Car interface {
	Create(car model.Car) error
	GetAll(filter string, page int) ([]model.Car, error)
	GetById(carID int) (model.Car, error)
	Delete(carID int) error
	Update(carID int, carData *model.Car) error
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
	GetByRegNum(regNum string) (model.Car, error)
}

type Repository struct {
	Car
	People
	ClientApi
}

func NewRepository(db *gorm.DB, baseURL string) *Repository {
	return &Repository{
		Car:       NewCarPostgres(db),
		People:    NewPeoplePostgres(db),
		ClientApi: NewClientApiRepository(baseURL),
	}
}
