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

// UpdateProduct untuk mengubah produk yang sudah ada
func UpdateProduk(w http.ResponseWriter, r *http.Request) {
	// Mengambil slug dari query parameter
	slug := r.URL.Query().Get("slug")

	// Dekode body permintaan untuk mendapatkan detail menu baru
	var updatedMenu model.Produk
	err := json.NewDecoder(r.Body).Decode(&updatedMenu)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	// Set waktu pembuatan
	updatedMenu.UpdatedAt = time.Now()

	// Membuat filter untuk mencocokkan toko berdasarkan slug
	filter := bson.M{"slug": slug, "menu._id": updatedMenu.ID}

	update := bson.M{
		"$set": bson.M{
			"menu.$": updatedMenu,
		},
	}

	collection := config.Mongoconn.Collection("kategori")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	result, err := collection.UpdateOne(ctx, filter, update)
	if err != nil {
		http.Error(w, "Failed to update produk", http.StatusInternalServerError)
		return
	}

	if result.MatchedCount == 0 {
		http.Error(w, "Produk not found", http.StatusNotFound)
		return
	}

	// Kirim respons berhasil
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"message": "Produk updated successfully"}`))
}