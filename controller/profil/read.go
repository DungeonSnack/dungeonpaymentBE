package profil

import (
	"encoding/json"
		"dungeonSnackBE/config"
	model "dungeonSnackBE/model"
	"net/http"
    "time"
    "context"
    "go.mongodb.org/mongo-driver/bson"
)

func GetProfile(w http.ResponseWriter, r *http.Request) {
    collection := config.Mongoconn.Collection("user")

    // Create a context with a timeout
    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()

    // Find all users in the collection
    cursor, err := collection.Find(ctx, bson.M{})
    if err != nil {
        http.Error(w, "Error fetching users", http.StatusInternalServerError)
        return
    }
    defer cursor.Close(ctx)

    // Create a slice to hold the users
    var users []model.Users

    // Iterate over the cursor and decode each user document
    for cursor.Next(ctx) {
        var user model.Users
        if err := cursor.Decode(&user); err != nil {
            http.Error(w, "Error decoding user data", http.StatusInternalServerError)
            return
        }
        users = append(users, user)
    }

    // Check for cursor errors
    if err := cursor.Err(); err != nil {
        http.Error(w, "Cursor error", http.StatusInternalServerError)
        return
    }

    // Return the users as JSON
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(users)
}