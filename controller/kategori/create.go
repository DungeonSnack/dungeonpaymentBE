package kategori

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

func AddKategori(w http.ResponseWriter, r *http.Request) {
	var kategori model.Kategori
	err := json.NewDecoder(r.Body).Decode(&kategori)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	kategori.ID = primitive.NewObjectID()
	kategori.Slug = slug.GenerateSlug(kategori.NamaKategori)

	for i := range kategori.Produk {
		kategori.Produk[i].ID = primitive.NewObjectID() // Generate new ObjectID
		kategori.Produk[i].CreatedAt = time.Now()       // Set createdAt to current time
		kategori.Produk[i].UpdatedAt = time.Now()       // Set updatedAt to current time
	}

	collection := config.Mongoconn.Collection("kategori")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	_, err = collection.InsertOne(ctx, kategori)
	if err != nil {
		http.Error(w, "Failed to create kategori", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(kategori)
}