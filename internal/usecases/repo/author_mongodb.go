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

type AuthorRepo struct {
	collection *mongo.Collection // or *mongo.Database
	// logger
}

func NewAuthorRepo(db *mongo.Database) *AuthorRepo {
	return &AuthorRepo{
		collection: db.Collection("authors"),
	}
}

func (r *AuthorRepo) Create(ctx context.Context, author entity.Author) (string, error) {
	result, err := r.collection.InsertOne(ctx, author)
	if err != nil {
		return "", fmt.Errorf("failed to create author due to error: %v", err)
	}

	oid, ok := result.InsertedID.(primitive.ObjectID)
	if ok {
		return oid.Hex(), nil
	}

	return "", fmt.Errorf("failed to convert oid to hex (oid):%s", oid)
}

func (r *AuthorRepo) FindByID(ctx context.Context, id string) (entity.Author, error) {

	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return entity.Author{}, fmt.Errorf("AuthorRepo - FindByID - ObjectID from Hex id:%s error: %v", id, err)
	}

	result := r.collection.FindOne(ctx, bson.M{"_id": oid})

	if result.Err() != nil {
		if result.Err() == mongo.ErrNoDocuments {
			return entity.Author{}, fmt.Errorf("r.collection.FindOne. error: %v", apperrors.ErrAuthorNotFound)
		}
		return entity.Author{}, fmt.Errorf("AuthorRepo - FindByID - FindOne. error: %v", err)
	}

	var author entity.Author

	if err := result.Decode(author); err != nil {
		return entity.Author{}, fmt.Errorf("AuthorRepo - FindByID - Decode result. error: %v", err)
	}

	return author, nil
}

func (r *AuthorRepo) FindByEmail(ctx context.Context, email string) (entity.Author, error) {
	var author entity.Author
	if err := r.collection.FindOne(ctx, bson.M{"email": email}).Decode(&author); err != nil {
		// TODO apperror's
		if err == mongo.ErrNoDocuments {
			return entity.Author{}, fmt.Errorf("r.collection.FindOne. error: %v", apperrors.ErrAuthorNotFound)
		}
		return entity.Author{}, fmt.Errorf("AuthorRepo - FindByEmail - FindOne and Decode. error: %v", err)
	}
	return author, nil
}

func (r *AuthorRepo) FindAll(ctx context.Context) ([]entity.Author, error) {
	panic("implement me")
}

func (r *AuthorRepo) Delete(ctx context.Context, id string) error {
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return fmt.Errorf("AuthorRepo - Delete - ObjectID from Hex. error: %v", err)
	}

	filter := bson.M{"_id": oid}
	dresult, err := r.collection.DeleteOne(ctx, filter)
	if err != nil {
		return fmt.Errorf("AuthorRepo - Delete - DeleteOne. error: %v", err)
	}

	if dresult.DeletedCount != 1 {
		return fmt.Errorf("AuthorRepo - Delete - dresult.DeleteCount = %d. error: %v", dresult.DeletedCount, err)
	}

	return nil
}
