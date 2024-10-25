package model

import (
	"time"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Produk struct {
	ID          primitive.ObjectID `bson:"produk_id,omitempty" json:"id,omitempty"`
	NamaProduk  string             `bson:"nama_produk,omitempty" json:"nama_produk,omitempty"`
	Harga       int                `bson:"harga,omitempty" json:"harga,omitempty"`
	Deskripsi   string             `bson:"deskripsi,omitempty" json:"deskripsi,omitempty"`
	Kategori    string             `bson:"kategori,omitempty" json:"kategori,omitempty"`
	CreatedAt   time.Time          `bson:"createdAt,omitempty" json:"createdAt,omitempty"`
	UpdatedAt   time.Time          `bson:"updatedAt,omitempty" json:"updatedAt,omitempty"`
}