package model

type Car struct {
	ID      int    `json:"-" gorm:"primaryKey"`
	RegNum  string `json:"regNum" gorm:"uniqueIndex"`
	Mark    string `json:"mark"`
	Model   string `json:"model"`
	Year    int    `json:"year"`
	OwnerID int    `json:"-"`
	Owner   People `json:"owner" gorm:"foreignKey:OwnerID;references:ID"`
}
