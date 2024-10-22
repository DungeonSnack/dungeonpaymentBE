package order

import (
	"context"
	"dungeonSnackBE/config"
	"dungeonSnackBE/model"
	"encoding/json"
	"net/http"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// CreateOrder untuk menambahkan order baru
func CreateOrder(w http.ResponseWriter, r *http.Request) {
	// Dekode body permintaan untuk mendapatkan detail order baru
	var newOrder model.Order
	err := json.NewDecoder(r.Body).Decode(&newOrder)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	// Set id order baru dan waktu pembuatan
	newOrder.ID = primitive.NewObjectID() // Set ObjectID baru untuk order
	newOrder.CreatedAt = time.Now()       // Set waktu saat ini untuk createdAt
	newOrder.UpdatedAt = time.Now()       // Set waktu saat ini untuk updatedAt

	// Membuat filter untuk mencocokkan user berdasarkan user_id
	filter := bson.M{"_id": newOrder.UserId}

	update := bson.M{
		"$push": bson.M{
			"Order": newOrder,
		},
	}

	collection := config.Mongoconn.Collection("user")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	result, err := collection.UpdateOne(ctx, filter, update)
	if err != nil {
		http.Error(w, "Failed to add order", http.StatusInternalServerError)
		return
	}

	if result.MatchedCount == 0 {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	// Kirim respons berhasil
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(newOrder)
}