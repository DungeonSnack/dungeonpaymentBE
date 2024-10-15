package controller

import (
	"dungeonSnackBE/model/produk"
	"encoding/json"
	"net/http"
	"time"

	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Definisikan kategori yang diperbolehkan
var allowedCategories = []string{"minuman", "makanan berat", "makanan ringan", "jelly"}

// Fungsi untuk memeriksa apakah kategori valid
func isValidCategory(category string) bool {
	for _, v := range allowedCategories {
		if v == category {
			return true
		}
	}
	return false
}

// CreateProduct untuk menambahkan produk baru
func CreateProduct(w http.ResponseWriter, r *http.Request, db *mongo.Client) {
	collection := db.Database("dsdatabase").Collection("product")
	var newProduct model.Product

	// Decode JSON body ke struct Product
	if err := json.NewDecoder(r.Body).Decode(&newProduct); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Validasi kategori
	if !isValidCategory(newProduct.Kategori) {
		http.Error(w, "Kategori tidak valid. Pilih antara: minuman, makanan berat, makanan ringan, jelly.", http.StatusBadRequest)
		return
	}

	newProduct.ID = primitive.NewObjectID()
	newProduct.CreatedAt = time.Now()
	newProduct.UpdatedAt = time.Now()

	// Insert ke database
	_, err := collection.InsertOne(context.TODO(), newProduct)
	if err != nil {
		http.Error(w, "Gagal menambahkan produk", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

// GetProducts untuk mengambil semua produk
func GetProducts(w http.ResponseWriter, r *http.Request, db *mongo.Client) {
	collection := db.Database("dsdatabase").Collection("product")
	var products []model.Product

	cursor, err := collection.Find(context.TODO(), bson.M{})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer cursor.Close(context.Background())

	for cursor.Next(context.Background()) {
		var product model.Product
		cursor.Decode(&product)
		products = append(products, product)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(products)
}

// DeleteProduct untuk menghapus produk berdasarkan ID
func DeleteProduct(w http.ResponseWriter, r *http.Request, db *mongo.Client) {
	collection := db.Database("dsdatabase").Collection("product")
	var productToDelete model.Product

	if err := json.NewDecoder(r.Body).Decode(&productToDelete); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	_, err := collection.DeleteOne(context.TODO(), bson.M{"_id": productToDelete.ID})
	if err != nil {
		http.Error(w, "Gagal menghapus produk", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

// UpdateProduct untuk mengupdate produk
func UpdateProduct(w http.ResponseWriter, r *http.Request, db *mongo.Client) {
    collection := db.Database("dsdatabase").Collection("product")
    var productToUpdate model.Product

    if err := json.NewDecoder(r.Body).Decode(&productToUpdate); err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    // Validasi kategori
    if !isValidCategory(productToUpdate.Kategori) {
        http.Error(w, "Kategori tidak valid. Pilih antara: minuman, makanan berat, makanan ringan, jelly.", http.StatusBadRequest)
        return
    }

    productToUpdate.UpdatedAt = time.Now()

    _, err := collection.UpdateOne(context.TODO(), bson.M{"_id": productToUpdate.ID}, bson.M{
        "$set": bson.M{
            "nama_produk": productToUpdate.NamaProduk,
            "harga":       productToUpdate.Harga,
            "deskripsi":   productToUpdate.Deskripsi,
            "kategori":    productToUpdate.Kategori,
            "updatedAt":   productToUpdate.UpdatedAt,
        },
    })
    if err != nil {
        http.Error(w, "Gagal mengupdate produk", http.StatusInternalServerError)
        return
    }

    w.WriteHeader(http.StatusOK)
}

