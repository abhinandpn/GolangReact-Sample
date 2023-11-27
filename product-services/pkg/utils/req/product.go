package req

type ReqProduct struct {
	ProductName string  `json:"productname"`
	Discription string  `json:"discription"`
	Price       float64 `json:"price"`
	Category    string  `json:"category"`
}
