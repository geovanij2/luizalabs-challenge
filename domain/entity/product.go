package entity

type Product struct {
	Id          string  `json:"id" db:"product_id"`
	Price       uint64  `json:"price" db:"price"`
	Image       string  `json:"image" db:"image"`
	Brand       string  `json:"brand" db:"brand"`
	Title       string  `json:"title" db:"title"`
	ReviewScore float64 `json:"reviewScore" db:"review_score"`
}
