package model

import (
	"time"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Toko struct {
	ID          primitive.ObjectID `json:"toko_id,omitempty" bson:"_id,omitempty"`
	NamaToko    string             `json:"nama_toko" bson:"nama_toko"`
	Slug        string             `json:"slug" bson:"slug"`
	Produk      []Produk          `json:"menu" bson:"produk"`
	CreatedAt   time.Time          `json:"created_at" bson:"created_at"`
	UpdatedAt   time.Time          `json:"updated_at" bson:"updated_at"`
}
