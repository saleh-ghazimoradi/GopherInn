package domain

import "go.mongodb.org/mongo-driver/v2/bson"

type User struct {
	Id        bson.ObjectID `bson:"_id,omitempty"`
	FirstName string        `bson:"first_name"`
	LastName  string        `bson:"last_name"`
	Email     string        `bson:"email"`
	Password  string        `bson:"password"`
}
