package produk

import (
	"context"
	"dungeonSnackBE/config"
	"dungeonSnackBE/model"
	"encoding/json"
	"net/http"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// GetProduk untuk mendapatkan semua produk
func GetProduk(w http.ResponseWriter, r *http.Request) {
	// Koneksi ke koleksi MongoDB produk
	collection := config.Mongoconn.Collection("produk")

	// Membuat context dengan timeout
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Membuat variabel untuk menampung semua produk
	var produks []model.Menu

	// Membuat filter kosong
	filter := map[string]interface{}{}

	// Membuat variabel untuk menampung semua produk
	cursor, err := collection.Find(ctx, filter)
	if err != nil {
		http.Error(w, "Failed to fetch produk", http.StatusInternalServerError)
		return
	}
	defer cursor.Close(ctx)

	// Mengambil semua data produk
	for cursor.Next(ctx) {
		var produk model.Menu
		cursor.Decode(&produk)
		produks = append(produks, produk)
	}

	// Mengirimkan respons sukses
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(produks)
}

// GetProdukByID untuk mendapatkan produk berdasarkan ID
func GetProdukByID(w http.ResponseWriter, r *http.Request) {
	// Mendapatkan ID produk dari parameter URL
	id := r.URL.Query().Get("id")

	// Memastikan ID tidak kosong
	if id == "" {
		http.Error(w, "ID is required", http.StatusBadRequest)
		return
	}

	// Mengubah ID dari string ke ObjectID
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		http.Error(w, "Invalid ID format", http.StatusBadRequest)
		return
	}

	// Koneksi ke koleksi MongoDB produk
	collection := config.Mongoconn.Collection("produk")

	// Membuat context dengan timeout
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Membuat variabel untuk menampung produk
	var produk model.Menu

	// Melakukan pencarian berdasarkan ID
	err = collection.FindOne(ctx, model.Menu{ID: objID}).Decode(&produk)
	if err != nil {
		http.Error(w, "produk not found", http.StatusNotFound)
		return
	}

	// Mengirimkan respons sukses
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(produk)
}

// GetProdukBySlug untuk mendapatkan produk berdasarkan slug
func GetProdukBySlug(w http.ResponseWriter, r *http.Request) {
	// Mendapatkan slug produk dari parameter URL
	slug := r.URL.Query().Get("slug")

	// Memastikan slug tidak kosong
	if slug == "" {
		http.Error(w, "Slug is required", http.StatusBadRequest)
		return
	}

	// Koneksi ke koleksi MongoDB produk
	collection := config.Mongoconn.Collection("produk")

	// Membuat context dengan timeout
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Membuat variabel untuk menampung produk
	var produk model.Menu

	// Melakukan pencarian berdasarkan slug
	err := collection.FindOne(ctx, model.Menu{Slug: slug}).Decode(&produk)
	if err != nil {
		http.Error(w, "produk not found", http.StatusNotFound)
		return
	}

	// Mengirimkan respons sukses
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(produk)
}

// GetProdukByCategory untuk mendapatkan produk berdasarkan kategori
func GetProdukByCategory(w http.ResponseWriter, r *http.Request) {
	// Mendapatkan kategori produk dari parameter URL
	category := r.URL.Query().Get("category")

	// Memastikan kategori tidak kosong
	if category == "" {
		http.Error(w, "Category is required", http.StatusBadRequest)
		return
	}

	// Koneksi ke koleksi MongoDB produk
	collection := config.Mongoconn.Collection("produk")

	// Membuat context dengan timeout
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Membuat variabel untuk menampung semua produk
	var produks []model.Menu

	// Membuat filter berdasarkan kategori
	filter := map[string]interface{}{
		"category": category,
	}

	// Membuat variabel untuk menampung semua produk
	cursor, err := collection.Find(ctx, filter)
	if err != nil {
		http.Error(w, "Failed to fetch produk", http.StatusInternalServerError)
		return
	}
	defer cursor.Close(ctx)

	// Mengambil semua data produk
	for cursor.Next(ctx) {
		var produk model.Menu
		cursor.Decode(&produk)
		produks = append(produks, produk)
	}

	// Mengirimkan respons sukses
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(produks)
}

// GetProdukByToko untuk mendapatkan produk berdasarkan toko
func GetProdukByToko(w http.ResponseWriter, r *http.Request) {
	// Mendapatkan toko produk dari parameter URL
	toko := r.URL.Query().Get("toko")

	// Memastikan toko tidak kosong
	if toko == "" {
		http.Error(w, "Toko is required", http.StatusBadRequest)
		return
	}

	// Koneksi ke koleksi MongoDB produk
	collection := config.Mongoconn.Collection("produk")

	// Membuat context dengan timeout
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Membuat variabel untuk menampung semua produk
	var produks []model.Menu

	// Membuat filter berdasarkan toko
	filter := map[string]interface{}{
		"toko": toko,
	}

	// Membuat variabel untuk menampung semua produk
	cursor, err := collection.Find(ctx, filter)
	if err != nil {
		http.Error(w, "Failed to fetch produk", http.StatusInternalServerError)
		return
	}
	defer cursor.Close(ctx)

	// Mengambil semua data produk
	for cursor.Next(ctx) {
		var produk model.Menu
		cursor.Decode(&produk)
		produks = append(produks, produk)
	}

	// Mengirimkan respons sukses
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(produks)
}

// GetProdukByPrice untuk mendapatkan produk berdasarkan harga
func GetProdukByPrice(w http.ResponseWriter, r *http.Request) {
	// Mendapatkan harga produk dari parameter URL
	price := r.URL.Query().Get("price")

	// Memastikan harga tidak kosong
	if price == "" {
		http.Error(w, "Price is required", http.StatusBadRequest)
		return
	}

	// Koneksi ke koleksi MongoDB produk
	collection := config.Mongoconn.Collection("produk")

	// Membuat context dengan timeout
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Membuat variabel untuk menampung semua produk
	var produks []model.Menu

	// Membuat filter berdasarkan harga
	filter := map[string]interface{}{
		"price": price,
	}

	// Membuat variabel untuk menampung semua produk
	cursor, err := collection.Find(ctx, filter)
	if err != nil {
		http.Error(w, "Failed to fetch produk", http.StatusInternalServerError)
		return
	}
	defer cursor.Close(ctx)

	// Mengambil semua data produk
	for cursor.Next(ctx) {
		var produk model.Menu
		cursor.Decode(&produk)
		produks = append(produks, produk)
	}

	// Mengirimkan respons sukses
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(produks)
}