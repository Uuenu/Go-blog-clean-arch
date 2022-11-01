package mongodb

import "go.mongodb.org/mongo-driver/mongo"

type userStorage struct {
	db *mongo.Database
}

func NewUserStorage(db *mongo.Database) *userStorage {
	return &userStorage{
		db: db,
	}
}
