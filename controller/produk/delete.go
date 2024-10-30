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

// DeleteProduk untuk menghapus produk
func DeleteProduk(w http.ResponseWriter, r *http.Request) {
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

	// Melakukan penghapusan berdasarkan ID
	result, err := collection.DeleteOne(ctx, model.Menu{ID: objID})
	if err != nil {
		http.Error(w, "Failed to delete produk", http.StatusInternalServerError)
		return
	}

	// Memeriksa apakah ada data yang dihapus
	if result.DeletedCount == 0 {
		http.Error(w, "produk not found", http.StatusNotFound)
		return
	}

	// Mengirimkan respons sukses
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{
		"message": "produk deleted successfully",
	})
}