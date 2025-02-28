package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type TodoList struct {
	ID     primitive.Object `json:"_id.omitempty" bson:"_id.omitempty"`
	Task   string           `json:"_task.omitempty" bson:"_task.omitempty"`
	Status bool             `json:"_status.omitempty" bson:"_status.omitempty"`
}
