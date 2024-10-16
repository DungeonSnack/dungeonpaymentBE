package produk

import (
	"context"
	"dungeonSnackBE/config"
	"dungeonSnackBE/model"
	"encoding/json"
	"net/http"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// GetProduct untuk mendapatkan detail produk berdasarkan slug
func GetProduk(w http.ResponseWriter, r *http.Request) {
	// Mengambil slug dari query parameter
	slug := r.URL.Query().Get("slug")

	// Membuat filter untuk mencocokkan toko berdasarkan slug
	filter := bson.M{"slug": slug}

	collection := config.Mongoconn.Collection("kategori")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var kategori model.Kategori
	err := collection.FindOne(ctx, filter).Decode(&kategori)
	if err != nil {
		http.Error(w, "Toko not found", http.StatusNotFound)
		return
	}

	// Kirim respons berhasil
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(kategori)
}

// GetAllProduct untuk mendapatkan semua produk
func GetAllProduk(w http.ResponseWriter, r *http.Request) {
	collection := config.Mongoconn.Collection("kategori")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		http.Error(w, "Failed to fetch produk", http.StatusInternalServerError)
		return
	}
	defer cursor.Close(ctx)

	var kategori []model.Kategori
	for cursor.Next(ctx) {
		var k model.Kategori
		cursor.Decode(&k)
		kategori = append(kategori, k)
	}

	if err := cursor.Err(); err != nil {
		http.Error(w, "Failed to fetch produk", http.StatusInternalServerError)
		return
	}

	// Kirim respons berhasil
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(kategori)
}

// GetProductByCategory untuk mendapatkan detail produk berdasarkan kategori
func GetProdukByKategori(w http.ResponseWriter, r *http.Request) {
	// Mengambil slug dari query parameter
	slug := r.URL.Query().Get("slug")

	// Membuat filter untuk mencocokkan toko berdasarkan slug
	filter := bson.M{"slug": slug}

	collection := config.Mongoconn.Collection("kategori")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var kategori model.Kategori
	err := collection.FindOne(ctx, filter).Decode(&kategori)
	if err != nil {
		http.Error(w, "Toko not found", http.StatusNotFound)
		return
	}

	// Kirim respons berhasil
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(kategori.Produk)
}

// GetProductByCategoryAndProduct untuk mendapatkan detail produk berdasarkan kategori dan produk
func GetProdukByKategoriAndProduk(w http.ResponseWriter, r *http.Request) {
	// Mengambil slug dari query parameter
	slug := r.URL.Query().Get("slug")
	produkSlug := r.URL.Query().Get("produkSlug")

	// Membuat filter untuk mencocokkan toko berdasarkan slug
	filter := bson.M{"slug": slug, "menu.slug": produkSlug}

	collection := config.Mongoconn.Collection("kategori")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var kategori model.Kategori
	err := collection.FindOne(ctx, filter).Decode(&kategori)
	if err != nil {
		http.Error(w, "Produk not found", http.StatusNotFound)
		return
	}

	// Kirim respons berhasil
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(kategori.Produk)
}

// GetProductByCategoryAndProductID untuk mendapatkan detail produk berdasarkan kategori dan produk ID
func GetProdukByKategoriAndProdukID(w http.ResponseWriter, r *http.Request) {
	// Mengambil slug dari query parameter
	slug := r.URL.Query().Get("slug")
	produkID := r.URL.Query().Get("produkID")

	// Membuat filter untuk mencocokkan toko berdasarkan slug
	filter := bson.M{"slug": slug, "menu._id": produkID}

	collection := config.Mongoconn.Collection("kategori")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var kategori model.Kategori
	err := collection.FindOne(ctx, filter).Decode(&kategori)
	if err != nil {
		http.Error(w, "Produk not found", http.StatusNotFound)
		return
	}

	// Kirim respons berhasil
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(kategori.Produk)
}