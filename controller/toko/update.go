package toko

import (
	"context"
	"dungeonSnackBE/config"
	"dungeonSnackBE/helper/slug"
	"dungeonSnackBE/model"
	"encoding/json"
	"net/http"
	"time"

	"go.mongodb.org/mongo-driver/bson"
)

// UpdateToko adalah fungsi untuk mengupdate data toko
func UpdateToko(w http.ResponseWriter, r *http.Request) {
	// Mendapatkan user dari context (misalnya dari middleware)
	user, ok := r.Context().Value("user").(model.Users)
	if !ok || user.Role != "penjual" {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	// Decode payload toko yang akan diupdate
	var toko model.Toko
	err := json.NewDecoder(r.Body).Decode(&toko)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	// Set slug dan waktu update
	toko.Slug = slug.GenerateSlug(toko.NamaToko)
	toko.UpdatedAt = time.Now()

	// Akses koleksi toko di MongoDB
	collection := config.Mongoconn.Collection("toko")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Lakukan update pada dokumen toko berdasarkan toko_id
	_, err = collection.UpdateOne(ctx, bson.M{"_id": toko.ID}, bson.M{"$set": toko})
	if err != nil {
		http.Error(w, "Failed to update toko", http.StatusInternalServerError)
		return
	}

	// Kembalikan respon sukses
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(toko)
}
