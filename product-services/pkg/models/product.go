package models

type Product struct {
	Id          uint   `json:"id" gorm:"primaryKey;unique;not null"`
	ProductName string `json:"productname" gorm:"not null"`
	Discription string `json:"discription" gorm:"not null;unique"`
	Image       string `json:"image" gorm:"not null"`
	Stock       uint   `json:"stock" gorm:"not null"`
}
