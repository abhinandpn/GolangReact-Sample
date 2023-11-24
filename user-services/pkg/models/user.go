package models

type User struct {
	Id       uint   `json:"id" gorm:"primaryKey;unique;not null"`
	UserName string `json:"username" gorm:"not null"`
	Email    string `json:"email" gorm:"not null;unique"`
	Password string `json:"password" gorm:"not null"`
	IsBlock  bool
}
