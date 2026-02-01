package main

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"
)

// Model Kategori
type Category struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

// Data Dummy Kategori
var categories = []Category{
	{ID: 1, Name: "Makanan", Description: "Category Makanan"},
	{ID: 2, Name: "Minuman", Description: "Category Minuman"},
}
var lastCategoryID = 2

func getCategoryByID(w http.ResponseWriter, r *http.Request) {
	// Parse ID dari URL path
	// URL: /categories/123 -> ID = 123
	idStr := strings.TrimPrefix(r.URL.Path, "/categories/")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid Category ID", http.StatusBadRequest)
		return
	}

	// Cari category dengan ID tersebut
	for _, c := range categories {
		if c.ID == id {
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(c)
			return
		}
	}

	// Kalau tidak found
	http.Error(w, "category belum ada", http.StatusNotFound)
}

// PUT localhost:8080/categories/{id}
func updateCategory(w http.ResponseWriter, r *http.Request) {
	// get id dari request
	idStr := strings.TrimPrefix(r.URL.Path, "/categories/")

	// ganti int
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid category ID", http.StatusBadRequest)
		return
	}

	// get data dari request
	var updateCat Category
	err = json.NewDecoder(r.Body).Decode(&updateCat)
	if err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	// loop category, cari id, ganti sesuai data dari request
	for i := range categories {
		if categories[i].ID == id {
			updateCat.ID = id
			categories[i] = updateCat

			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(updateCat)
			return
		}
	}

	http.Error(w, "category belum ada", http.StatusNotFound)
}

// Handler untuk /categories
func categoriesHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		// Ambil semua kategori
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(categories)
	} else if r.Method == "POST" {
		// Tambah kategori
		var kBaru Category
		if err := json.NewDecoder(r.Body).Decode(&kBaru); err != nil {
			http.Error(w, "Invalid request", http.StatusBadRequest)
			return
		}
		lastCategoryID++
		kBaru.ID = lastCategoryID
		categories = append(categories, kBaru)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(kBaru)
	}
}

func deleteCategory(w http.ResponseWriter, r *http.Request) {
	// get id
	idStr := strings.TrimPrefix(r.URL.Path, "/categories/")

	// ganti id int
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid category ID", http.StatusBadRequest)
		return
	}

	// loop category cari ID, dapet index yang mau dihapus
	for i, c := range categories {
		if c.ID == id {
			// bikin slice baru dengan data sebelum dan sesudah index
			categories = append(categories[:i], categories[i+1:]...)

			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(map[string]string{
				"message": "sukses delete",
			})
			return
		}
	}

	http.Error(w, "category belum ada", http.StatusNotFound)
}

// Handler untuk /categories/{id}
func categoryByIDHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		getCategoryByID(w, r)
	} else if r.Method == "PUT" {
		updateCategory(w, r)
	} else if r.Method == "DELETE" {
		deleteCategory(w, r)
	}
}
