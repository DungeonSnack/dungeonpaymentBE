package order

import (
	"dungeonSnackBE/model"
	"encoding/json"
	"net/http"
	"time"
	"dungeonSnackBE/config"
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func UpdateorderByID(w http.ResponseWriter, r *http.Request) {
	// Mendapatkan ID toko dari parameter URL
	id := r.URL.Query().Get("id")

	// Memastikan ID tidak kosong
	if id == "" {
		http.Error(w, "ID is required", http.StatusBadRequest)
		return
	}

	// Mengubah ID dari string ke ObjectID
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		http.Error(w, "Invalid ID format", http.StatusBadRequest)
		return
	}

	// Mengurai JSON dari body request
	var updatedorder model.Order
	err = json.NewDecoder(r.Body).Decode(&updatedorder)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	// Mengatur field yang akan diupdate
	updateFields := bson.M{
		"Quantity": updatedorder.Quantity,
		"updatedAt": time.Now(),
	}

	// Koneksi ke koleksi MongoDB order
	collection := config.Mongoconn.Collection("order")

	// Membuat context dengan timeout
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Melakukan update berdasarkan ID
	result, err := collection.UpdateOne(
		ctx,
		bson.M{"order_id": objID},
		bson.M{"$set": updateFields},
	)
	if err != nil {
		http.Error(w, "Failed to update order", http.StatusInternalServerError)
		return
	}

	// Memeriksa apakah data berhasil diupdate
	if result.MatchedCount == 0 {
		http.Error(w, "order not found", http.StatusNotFound)
		return
	}

	// Mengirimkan respons sukses
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{
		"message": "order updated successfully",
	})
}