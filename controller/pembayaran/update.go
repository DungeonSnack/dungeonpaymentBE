package pembayaran

import (
	"context"
	"dungeonSnackBE/config"
	"dungeonSnackBE/model"
	"encoding/json"
	"net/http"
	"time"

	"go.mongodb.org/mongo-driver/bson"
)

// UpdateProduct untuk mengubah pembayaran yang sudah ada
func Updatepembayaran(w http.ResponseWriter, r *http.Request) {
	// Mengambil slug dari query parameter
	slug := r.URL.Query().Get("slug")

	// Dekode body permintaan untuk mendapatkan detail pembayaran baru
	var updatedpembayaran model.Payment
	err := json.NewDecoder(r.Body).Decode(&updatedpembayaran)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	// Set waktu pembuatan
	updatedpembayaran.UpdatedAt = time.Now()

	// Membuat filter untuk mencocokkan pembayaran berdasarkan slug
	filter := bson.M{"slug": slug, "pembayaran._id": updatedpembayaran.ID}

	update := bson.M{
		"$set": bson.M{
			"pembayaran.$": updatedpembayaran,
		},
	}

	collection := config.Mongoconn.Collection("pembayaran")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	result, err := collection.UpdateOne(ctx, filter, update)
	if err != nil {
		http.Error(w, "Failed to update pembayaran", http.StatusInternalServerError)
		return
	}

	if result.MatchedCount == 0 {
		http.Error(w, "pembayaran not found", http.StatusNotFound)
		return
	}

	// Kirim respons berhasil
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"message": "pembayaran updated successfully"}`))
}
