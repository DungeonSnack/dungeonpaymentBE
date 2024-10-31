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

// UpdateProduk untuk mengupdate produk
func UpdateProduk(w http.ResponseWriter, r *http.Request) {
	// Mendapatkan user dari context (misal dari middleware)
	user, ok := r.Context().Value("user").(model.Users)
	if !ok || user.Role != "penjual" {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	// Ambil ID produk dari query parameter
	id := r.URL.Query().Get("id")
	if id == "" {
		http.Error(w, "ID is required", http.StatusBadRequest)
		return
	}

	// Konversi ID ke ObjectID MongoDB
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		http.Error(w, "Invalid ID format", http.StatusBadRequest)
		return
	}

	// Decode payload untuk produk yang diperbarui
	var updatedProduk model.Menu
	err = json.NewDecoder(r.Body).Decode(&updatedProduk)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	// Set data yang akan diperbarui
	updateFields := bson.M{
		"nama_produk": updatedProduk.NamaProduk,
		"harga":       updatedProduk.Price,
		"stok":        updatedProduk.Stok,
		"updatedAt":   time.Now(),
	}

	// Akses koleksi produk di MongoDB
	collection := config.Mongoconn.Collection("produk")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Lakukan operasi update pada produk berdasarkan ID
	result, err := collection.UpdateOne(
		ctx,
		bson.M{"_id": objID},
		bson.M{"$set": updateFields},
	)
	if err != nil {
		http.Error(w, "Failed to update produk", http.StatusInternalServerError)
		return
	}

	// Cek apakah produk ditemukan dan diperbarui
	if result.MatchedCount == 0 {
		http.Error(w, "Produk not found", http.StatusNotFound)
		return
	}

	// Kembalikan respon sukses
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{
		"message": "Produk updated successfully",
	})
}
