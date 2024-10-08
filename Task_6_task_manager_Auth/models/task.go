package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Task struct {
	ID          primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Title       string             `bson:"title" json:"title" binding:"required"`
	Description string             `bson:"description" json:"description" binding:"required"`
	DueDate     string             `bson:"due_date" json:"due_date" binding:"required"`
	Status      string             `bson:"status" json:"status"`
}
