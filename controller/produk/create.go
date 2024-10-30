package produk

import (
	"context"
	"dungeonSnackBE/config"
	"dungeonSnackBE/helper/slug"
	"dungeonSnackBE/model"
	"encoding/json"
	"net/http"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// CreateProduk untuk menambahkan produk baru
func CreateProduk(w http.ResponseWriter, r *http.Request) {
	var produk model.Menu
	err := json.NewDecoder(r.Body).Decode(&produk)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	produk.ID = primitive.NewObjectID()
	produk.Slug = slug.GenerateSlug(produk.NamaProduk)
	produk.CreatedAt = time.Now()
	produk.UpdatedAt = time.Now()

	collection := config.Mongoconn.Collection("produk")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	_, err = collection.InsertOne(ctx, produk)
	if err != nil {
		http.Error(w, "Failed to create produk", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(produk)
}
