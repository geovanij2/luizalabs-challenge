package entity

type Product struct {
	Id          string  `json:"id"`
	Price       uint64  `json:"price"`
	Image       string  `json:"image"`
	Brand       string  `json:"brand"`
	Title       string  `json:"title"`
	ReviewScore float64 `json:"reviewScore"`
}
