package toko

import (
	"dungeonSnackBE/model"
	"encoding/json"
	"net/http"
	"time"
	"dungeonSnackBE/helper/slug"
	"dungeonSnackBE/config"
	"context"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GetToko(w http.ResponseWriter, r *http.Request) {
	var toko model.Toko
	err := json.NewDecoder(r.Body).Decode(&toko)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	toko.Slug = slug.GenerateSlug(toko.NamaToko)

	collection := config.Mongoconn.Collection("toko")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err = collection.FindOne(ctx, primitive.D{primitive.E{Key: "slug", Value: toko.Slug}}).Decode(&toko)
	if err != nil {
		http.Error(w, "Failed to get toko", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(toko)
}