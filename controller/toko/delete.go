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

func DeleteToko(w http.ResponseWriter, r *http.Request) {
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

	_, err = collection.DeleteOne(ctx, primitive.D{primitive.E{Key: "slug", Value: toko.Slug}})
	if err != nil {
		http.Error(w, "Failed to delete toko", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(toko)
}