package repo

import (
	"context"
	"fmt"
	"go-blog-ca/internal/domain/entity"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type SessionRepo struct {
	collection *mongo.Collection
	//logger
}

func NewSessionRepo(db *mongo.Database) *SessionRepo {
	return &SessionRepo{
		collection: db.Collection("sessions"),
	}
}

func (r *SessionRepo) Create(ctx context.Context, s entity.Session) error {
	_, err := r.collection.InsertOne(ctx, s)
	if err != nil {
		return fmt.Errorf("SessionRepo - Create - r.collection.InsertOne: %w", err)
	}
	return nil
}
func (r *SessionRepo) FindByID(ctx context.Context, sid string) (entity.Session, error) {
	var s entity.Session

	if err := r.collection.FindOne(ctx, bson.M{"_id": sid}).Decode(&s); err != nil {

		if err == mongo.ErrNoDocuments {
			// TODO apperrors
			//return entity.Session{}, fmt.Errorf("r.FindOne.Decode: %w", apperrors.ErrSessionNotFound)
			return entity.Session{}, nil
		}

		return entity.Session{}, fmt.Errorf("SessionRepo - FindByID - FindOne: %w", err)
	}

	return s, nil
}
func (r *SessionRepo) FindAll(ctx context.Context, aid string) ([]entity.Session, error) {
	// find all authors session
	panic("implement me")
}
func (r *SessionRepo) Delete(ctx context.Context, sid string) error {

	dresult, err := r.collection.DeleteOne(ctx, bson.M{"_id": sid})
	if err != nil {

		if err == mongo.ErrNoDocuments {
			//TODO apperrors
			// reutnr fmt.Errorf("r.DeleteOne: %w", err)
			return nil
		}

		return fmt.Errorf("SessionRepo - Delete - Deleteone: %w", err)
	}

	if dresult.DeletedCount != 0 {
		return fmt.Errorf("SessionRepo - Delete - DeleteCount: %d", dresult.DeletedCount)
	}

	return nil
}

func (r *SessionRepo) DeleteAll(ctx context.Context, aid, sid string) error {
	// delete add authors session excluding current session
	panic("implement me")
}
