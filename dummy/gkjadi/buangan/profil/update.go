package profil

import (
	"encoding/json"
	"net/http"
	"time"
	"dungeonSnackBE/config"
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"dungeonSnackBE/model"
)

func UpdateProfile(w http.ResponseWriter, r *http.Request) {
	vars := r.URL.Query()
	userID := vars.Get("id")

	// Convert the user_id from string to MongoDB ObjectID
	oid, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	// Parse the body to get the updated data
	var updatedUser model.Users
	err = json.NewDecoder(r.Body).Decode(&updatedUser)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	// Set the updated timestamp
	updatedUser.UpdatedAt = time.Now()

	// Set up the MongoDB collection
	collection := config.Mongoconn.Collection("user")

	// Create a context with timeout
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Define the update query using $set to update only specific fields
	update := bson.M{
		"$set": bson.M{
			"nama":       updatedUser.Nama,
			"no_hp":    updatedUser.No_HP,
			"email":      updatedUser.Email,
			"updated_at": updatedUser.UpdatedAt,
		},
	}

	// Perform the update operation
	result, err := collection.UpdateOne(ctx, bson.M{"_id": oid}, update)
	if err != nil {
		http.Error(w, "Failed to update user", http.StatusInternalServerError)
		return
	}

	// Check if any document was actually updated
	if result.MatchedCount == 0 {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	// Return success response
	response := map[string]interface{}{
		"message": "User updated successfully",
		"user_id": userID,
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}