package models

// @description Represents a product object
type Product struct {
	ID          uint64  `json:"id"`
	ProductName string  `json:"product_name"`
	DrugGeneric string  `json:"drug_generic"`
	Company     string  `json:"company"`
	Score       float64 `json:"score"`
}
