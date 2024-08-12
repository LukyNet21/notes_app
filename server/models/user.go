package models

type User struct {
	ID        uint   `json:"id" gorm:"primaryKey"`
	Username  string `json:"username" gorm:"unique;not null;min=3;max=30"`
	Email     string `json:"email" gorm:"unique;not null"`
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
	Password  string `json:"password" gorm:"unique;not null"`
}
