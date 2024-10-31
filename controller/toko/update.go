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

// UpdateToko is a function to update a toko
func UpdateToko(w http.ResponseWriter, r *http.Request) {
	var toko model.Toko
	err := json.NewDecoder(r.Body).Decode(&toko)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	toko.Slug = slug.GenerateSlug(toko.NamaToko)
	toko.UpdatedAt = time.Now() // Set updatedAt to current time

	collection := config.Mongoconn.Collection("toko")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	_, err = collection.UpdateOne(ctx, bson.M{"toko_id": toko.ID}, bson.M{"$set": toko})
	if err != nil {
		http.Error(w, "Failed to update toko", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(toko)
}