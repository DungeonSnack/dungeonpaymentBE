package controller

import (
	"context"
	"dungeonSnackBE/model/pengguna"
	"encoding/json"
	"net/http"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
)

// Fungsi Registrasi

func Register(w http.ResponseWriter, r *http.Request, db *mongo.Client) {
    // Menghubungkan ke database dsdatabase dan koleksi user
    collection := db.Database("dsdatabase").Collection("user")
    var newUser model.Users

    // Decode JSON input ke struct Users
    if err := json.NewDecoder(r.Body).Decode(&newUser); err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    // Cek apakah email sudah ada
    var existingUser model.Users
    err := collection.FindOne(context.TODO(), bson.M{"email": newUser.Email}).Decode(&existingUser)
    if err == nil {
        http.Error(w, "User already exists", http.StatusConflict)
        return
    }

    // Hash password
    hashedPassword, err := bcrypt.GenerateFromPassword([]byte(newUser.Password), bcrypt.DefaultCost)
    if err != nil {
        http.Error(w, "Failed to hash password", http.StatusInternalServerError)
        return
    }

    // Tetapkan role pembeli secara default
    newUser.Role = "pembeli"  // Mengatur role pembeli

    // Buat user baru dengan hashed password dan waktu pembuatan
    newUser.Password = string(hashedPassword)
    newUser.CreatedAt = time.Now()
    newUser.UpdatedAt = time.Now()
    newUser.ID = primitive.NewObjectID()  // Membuat ID baru

    // Simpan user baru ke MongoDB
    _, err = collection.InsertOne(context.TODO(), newUser)
    if err != nil {
        http.Error(w, "Failed to register user", http.StatusInternalServerError)
        return
    }

    w.WriteHeader(http.StatusCreated)
    json.NewEncoder(w).Encode("User registered successfully")
}