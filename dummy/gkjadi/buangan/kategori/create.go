package kategori

import (
	"context"
	"dungeonSnackBE/config"
	"dungeonSnackBE/helper/slug"
	"dungeonSnackBE/model"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Add Kategori to database
func AddKategori(w http.ResponseWriter, r *http.Request) {
	var kategori model.Kategori
	err := json.NewDecoder(r.Body).Decode(&kategori)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	kategori.ID = primitive.NewObjectID()
	kategori.Slug = slug.GenerateSlug(kategori.NamaKategori)

	for i := range kategori.Produk {
		fmt.Println("Setting time for kategori produk:", kategori.Produk[i].NamaProduk)
		kategori.Produk[i].ID = primitive.NewObjectID()
		kategori.Produk[i].CreatedAt = time.Now()
		kategori.Produk[i].UpdatedAt = time.Now()
		fmt.Println("CreatedAt:", kategori.Produk[i].CreatedAt)
	}

	collection := config.Mongoconn.Collection("kategori")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	_, err = collection.InsertOne(ctx, kategori)
	if err != nil {
		http.Error(w, "Failed to create kategori", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(kategori)
}
