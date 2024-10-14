package model

import (
	"time"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Kategori struct {
	ID          primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	NamaKategori string             `json:"nama_kategori" bson:"nama_kategori"`
	CreatedAt   time.Time           `json:"created_at" bson:"created_at"`
	UpdatedAt 	time.Time 			`bson:"updatedAt,omitempty" json:"updatedAt,omitempty"` 
}