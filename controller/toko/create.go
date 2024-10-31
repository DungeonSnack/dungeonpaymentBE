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
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func CreateToko(w http.ResponseWriter, r *http.Request) {
	// Mendapatkan user dari context (asumsi user diambil dari middleware)
	user, ok := r.Context().Value("user").(model.Users)
	if !ok || user.Role != "penjual" {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	// Cek apakah user sudah memiliki toko
	collection := config.Mongoconn.Collection("toko")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var existingToko model.Toko
	err := collection.FindOne(ctx, bson.M{"userId": user.ID}).Decode(&existingToko)
	if err == nil {
		http.Error(w, "User sudah memiliki toko", http.StatusBadRequest)
		return
	}

	// Decode payload toko baru
	var toko model.Toko
	err = json.NewDecoder(r.Body).Decode(&toko)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	// Set data toko
	toko.Slug = slug.GenerateSlug(toko.NamaToko)
	toko.ID = primitive.NewObjectID()
	toko.UserID = user.ID // Menyimpan ID user
	toko.CreatedAt = time.Now()
	toko.UpdatedAt = time.Now()

	// Insert toko baru
	_, err = collection.InsertOne(ctx, toko)
	if err != nil {
		http.Error(w, "Failed to insert toko", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(toko)
}
