package models

type Product struct {
<<<<<<< HEAD
	Id        uint   `json:"id" gorm:"primaryKey;unique;not null"`
	UserName  string `json:"username" gorm:"not null"`
	Email     string `json:"email" gorm:"not null;unique"`
	Password  string `json:"password" gorm:"not null"`
	IsBlocked bool
=======
	Id          uint   `json:"id" gorm:"primaryKey;unique;not null"`
	ProductName string `json:"productname" gorm:"not null"`
	Discription string `json:"discription" gorm:"not null;unique"`
	Image       string `json:"image" gorm:"not null"`
	Stock       uint   `json:"stock" gorm:"not null"`
>>>>>>> user
}
