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

func UpdateUserRole(w http.ResponseWriter, r *http.Request, db *mongo.Client) {
	collection := db.Database("dsdatabase").Collection("user")
	var userRoleUpdate struct {
		ID   primitive.ObjectID `json:"id"`
		Role string             `json:"role"`
	}

	// Decode JSON dari body request
	if err := json.NewDecoder(r.Body).Decode(&userRoleUpdate); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Update field "role" pada user yang sesuai dengan ID
	_, err := collection.UpdateOne(
		context.TODO(),
		bson.M{"_id": userRoleUpdate.ID},
		bson.M{"$set": bson.M{"role": userRoleUpdate.Role, "updatedAt": time.Now()}},
	)
	if err != nil {
		http.Error(w, "Gagal mengupdate role user", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
