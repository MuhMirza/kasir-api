package main

import (
	"fmt"
	"net/http"
)

func main() {
	// Routing Produk (Logic-nya ada di file produk.go)
	// Asumsi kamu sudah memindahkan logic produk ke file produk.go
	http.HandleFunc("/api/produk", produkHandler)
	http.HandleFunc("/api/produk/", produkByIDHandler)

	// Routing Kategori (Logic-nya ada di file kategori.go)
	http.HandleFunc("/categories", categoriesHandler)
	http.HandleFunc("/categories/", categoryByIDHandler)

	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("OK"))
	})

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("gagal running server")
	}
}
