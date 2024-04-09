package repository

import (
	"errors"
	"fmt"

	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/nekvil/cars-go-api/internal/model"
	"gorm.io/gorm"
)

type CarPostgres struct {
	db *gorm.DB
}

func NewCarPostgres(db *gorm.DB) *CarPostgres {
	return &CarPostgres{db: db}
}

func (r *CarPostgres) Create(car model.Car) error {
	return r.db.Create(&car).Error
}

func (r *CarPostgres) GetAll(filter string, page int) ([]model.Car, error) {
	const pageSize = 10
	offset := (page - 1) * pageSize

	query := `
		SELECT *
		FROM car`

	if filter != "" {
		query += fmt.Sprintf(`
		WHERE to_tsvector(reg_num || ' ' || mark || ' ' || model || ' ' || year::text)
			@@ to_tsquery('%s:*')`, filter)
	}

	query += fmt.Sprintf(`
		ORDER BY id
		OFFSET %d
		LIMIT %d`, offset, pageSize)

	var cars []model.Car
	if err := r.db.Preload("Owner").Raw(query).Find(&cars).Error; err != nil {
		return nil, fmt.Errorf("failed to get cars: %w", err)
	}
	return cars, nil
}

func (r *CarPostgres) GetById(id int) (model.Car, error) {
	var existingCar model.Car
	if err := r.db.First(&existingCar, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return model.Car{}, fmt.Errorf("car with ID %d not found", id)
		}
		return model.Car{}, fmt.Errorf("failed to check if car exists: %w", err)
	}
	return existingCar, nil
}

func (r *CarPostgres) Delete(id int) error {
	return r.db.Delete(&model.Car{}, id).Error
}

func (r *CarPostgres) Update(id int, car *model.Car) error {
	var existingCar model.Car
	if err := r.db.Model(&existingCar).Where("id = ?", id).Updates(car).Error; err != nil {
		return fmt.Errorf("failed to update car: %w", err)
	}
	return nil
}

func (r *CarPostgres) GetByRegNum(regNum string) (model.Car, error) {
	var car model.Car
	if err := r.db.Where("reg_num = ?", regNum).First(&car).Error; err != nil {
		return model.Car{}, fmt.Errorf("failed to get car by registration number: %w", err)
	}
	return car, nil
}
