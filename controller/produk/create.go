package produk

import (
	"context"
	"dungeonSnackBE/config"
	"dungeonSnackBE/helper/slug"
	"dungeonSnackBE/model"
	"encoding/json"
	"net/http"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// CreateProduk untuk menambahkan produk baru
func CreateProduk(w http.ResponseWriter, r *http.Request) {
	// Mendapatkan user dari context (misal dari middleware)
	user, ok := r.Context().Value("user").(model.Users)
	if !ok || user.Role != "penjual" {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	// Mengecek apakah user memiliki toko
	var toko model.Toko
	collectionToko := config.Mongoconn.Collection("toko")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err := collectionToko.FindOne(ctx, bson.M{"userId": user.ID}).Decode(&toko)
	if err != nil {
		http.Error(w, "User tidak memiliki toko", http.StatusBadRequest)
		return
	}

	// Decode payload produk baru
	var produk model.Menu
	err = json.NewDecoder(r.Body).Decode(&produk)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	// Set data produk baru
	produk.ID = primitive.NewObjectID()
	produk.TokoID = toko.ID // Menyimpan ID toko yang terkait
	produk.Slug = slug.GenerateSlug(produk.NamaProduk)
	produk.CreatedAt = time.Now()
	produk.UpdatedAt = time.Now()

	// Menyimpan produk ke koleksi "produk"
	collectionProduk := config.Mongoconn.Collection("produk")
	_, err = collectionProduk.InsertOne(ctx, produk)
	if err != nil {
		http.Error(w, "Failed to create produk", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(produk)
}

