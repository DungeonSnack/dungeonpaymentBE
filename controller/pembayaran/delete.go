package pembayaran

import (
	"context"
	"dungeonSnackBE/config"
	"github.com/gorilla/mux"
	"net/http"
	"time"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// DeleteProduct untuk menghapus pembayaran yang sudah ada
func Deletepembayaran(w http.ResponseWriter, r *http.Request) {
	// Mengambil slug toko dan pembayaran_id dari query parameter
	params := mux.Vars(r)
	MarketSlug := params["slug"]
	pembayaranIDHex := r.URL.Query().Get("pembayaran_id")

	// Konversi pembayaran_id dari hex string ke ObjectID
	pembayaranID, err := primitive.ObjectIDFromHex(pembayaranIDHex)
	if err != nil {
		http.Error(w, "Invalid pembayaran ID", http.StatusBadRequest)
		return
	}

	// Membuat filter untuk mencari toko berdasarkan slug dan pembayaran_id
	filter := bson.M{
		"slug":      MarketSlug,
		"pembayaran._id":  pembayaranID,
	}

	// Membuat update untuk menghapus pembayaran dari array pembayaran
	update := bson.M{
		"$pull": bson.M{
			"pembayaran": bson.M{
				"_id": pembayaranID,
			},
		},
	}

	collection := config.Mongoconn.Collection("market")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	result, err := collection.UpdateOne(ctx, filter, update)
	if err != nil {
		http.Error(w, "Failed to delete pembayaran", http.StatusInternalServerError)
		return
	}

	// Cek apakah pembayaran berhasil dihapus
	if result.MatchedCount == 0 {
		http.Error(w, "Toko or pembayaran not found", http.StatusNotFound)
		return
	}

	// Mengembalikan respons sukses
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"message": "pembayaran deleted successfully"}`))
}