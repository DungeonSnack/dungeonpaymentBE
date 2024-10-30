package model

import (
	"time"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Menu struct {
	ID          primitive.ObjectID `bson:"produk_id" json:"id,omitempty"`
	NamaProduk string             `bson:"nama_produk" json:"nama_produk,omitempty"`
	Price 	 float32            `bson:"price" json:"price,omitempty"`
	Description string             `bson:"description" json:"description,omitempty"`
	Stok        int                `bson:"stok" json:"stok,omitempty"`
	Slug         string             `bson:"slug" json:"slug"`
	CreatedAt   time.Time          `bson:"createdAt,omitempty" json:"createdAt,omitempty"`
	UpdatedAt   time.Time          `bson:"updatedAt,omitempty" json:"updatedAt,omitempty"`
	Image       string             `bson:"image_menu" json:"image,omitempty"`
}

type Produk struct {
	ID          primitive.ObjectID `bson:"produk_id" json:"id,omitempty"`
	NamaProduk string             `bson:"nama_produk" json:"nama_produk,omitempty"`
	Price 	 float32            `bson:"price" json:"price,omitempty"`
	Description string             `bson:"description" json:"description,omitempty"`
	Stok        int                `bson:"stok" json:"stok,omitempty"`
	Slug         string             `bson:"slug" json:"slug"`
	CreatedAt   time.Time          `bson:"createdAt,omitempty" json:"createdAt,omitempty"`
	UpdatedAt   time.Time          `bson:"updatedAt,omitempty" json:"updatedAt,omitempty"`
	Image       string             `bson:"image_produk" json:"image,omitempty"`
}
type Toko struct {
	ID          primitive.ObjectID `bson:"toko_id" json:"id,omitempty"`
	NamaToko string             `bson:"nama_toko" json:"nama_toko,omitempty"`
	Alamat string             `bson:"alamat" json:"alamat,omitempty"`
	Category string             `bson:"category" json:"category,omitempty"`
	Slug         string             `bson:"slug" json:"slug"`
	CreatedAt   time.Time          `bson:"createdAt,omitempty" json:"createdAt,omitempty"`
	UpdatedAt   time.Time          `bson:"updatedAt,omitempty" json:"updatedAt,omitempty"`
	Image       string             `bson:"image_toko" json:"image,omitempty"`
}

type Category struct {
	ID          primitive.ObjectID `bson:"category_id" json:"id,omitempty"`
	NamaCategory string             `bson:"nama_category" json:"nama_category,omitempty"`
}