package models

// Produk represents a product in the cashier system
type Product struct {
	ID    int    `json:"id"`
	Name  string `json:"nama"`
	Price int    `json:"harga"`
	Stock int    `json:"stok"`
}
