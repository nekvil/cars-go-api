package model

type People struct {
	ID         int    `json:"-" gorm:"primaryKey"`
	Name       string `json:"name"`
	Surname    string `json:"surname"`
	Patronymic string `json:"patronymic"`
}
