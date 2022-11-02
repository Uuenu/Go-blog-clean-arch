package mongodb

import "go.mongodb.org/mongo-driver/mongo"

type authorStorage struct {
	db *mongo.Database
}

func NewUserStorage(db *mongo.Database) *authorStorage {
	return &authorStorage{
		db: db,
	}
}

func (s authorStorage) Create() error {
	panic("implement me")
}

func (s authorStorage) Update() error {
	panic("implement me")
}
