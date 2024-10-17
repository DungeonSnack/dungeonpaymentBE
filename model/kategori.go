package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Kategori struct
type Kategori struct {
	ID           primitive.ObjectID `json:"kategori_id,omitempty" bson:"_id,omitempty"`
	NamaKategori string             `json:"nama_kategori" bson:"nama_kategori"`
	Slug         string             `bson:"slug" json:"slug"`
	Produk       []Produk           `bson:"produk" json:"produk"`
	CreatedAt    time.Time          `json:"created_at" bson:"created_at"`
	UpdatedAt    time.Time          `bson:"updatedAt,omitempty" json:"updatedAt,omitempty"`
}
