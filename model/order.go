package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// order struct
type Order struct {
	ID           primitive.ObjectID `json:"order_id,omitempty" bson:"_id,omitempty"`
	Namaorder string             `json:"nama_order" bson:"nama_order"`
	Slug         string             `bson:"slug" json:"slug"`
	Pembayaran       []Pembayaran           `bson:"pembayaran" json:"pembayaran"`
	CreatedAt    time.Time          `json:"created_at" bson:"created_at"`
	UpdatedAt    time.Time          `bson:"updatedAt,omitempty" json:"updatedAt,omitempty"`
}
