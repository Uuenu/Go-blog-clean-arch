package repo

import (
	"context"
	"fmt"
	"go-blog-ca/internal/domain/entity"
	"go-blog-ca/pkg/apperrors"
	"go-blog-ca/pkg/logging"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type SessionRepo struct {
	collection *mongo.Collection
	log        logging.Logger
}

func NewSessionRepo(db *mongo.Database, l logging.Logger) *SessionRepo {
	return &SessionRepo{
		collection: db.Collection("sessions"),
		log:        l,
	}
}

func (r *SessionRepo) Create(ctx context.Context, s entity.Session) (string, error) {

	//time.Sleep(time.Second * 10)

	result, err := r.collection.InsertOne(ctx, s)
	if err != nil {
		return "", fmt.Errorf("SessionRepo - Create - r.collection.InsertOne: %w", err)
	}
	session_id := primitive.ObjectID.Hex(result.InsertedID.(primitive.ObjectID))

	r.log.Infof("ObjectID: %v", result.InsertedID)
	r.log.Infof("SessionID: %v", session_id)
	return session_id, nil
}

func (r *SessionRepo) FindByID(ctx context.Context, sid string) (entity.Session, error) {
	//TODO: UniqueString isnt OBjectID

	oid, err := primitive.ObjectIDFromHex(sid)
	if err != nil {
		return entity.Session{}, fmt.Errorf("SessionRepo - FindByID - ObjectIDFromHex: %v", err)
	}

	filter := bson.M{"_id": oid}

	r.log.Infof("SessionRepo FindByID. oid: %v, sid: %v", oid, sid)

	result := r.collection.FindOne(ctx, filter)

	r.log.Infof("Mongo result: %v", result)

	var s entity.Session

	if err := result.Decode(&s); err != nil {
		return entity.Session{}, fmt.Errorf("SessionRepo - FindByID - result.Decode: %v", err)
	}

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
	r.log.Infof("SessionID from SessionRepo(Delete): %v", sid)

	oid, err := primitive.ObjectIDFromHex(sid)
	if err != nil {
		return fmt.Errorf("SessionRepo - Delete - ObjectIDFromHex: %v", err)
	}

	dresult, err := r.collection.DeleteOne(ctx, bson.M{"_id": oid})
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
