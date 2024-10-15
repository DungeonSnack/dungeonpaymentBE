package model

import (
	"time"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Users struct {
	ID       primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	Nama     string             `bson:"nama,omitempty" json:"nama,omitempty"`
	No_HP    string             `bson:"no_hp,omitempty" json:"no_hp,omitempty"`
	Email    string             `bson:"email,omitempty" json:"email,omitempty"`
	Role     string             `bson:"role,omitempty" json:"role,omitempty"`
	Password string             `bson:"password,omitempty" json:"password,omitempty"`
	CreatedAt time.Time         `bson:"createdAt,omitempty" json:"createdAt,omitempty"`
	UpdatedAt time.Time         `bson:"updatedAt,omitempty" json:"updatedAt,omitempty"`
}
