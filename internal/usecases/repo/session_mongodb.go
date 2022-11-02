package repo

import "go.mongodb.org/mongo-driver/mongo"

type SessionRepo struct {
	collection *mongo.Collection
}

func NewSessionRepo(db *mongo.Database) *SessionRepo {
	return &SessionRepo{
		collection: db.Collection("sessions"),
	}
}
