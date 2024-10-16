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


// CreateProduct untuk menambahkan produk baru
func AddProduk(w http.ResponseWriter, r *http.Request) {
	// Mengambil slug dari query parameter
	slug := r.URL.Query().Get("slug")

	// Dekode body permintaan untuk mendapatkan detail menu baru
	var newMenu model.Produk
	err := json.NewDecoder(r.Body).Decode(&newMenu)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	// Set id menu baru dan waktu pembuatan
	newMenu.ID = primitive.NewObjectID() // Set ObjectID baru untuk menu
	newMenu.CreatedAt = time.Now()       // Set waktu saat ini untuk createdAt
	newMenu.UpdatedAt = time.Now()       // Set waktu saat ini untuk updatedAt

	// Membuat filter untuk mencocokkan toko berdasarkan slug
	filter := bson.M{"slug": slug}

	update := bson.M{
		"$push": bson.M{
			"menu": newMenu,
		},
	}

	collection := config.Mongoconn.Collection("kategori")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	result, err := collection.UpdateOne(ctx, filter, update)
	if err != nil {
		http.Error(w, "Failed to add produk", http.StatusInternalServerError)
		return
	}

	if result.MatchedCount == 0 {
		http.Error(w, "Toko not found", http.StatusNotFound)
		return
	}

	// Kirim respons berhasil
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"message": "Produk added successfully"}`))
}