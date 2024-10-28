package order

import (
	"context"
	"encoding/json"
	"net/http"
	"time"
	"dungeonSnackBE/config"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Fungsi untuk menghapus order berdasarkan ID
func DeleteorderByID(w http.ResponseWriter, r *http.Request) {
	// Mendapatkan ID toko dari parameter URL
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

	// Koneksi ke koleksi MongoDB order
	collection := config.Mongoconn.Collection("order")

	// Membuat context dengan timeout
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Melakukan penghapusan berdasarkan ID
	result, err := collection.DeleteOne(ctx, bson.M{"_id": objID})
	if err != nil {
		http.Error(w, "Failed to delete order", http.StatusInternalServerError)
		return
	}

	// Memeriksa apakah ada data yang dihapus
	if result.DeletedCount == 0 {
		http.Error(w, "order not found", http.StatusNotFound)
		return
	}

	// Mengirimkan respons sukses
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{
		"message": "order deleted successfully",
	})
}