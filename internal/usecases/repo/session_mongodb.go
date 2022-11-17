package repo

import (
	"context"
	"fmt"
	"go-blog-ca/internal/domain/entity"
	"go-blog-ca/pkg/apperrors"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type SessionRepo struct {
	collection *mongo.Collection
	//l          *logging.Logger
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
	oid, err := primitive.ObjectIDFromHex(sid)
	if err != nil {
		return entity.Session{}, fmt.Errorf("primitive.OIDFromHex: %s", sid)
	}
	if err := r.collection.FindOne(ctx, bson.M{"_id": oid}).Decode(&s); err != nil {

		if err == mongo.ErrNoDocuments {
			return entity.Session{}, fmt.Errorf("r.FindOne.Decode: %w", apperrors.ErrSessionNotFound)
		}

		return entity.Session{}, fmt.Errorf("SessionRepo - FindByID - FindOne: %w", err)
	}
	//r.l.Infof("Session: %s")
	return s, nil
}
func (r *SessionRepo) FindAll(ctx context.Context, aid string) ([]entity.Session, error) {

	filter := bson.M{"account_id": bson.M{"": aid}}

	cursor, err := r.collection.Find(ctx, filter)
	if err != nil {
		return nil, fmt.Errorf("SessionRepo - FindAll - Find. error: %v", err)
	}

	var sessions []entity.Session

	if err := cursor.All(ctx, sessions); err != nil {
		return nil, fmt.Errorf("SessionRepo - FindAll - cursor.All. error: %v", err)
	}

	return sessions, nil
}
func (r *SessionRepo) Delete(ctx context.Context, sid string) error {

	dresult, err := r.collection.DeleteOne(ctx, bson.M{"_id": sid})
	if err != nil {

		if err == mongo.ErrNoDocuments {
			return apperrors.ErrSessionNotFound
		}

		return fmt.Errorf("SessionRepo - Delete - Deleteone: %w", err)
	}

	if dresult.DeletedCount == 0 {
		return fmt.Errorf("SessionRepo - Delete - DeleteCount: %d", dresult.DeletedCount)
	}

	return nil
}

func (r *SessionRepo) DeleteAll(ctx context.Context, aid, currSid string) error {
	// delete add authors session excluding current session
	author_oid, err := primitive.ObjectIDFromHex(aid)
	if err != nil {
		return fmt.Errorf("SessionRepo - DeleteAll - ObjectID from Hex. error: %v", err)
	}

	//TODO add excluding sid session
	filter := bson.M{
		"_id":       bson.M{"$ne": currSid}, // all session without currSid
		"author_id": author_oid,
	}

	dresults, err := r.collection.DeleteMany(ctx, filter)
	if err != nil {
		return fmt.Errorf("SessionRepo - DeleteAll - DeleteMany. error: %v", err)
	}

	if dresults.DeletedCount == 0 {
		return fmt.Errorf("SessionRepo - DeleteAll - DeletedCount. error: %v", apperrors.ErrSessionNotTerminated)
	}

	return nil
}
