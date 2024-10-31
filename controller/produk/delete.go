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

// DeleteProduk untuk menghapus produk
func DeleteProduk(w http.ResponseWriter, r *http.Request) {
	// Mendapatkan user dari context (misalnya dari middleware)
	user, ok := r.Context().Value("user").(model.Users)
	if !ok || user.Role != "penjual" {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	// Mendapatkan ID produk dari parameter URL
	id := r.URL.Query().Get("id")
	if id == "" {
		http.Error(w, "ID is required", http.StatusBadRequest)
		return
	}

	// Konversi ID dari string ke ObjectID
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

	// Melakukan penghapusan berdasarkan ID dan memastikan ID user cocok
	result, err := collection.DeleteOne(ctx, bson.M{
		"_id": objID,
		"user_id": user.ID, // Hanya produk milik user penjual yang bisa dihapus
	})
	if err != nil {
		http.Error(w, "Failed to delete produk", http.StatusInternalServerError)
		return
	}

	// Memeriksa apakah ada data yang dihapus
	if result.DeletedCount == 0 {
		http.Error(w, "Produk not found", http.StatusNotFound)
		return
	}

	// Mengirimkan respons sukses
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{
		"message": "Produk deleted successfully",
	})
}
