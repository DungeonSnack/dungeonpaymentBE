package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// order struct
type Order struct {
	ID           primitive.ObjectID `json:"order_id,omitempty" bson:"_id,omitempty"`
	Namaorderan string             `json:"nama_orderan" bson:"nama_orderan"`
	Price       float32            `bson:"price" json:"price,omitempty"`
	Description string             `bson:"description" json:"description,omitempty"`
	Slug         string             `bson:"slug" json:"slug"`
	Pembayaran       []Pembayaran           `bson:"pembayaran" json:"pembayaran"`
	CreatedAt    time.Time          `json:"created_at" bson:"created_at"`
	UpdatedAt    time.Time          `bson:"updatedAt,omitempty" json:"updatedAt,omitempty"`
}
