package repository

import (
	"fmt"

	"github.com/nekvil/cars-go-api/internal/model"
	"gorm.io/gorm"
)

type PeoplePostgres struct {
	db *gorm.DB
}

func NewPeoplePostgres(db *gorm.DB) *PeoplePostgres {
	return &PeoplePostgres{db: db}
}

func (r *PeoplePostgres) Create(people model.People) (model.People, error) {
	if err := r.db.Create(&people).Error; err != nil {
		return model.People{}, err
	}
	return people, nil
}

func (r *PeoplePostgres) GetAll(id int) ([]model.People, error) {
	var people []model.People
	if err := r.db.Find(&people).Error; err != nil {
		return nil, fmt.Errorf("failed to get all people: %w", err)
	}
	return people, nil
}

func (r *PeoplePostgres) GetById(id int) (model.People, error) {
	var person model.People
	if err := r.db.First(&person, id).Error; err != nil {
		return model.People{}, fmt.Errorf("failed to get person by ID: %w", err)
	}
	return person, nil
}

func (r *PeoplePostgres) Delete(id int) error {
	return r.db.Delete(&model.People{}, id).Error
}

func (r *PeoplePostgres) Update(id int, peopleData *model.People) error {
	return r.db.Model(&model.People{}).Where("id = ?", id).Updates(peopleData).Error
}

func (r *PeoplePostgres) GetByFullName(name, surname, patronymic string) (model.People, error) {
	var owner model.People
	if err := r.db.Where("name = ? AND surname = ? AND patronymic = ?", name, surname, patronymic).First(&owner).Error; err != nil {
		return model.People{}, err
	}

	return owner, nil
}
