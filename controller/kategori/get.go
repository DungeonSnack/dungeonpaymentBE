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

func GetKategori (w http.ResponseWriter, r *http.Request) {
	var kategori []model.Kategori

	collection := config.Mongoconn.Collection("kategori")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		http.Error(w, "Failed to fetch kategori", http.StatusInternalServerError)
		return
	}
	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		var k model.Kategori
		cursor.Decode(&k)
		kategori = append(kategori, k)
	}

	if err := cursor.Err(); err != nil {
		http.Error(w, "Failed to fetch kategori", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(kategori)
}