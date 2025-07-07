package models

type Chat struct {
	UserID  uint   `gorm:"primaryKey"`
	Date    string `gorm:"primaryKey;size:10"`
	Branch  uint   `gorm:"primaryKey"`
	Content string `gorm:"type:text;not null"`
	IsAI    bool   `gorm:"default:false"`
}
