package controller

import (
	"dungeonSnackBE/model/kategori"
	"encoding/json"
	"net/http"
	"time"

	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func CreateKategori(w http.ResponseWriter, r *http.Request, db *mongo.Client) {
	collection := db.Database("dsdatabase").Collection("kategori")
	var newKategori model.Kategori

	if err := json.NewDecoder(r.Body).Decode(&newKategori); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var existingKategori model.Kategori
	err := collection.FindOne(context.TODO(), bson.M{"nama_kategori": newKategori.NamaKategori}).Decode(&existingKategori)
	if err == nil {
		http.Error(w, "Kategori sudah ada", http.StatusConflict)
		return
	}

	newKategori.CreatedAt = time.Now()
	newKategori.ID = primitive.NewObjectID()

	_, err = collection.InsertOne(context.TODO(), newKategori)
	if err != nil {
		http.Error(w, "Gagal membuat kategori", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func GetKategori(w http.ResponseWriter, r *http.Request, db *mongo.Client) {
	collection := db.Database("dsdatabase").Collection("kategori")
	var kategoris []model.Kategori

	cursor, err := collection.Find(context.TODO(), bson.M{})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer cursor.Close(context.Background())

	for cursor.Next(context.Background()) {
		var kategori model.Kategori
		cursor.Decode(&kategori)
		kategoris = append(kategoris, kategori)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(kategoris)
}

func DeleteKategori(w http.ResponseWriter, r *http.Request, db *mongo.Client) {
	collection := db.Database("dsdatabase").Collection("kategori")
	var kategoriToDelete model.Kategori

	if err := json.NewDecoder(r.Body).Decode(&kategoriToDelete); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	_, err := collection.DeleteOne(context.TODO(), bson.M{"_id": kategoriToDelete.ID})
	if err != nil {
		http.Error(w, "Gagal menghapus kategori", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func UpdateKategori(w http.ResponseWriter, r *http.Request, db *mongo.Client) {
	collection := db.Database("dsdatabase").Collection("kategori")
	var kategoriToUpdate model.Kategori

	if err := json.NewDecoder(r.Body).Decode(&kategoriToUpdate); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	_, err := collection.UpdateOne(context.TODO(), bson.M{"_id": kategoriToUpdate.ID}, bson.M{"$set": bson.M{"nama_kategori": kategoriToUpdate.NamaKategori}})
	if err != nil {
		http.Error(w, "Gagal mengupdate kategori", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
