package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// order struct
type Order struct {
	ID           primitive.ObjectID `json:"order_id,omitempty" bson:"_id,omitempty"`
	Quantity     int                `json:"quantity,omitempty" bson:"quantity,omitempty"`
	Payment	  []Payment          `json:"payment,omitempty" bson:"payment,omitempty"`
	Slug         string             `bson:"slug" json:"slug"`
}

// OrderDetail struct
type OrderDetail struct {
	ID           primitive.ObjectID `json:"order_detail_id,omitempty" bson:"_id,omitempty"`
	OrderID      primitive.ObjectID `json:"order_id,omitempty" bson:"order_id,omitempty"`
	ProductID    primitive.ObjectID `json:"product_id,omitempty" bson:"product_id,omitempty"`
	Quantity     int                `json:"quantity,omitempty" bson:"quantity,omitempty"`
	Price        float32            `json:"price,omitempty" bson:"price,omitempty"`
	TotalPrice   float32            `json:"total_price,omitempty" bson:"total_price,omitempty"`
	CreatedAt    time.Time          `json:"created_at,omitempty" bson:"created_at,omitempty"`
	UpdatedAt    time.Time          `json:"updated_at,omitempty" bson:"updated_at,omitempty"`
}

// payment struct
type Payment struct {
	ID           primitive.ObjectID `json:"payment_id,omitempty" bson:"_id,omitempty"`
	OrderID      primitive.ObjectID `json:"order_id,omitempty" bson:"order_id,omitempty"`
	TotalPrice   float32            `json:"total_price,omitempty" bson:"total_price,omitempty"`
	PaymentDate  time.Time          `json:"payment_date,omitempty" bson:"payment_date,omitempty"`
	PaymentProof string             `json:"payment_proof,omitempty" bson:"payment_proof,omitempty"`
	Status       string             `json:"status,omitempty" bson:"status,omitempty"`
	CreatedAt    time.Time          `json:"created_at,omitempty" bson:"created_at,omitempty"`
	UpdatedAt    time.Time          `json:"updated_at,omitempty" bson:"updated_at,omitempty"`
}
