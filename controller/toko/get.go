package toko

import (
	"context"
	"dungeonSnackBE/config"
	"dungeonSnackBE/model"
	"encoding/json"
	"net/http"
	"time"

	"go.mongodb.org/mongo-driver/bson"
)

func getToko(w http.ResponseWriter, r *http.Request) {
	// Mengambil slug dari query parameter
	slug := r.URL.Query().Get("slug")

	// Membuat filter untuk mencocokkan toko berdasarkan slug
	filter := bson.M{"slug": slug}

	collection := config.Mongoconn.Collection("toko")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var toko model.Toko
	err := collection.FindOne(ctx, filter).Decode(&toko)
	if err != nil {
		http.Error(w, "Toko not found", http.StatusNotFound)
		return
	}

	// Kirim respons berhasil
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(toko)
}