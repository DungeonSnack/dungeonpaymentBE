package produk

import (
	"context"
	"dungeonSnackBE/config"
	"dungeonSnackBE/model"
	"encoding/json"
	"net/http"
	"time"

	"go.mongodb.org/mongo-driver/bson"
)

// GetOrder untuk mendapatkan order berdasarkan user_id
func GetOrder(w http.ResponseWriter, r *http.Request) {
	// Dekode body permintaan untuk mendapatkan detail order baru
	var order model.Order
	err := json.NewDecoder(r.Body).Decode(&order)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	// Membuat filter untuk mencocokkan user berdasarkan user_id
	filter := bson.M{"_id": order.UserId}

	collection := config.Mongoconn.Collection("user")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err = collection.FindOne(ctx, filter).Decode(&order)
	if err != nil {
		http.Error(w, "Failed to get order", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(order)
}

// GetAllOrder untuk mendapatkan semua order
func GetAllOrder(w http.ResponseWriter, r *http.Request) {
	// Membuat filter untuk mencocokkan semua order
	filter := bson.M{}

	collection := config.Mongoconn.Collection("user")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	cursor, err := collection.Find(ctx, filter)
	if err != nil {
		http.Error(w, "Failed to get order", http.StatusInternalServerError)
		return
	}

	var orders []model.Order
	err = cursor.All(ctx, &orders)
	if err != nil {
		http.Error(w, "Failed to get order", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(orders)
}

//GetOrderById untuk mendapatkan order berdasarkan order_id
func GetOrderById(w http.ResponseWriter, r *http.Request) {
	// Dekode body permintaan untuk mendapatkan detail order baru
	var order model.Order
	err := json.NewDecoder(r.Body).Decode(&order)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	// Membuat filter untuk mencocokkan order berdasarkan order_id
	filter := bson.M{"order_id": order.ID}

	collection := config.Mongoconn.Collection("user")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err = collection.FindOne(ctx, filter).Decode(&order)
	if err != nil {
		http.Error(w, "Failed to get order", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(order)
}