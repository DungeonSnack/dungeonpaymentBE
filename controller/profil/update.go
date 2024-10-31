package profil

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

func UpdateProfil(w http.ResponseWriter, r *http.Request) {
	vars := r.URL.Query()
	userID := vars.Get("id")

	// Konversi user_id dari string ke ObjectID MongoDB
	oid, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		http.Error(w, "ID pengguna tidak valid", http.StatusBadRequest)
		return
	}

	// Parsing body untuk mendapatkan data yang diperbarui
	var updatedUser model.Users
	err = json.NewDecoder(r.Body).Decode(&updatedUser)
	if err != nil {
		http.Error(w, "Payload permintaan tidak valid", http.StatusBadRequest)
		return
	}

	// Atur timestamp yang diperbarui
	updatedUser.UpdatedAt = time.Now()

	// Siapkan field yang akan diperbarui
	updateFields := bson.M{
		"nama":       updatedUser.Nama,
		"no_hp":      updatedUser.No_HP,
		"email":      updatedUser.Email,
		"updated_at": updatedUser.UpdatedAt,
	}

	// Cek dan perbarui role jika diperlukan
	if updatedUser.Role == "penjual" {
		updateFields["role"] = "penjual"
	}

	// Mengatur koleksi MongoDB
	collection := config.Mongoconn.Collection("user")

	// Buat context dengan batas waktu
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Definisikan query pembaruan menggunakan $set untuk memperbarui field tertentu saja
	update := bson.M{"$set": updateFields}

	// Lakukan operasi pembaruan
	result, err := collection.UpdateOne(ctx, bson.M{"_id": oid}, update)
	if err != nil {
		http.Error(w, "Gagal memperbarui pengguna", http.StatusInternalServerError)
		return
	}

	// Cek apakah ada dokumen yang berhasil diperbarui
	if result.MatchedCount == 0 {
		http.Error(w, "Pengguna tidak ditemukan", http.StatusNotFound)
		return
	}

	// Kembalikan respon sukses
	response := map[string]interface{}{
		"message": "Pengguna berhasil diperbarui",
		"user_id": userID,
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
