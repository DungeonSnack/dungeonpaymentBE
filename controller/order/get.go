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
	"go.mongodb.org/mongo-driver/mongo/options"
)

func Getorder(w http.ResponseWriter, r *http.Request) {
	// Koneksi ke koleksi MongoDB toko
	collection := config.Mongoconn.Collection("order")

	// Membuat context dengan timeout
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Membuat opsi untuk mengabaikan field 'menu'
	projection := options.Find().SetProjection(bson.M{"menu": 0})

	// Menemukan semua toko, tanpa array 'menu'
	cursor, err := collection.Find(ctx, bson.M{}, projection)
	if err != nil {
		http.Error(w, "Failed to retrieve orders", http.StatusInternalServerError)
		return
	}
	defer cursor.Close(ctx)

	// Dekode hasil query ke dalam slice dari model Toko
	var orders []model.Order
	if err = cursor.All(ctx, &orders); err != nil {
		http.Error(w, "Failed to parse orders", http.StatusInternalServerError)
		return
	}

	// Kirim daftar toko (order) sebagai respons
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(orders)
}

func GetorderByID(w http.ResponseWriter, r *http.Request) {
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

	// Koneksi ke koleksi MongoDB order
	collection := config.Mongoconn.Collection("order")

	// Membuat context dengan timeout
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var order model.Order
	err = collection.FindOne(ctx, bson.M{"_id": objID}).Decode(&order)
	if err != nil {
		http.Error(w, "order not found", http.StatusNotFound)
		return
	}

	// Kirim data toko sebagai respons
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(order)
}