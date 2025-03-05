package models

type Information struct {
	ID    uint   `gorm:"primaryKey"`
	Key   string `gorm:"not null;unique"`
	Value string `gorm:"not null"`
}
