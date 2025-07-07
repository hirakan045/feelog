package models

type User struct {
	ID           uint   `gorm:"primaryKey;autoIncrement"`
	Name         string `gorm:"size:100;not null"`
	Email        string `gorm:"size:255;unique;not null"`
	PasswordHash string `gorm:"size:255;not null"`
}
