package controller

import (
	"context"
	"dungeonSnackBE/config"
	model "dungeonSnackBE/model/pengguna"
	"encoding/json"
	"net/http"
	"time"

	whatsauth "github.com/whatsauth/itmodel"
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
	var whatsapi whatsauth.Response
	err := json.NewDecoder(r.Body).Decode(&newUser)
	if err != nil {
		whatsapi.Response = "Invalid request payload"
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{
			"error":   whatsapi.Response,
			"message": "The JSON request body could not be decoded. Please check the structure of your request.",
		})
		return
	}

	// Cek apakah email sudah ada
	if newUser.Nama == "" {
		whatsapi.Response = "Nama is required"
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{
			"error":   whatsapi.Response,
			"message": "Please provide a valid name.",
		})
		return
	}

	if newUser.No_HP == "" {
		whatsapi.Response = "Phone number is required"
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{
			"error":   whatsapi.Response,
			"message": "Please provide a valid phone number.",
		})
		return
	}

	if newUser.Email == "" {
		whatsapi.Response = "Email is required"
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{
			"error":   whatsapi.Response,
			"message": "Please provide a valid email address.",
		})
		return
	}

	// Hash password
	if newUser.Password == "" {
		whatsapi.Response = "Password is required"
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{
			"error":   whatsapi.Response,
			"message": "Please provide a password.",
		})
		return
	}

	// Hash the user's password before saving it to the database
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(newUser.Password), bcrypt.DefaultCost)
	if err != nil {
		whatsapi.Response = "Failed to hash password"
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{
			"error":   whatsapi.Response,
			"message": "An error occurred while hashing the password.",
		})
		return
	}

	// Tetapkan role pembeli secara default
	newUser.Role = "pembeli" // Mengatur role pembeli

	// Buat user baru dengan hashed password dan waktu pembuatan
	newUser.Password = string(hashedPassword)
	newUser.CreatedAt = time.Now()
	newUser.UpdatedAt = time.Now()
	newUser.ID = primitive.NewObjectID() // Membuat ID baru

	collection = config.Mongoconn.Collection("user")
	// error cuy, collection := config.Mongoconn.Collection("user")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	result, err := collection.InsertOne(ctx, bson.M{
		"nama":       newUser.Nama,
		"no_telp":    newUser.No_HP,
		"email":      newUser.Email,
		"role":       newUser.Role,
		"password":   newUser.Password,
		"created_at": newUser.CreatedAt,
		"updated_at": newUser.UpdatedAt,
	})

	if err != nil {
		whatsapi.Response = "Failed to insert user"
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{
			"error":   whatsapi.Response,
			"message": "An error occurred while inserting the user into the database.",
		})
		return
	}

	response := map[string]interface{}{
		"message": "User registered successfully",
		"user_id": result.InsertedID,
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
