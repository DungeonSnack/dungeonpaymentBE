package pembayaran

import (
	"context"
	"dungeonSnackBE/config"
	"dungeonSnackBE/model"
	"encoding/json"
	"net/http"
	"time"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/bson"
)

// GetProduct untuk mendapatkan detail pembayaran berdasarkan slug
func GetpembayaranByorder(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	slug := params["slug"]

	collection := config.Mongoconn.Collection("order")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var order model.Order
	err := collection.FindOne(ctx, bson.M{"slug": slug}).Decode(&order)
	if err != nil {
			http.Error(w, "order not found", http.StatusNotFound)
			return
	}

	json.NewEncoder(w).Encode(order.Pembayaran)
}

func GetpembayaranByID(w http.ResponseWriter, r *http.Request) {
	pembayaranID := r.URL.Query().Get("pembayaran_id")

	// Konversi pembayaranID dari string ke primitive.ObjectID
	objectID, err := primitive.ObjectIDFromHex(pembayaranID)
	if err != nil {
			http.Error(w, "Invalid pembayaran ID format", http.StatusBadRequest)
			return
	}

	// Cari pembayaran berdasarkan pembayaran_id dari database
	var pembayaran model.Pembayaran
	collection := config.Mongoconn.Collection("order") // Pastikan koleksi yang benar
	err = collection.FindOne(context.TODO(), bson.M{"pembayaran._id": objectID}).Decode(&pembayaran)

	if err != nil {
			http.Error(w, "pembayaran not found", http.StatusNotFound)
			return
	}

	json.NewEncoder(w).Encode(pembayaran)
}