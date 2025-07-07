package models

type Diary struct {
	UserID  uint   `gorm:"primaryKey"`
	Date    string `gorm:"primaryKey;size:10"`
	Content string `gorm:"type:text"`
	Feels   uint8  `gorm:"not null"`
}
