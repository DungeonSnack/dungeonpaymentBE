package order

import (
	"context"
	"dungeonSnackBE/config"
	"dungeonSnackBE/model"
	"encoding/json"
	"net/http"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// UpdateOrder untuk mengupdate order
func UpdateOrder(w http.ResponseWriter, r *http.Request) {
	// Dekode body permintaan untuk mendapatkan detail order baru
	var order model.Order
	err := json.NewDecoder(r.Body).Decode(&order)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	// Set waktu pembuatan
	order.UpdatedAt = time.Now() // Set waktu saat ini untuk updatedAt

	// Membuat filter untuk mencocokkan user berdasarkan user_id
	filter := bson.M{"_id": order.UserId}

	update := bson.M{
		"$set": bson.M{
			"Order.$[o]": order,
		},
	}

	arrayFilter := bson.A{
		bson.M{"o._id": order.ID},
	}

	collection := config.Mongoconn.Collection("user")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	result, err := collection.UpdateOne(ctx, filter, update, &options.UpdateOptions{
		ArrayFilters: &options.ArrayFilters{
			Filters: arrayFilter,
		},
	})
	if err != nil {
		http.Error(w, "Failed to update order", http.StatusInternalServerError)
		return
	}

	if result.MatchedCount == 0 {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	// Kirim respons berhasil
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(order)
}