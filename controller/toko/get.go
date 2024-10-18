package toko

import (
	"dungeonSnackBE/model"
	"encoding/json"
	"net/http"
	"time"
	"dungeonSnackBE/config"
	"context"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GetToko(w http.ResponseWriter, r *http.Request) {
	var tokos []model.Toko

	collection := config.Mongoconn.Collection("toko")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	cursor, err := collection.Find(ctx, primitive.D{})
	if err != nil {
		http.Error(w, "Failed to get toko", http.StatusInternalServerError)
		return
	}
	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		var toko model.Toko
		cursor.Decode(&toko)
		tokos = append(tokos, toko)
	}

	json.NewEncoder(w).Encode(tokos)
}

