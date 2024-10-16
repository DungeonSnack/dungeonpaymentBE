package kategori

import (
	"dungeonSnackBE/config"
	"dungeonSnackBE/helper/slug"
	"encoding/json"
	"net/http"
	"time"

	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func DeleteKategori(w http.ResponseWriter, r *http.Request) {
	var kategori model.Kategori
	err := json.NewDecoder(r.Body).Decode(&kategori)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	collection := config.Mongoconn.Collection("kategori")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	_, err = collection.DeleteOne(ctx, bson.M{"_id": kategori.ID})
	if err != nil {
		http.Error(w, "Failed to delete kategori", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(kategori)
}