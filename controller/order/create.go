package order

import (
	"context"
	"dungeonSnackBE/config"
	"dungeonSnackBE/model"
	"encoding/json"
	"net/http"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func Addorder(w http.ResponseWriter, r *http.Request) {
	var order model.Order
	err := json.NewDecoder(r.Body).Decode(&order)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	for i := range order.Payment {
		order.Payment[i].ID = primitive.NewObjectID() // Generate new ObjectID
		order.Payment[i].CreatedAt = time.Now()       // Set createdAt to current time
		order.Payment[i].UpdatedAt = time.Now()       // Set updatedAt to current time
	}

	collection := config.Mongoconn.Collection("order")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	_, err = collection.InsertOne(ctx, order)
	if err != nil {
		http.Error(w, "Failed to create order", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(order)
}
