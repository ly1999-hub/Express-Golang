package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Person struct {
	ID          primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Name        string             `json:"name,omitempty" bson:"name,omitempty"`
	Age         int                `json:"age,omitempty" bson:"age,omitempty"`
	Description string             `json:"description,omitempty" bson:"description,omitempty"`
	Completed   bool               `bson:"completed" json:"completed"` // default: false
	CreatedAt   time.Time          `bson:"createdAt" json:"createdAt"` // server tự sinh
	UpdatedAt   time.Time          `bson:"updatedAt" json:"updatedAt"`
}
