package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// OrderItem struct
type OrderItem struct {
	ProductID primitive.ObjectID `json:"product_id" bson:"product_id"`
	Quantity  int                `json:"quantity" bson:"quantity"`
	Price     int                `json:"price" bson:"price"`
}

// order struct
type Order struct {
	ID           primitive.ObjectID `json:"order_id,omitempty" bson:"_id,omitempty"`
	UserId	   primitive.ObjectID `json:"_id" bson:"user_id"`
	TokoId	   primitive.ObjectID `json:"toko_id" bson:"toko_id"`
	OrderItems []OrderItem `json:"order_items" bson:"order_items"`
	TotalPrice  int `json:"total_price" bson:"total_price"`
	Status      string `json:"status" bson:"status"`
	CreatedAt   time.Time `json:"created_at" bson:"created_at"`
	UpdatedAt   time.Time `json:"updated_at" bson:"updated_at"`
}