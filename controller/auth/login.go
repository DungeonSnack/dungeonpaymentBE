package controller

import (
	"context"
	"dungeonSnackBE/model/pengguna"
	"encoding/json"
	"net/http"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
)

func Login(w http.ResponseWriter, r *http.Request, db *mongo.Client) {
    collection := db.Database("user").Collection("users")
    var userCredentials model.Users
    var foundUser model.Users

    if err := json.NewDecoder(r.Body).Decode(&userCredentials); err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    // Cari user berdasarkan email
    err := collection.FindOne(context.TODO(), bson.M{"email": userCredentials.Email}).Decode(&foundUser)
    if err != nil {
        http.Error(w, "Invalid email or password", http.StatusUnauthorized)
        return
    }

    // Cek password
    if err := bcrypt.CompareHashAndPassword([]byte(foundUser.Password), []byte(userCredentials.Password)); err != nil {
        http.Error(w, "Invalid email or password", http.StatusUnauthorized)
        return
    }

    // Kembalikan respon login berhasil
    w.WriteHeader(http.StatusOK)
    json.NewEncoder(w).Encode("Login successful")
}