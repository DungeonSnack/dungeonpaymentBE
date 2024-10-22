package produk

import (
	"context"
	"dungeonSnackBE/config"
	"dungeonSnackBE/model"
	"encoding/json"
	"net/http"
	"time"

	"go.mongodb.org/mongo-driver/bson"
)

// UpdateProduct untuk mengubah produk yang sudah ada
func UpdateProduk(w http.ResponseWriter, r *http.Request) {
	// Mengambil slug dari query parameter
	slug := r.URL.Query().Get("slug")

	// Dekode body permintaan untuk mendapatkan detail produk baru
	var updatedProduk model.Produk
	err := json.NewDecoder(r.Body).Decode(&updatedProduk)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	// Set waktu pembuatan
	updatedProduk.UpdatedAt = time.Now()

	// Membuat filter untuk mencocokkan kategori berdasarkan slug
	filter := bson.M{"slug": slug, "Produk._id": updatedProduk.ID}

	update := bson.M{
		"$set": bson.M{
			"Produk.$": updatedProduk,
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
