package models

type Medicine struct {
	ID           uint `gorm:"primaryKey"`
	MedicineID   uint `gorm:"index"`
	HowToUse     string
	Caliber      string `gorm:"size:250"`
	MedicineName string `gorm:"size:500"`
}

type ReceiptMedicine struct {
	ReceiptID  uint `gorm:"primaryKey"`
	MedicineID uint `gorm:"primaryKey"`
}
