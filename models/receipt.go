package models

type Receipt struct {
	ID          uint       `gorm:"primaryKey"`
	Date        string     `gorm:"not null"`
	PatientName string     `gorm:"not null;size:500"`
	FilePath    string     `gorm:"size:500"`
	Medicines   []Medicine `gorm:"many2many:receipt_medicines;"`
}
