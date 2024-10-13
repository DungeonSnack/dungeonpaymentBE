package controller

import (
    "context"
    "encoding/json"
    "net/http"
    "time"
    "go.mongodb.org/mongo-driver/bson"
    "go.mongodb.org/mongo-driver/mongo"
    "golang.org/x/crypto/bcrypt"
    "dungeonSnackBE/model"
)
// Fungsi Registrasi
func Register(w http.ResponseWriter, r *http.Request, db *mongo.Client) {
    collection := db.Database("user").Collection("users")
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

    // Buat user baru dengan hashed password dan waktu pembuatan
    newUser.Password = string(hashedPassword)
    newUser.CreatedAt = time.Now()
    newUser.UpdatedAt = time.Now()

    _, err = collection.InsertOne(context.TODO(), newUser)
    if err != nil {
        http.Error(w, "Failed to register user", http.StatusInternalServerError)
        return
    }

    w.WriteHeader(http.StatusCreated)
    json.NewEncoder(w).Encode("User registered successfully")
}

// Fungsi Login
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
