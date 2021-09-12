package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type CookieResponse struct {
	CodeAuth string             `json:"codeAuth,omitempty"`
	ID       primitive.ObjectID `json:"_id,omitempty"`
}
