package model

import (
	"time"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Produk struct {
	ID          primitive.ObjectID `bson:"produk_id" json:"id,omitempty"`
	NamaProduk string             `bson:"nama_produk" json:"nama_produk,omitempty"`
	Price 	 float32            `bson:"price" json:"price,omitempty"`
	Description string             `bson:"description" json:"description,omitempty"`
	Stok        int                `bson:"stok" json:"stok,omitempty"`
	Slug         string             `bson:"slug" json:"slug"`
	CreatedAt   time.Time          `bson:"createdAt,omitempty" json:"createdAt,omitempty"`
	UpdatedAt   time.Time          `bson:"updatedAt,omitempty" json:"updatedAt,omitempty"`
}
