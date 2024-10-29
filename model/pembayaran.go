package model

import (
	"time"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Pembayaran struct {
	ID          primitive.ObjectID `bson:"_id" json:"id,omitempty"`
	ProductName string             `bson:"product_name" json:"product_name,omitempty"`
	MetodeBayar   string             `bson:"metode_bayar" json:"category,omitempty"`
	CreatedAt   time.Time          `bson:"createdAt,omitempty" json:"createdAt,omitempty"`
	UpdatedAt   time.Time          `bson:"updatedAt,omitempty" json:"updatedAt,omitempty"`
}