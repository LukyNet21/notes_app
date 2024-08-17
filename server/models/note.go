package models

type Note struct {
	ID      uint   `json:"id" gorm:"primaryKey"`
	Name    string `json:"note_name" gorm:"not null"`
	Content string `json:"note_content" gorm:"not null"`
	UserID  uint
}
