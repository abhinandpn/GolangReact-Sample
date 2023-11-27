package models

type Product struct {
	Id          uint    `json:"id"`
	ProductName string  `json:"porductname"`
	Discription string  `json:"description"`
	Price       float64 `json:"price"`
	Category    string  `json:"category"`
}
