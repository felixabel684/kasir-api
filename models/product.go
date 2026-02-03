package models

type Product struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Price int    `json:"price"`
	Stock int    `json:"stock"`
	CategoryID int    `json:"category_id"`      // ID kategori
	Category   string `json:"category_name,omitempty"` // optional, akan diisi saat join
}