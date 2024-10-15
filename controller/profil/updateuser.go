package controller

import (
	"encoding/json"
	"net/http"
	"time"

	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func UpdateUserProfile(w http.ResponseWriter, r *http.Request, db *mongo.Client) {
	collection := db.Database("dsdatabase").Collection("user")
	var userProfileUpdate struct {
		ID    primitive.ObjectID `json:"id"`
		Nama  string             `json:"nama"`
		No_HP string             `json:"no_hp"`
		Email string             `json:"email"`
	}

	// Decode JSON dari body request
	if err := json.NewDecoder(r.Body).Decode(&userProfileUpdate); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Update nama, no_hp, dan email tanpa mengubah role
	_, err := collection.UpdateOne(
		context.TODO(),
		bson.M{"_id": userProfileUpdate.ID},
		bson.M{
			"$set": bson.M{
				"nama":      userProfileUpdate.Nama,
				"no_hp":     userProfileUpdate.No_HP,
				"email":     userProfileUpdate.Email,
				"updatedAt": time.Now(),
			},
		},
	)
	if err != nil {
		http.Error(w, "Gagal mengupdate profil user", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
