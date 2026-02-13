package models

import "go.mongodb.org/mongo-driver/v2/bson"

type UserCredentials struct {
	ID           bson.ObjectID `bson:"_id,omitempty"`
	User_ID      bson.ObjectID `bson:"user_id"`
	Role         Role          `bson:"role"`
	Email        string        `bson:"email"`
	PasswordHash string        `bson:"password"`
	TOTPSecret   string        `bson:"totp_secret,omitempty"`
	TOTPEnabled  bool          `bson:"totp_enabled"`
}
