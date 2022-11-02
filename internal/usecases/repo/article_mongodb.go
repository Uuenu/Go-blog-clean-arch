package repo

import (
	"context"
	"fmt"
	"go-blog-ca/internal/domain/entity"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type ArticleRepo struct {
	collection *mongo.Collection // or *mongo.Database
	// logger
}

func NewArticleRepo(db *mongo.Database) *ArticleRepo {
	return &ArticleRepo{
		collection: db.Collection("articles"),
	}
}

func (r *ArticleRepo) Create(ctx context.Context, article entity.Article) (string, error) {
	result, err := r.collection.InsertOne(ctx, article)
	if err != nil {
		return "", fmt.Errorf("failed to create article due to error %v", err)
	}

	oid, ok := result.InsertedID.(primitive.ObjectID)
	if ok {
		return oid.Hex(), nil
	}

	return "", fmt.Errorf("failed to convert oid to hex (oid):%s", oid)
}

func (r *ArticleRepo) FindById(ctx context.Context, id string) (entity.Article, error) {
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return entity.Article{}, fmt.Errorf("failed to convert hex to ObjectID. hex: %s", id)
	}

	filter := bson.M{"_id": oid}

	result := r.collection.FindOne(ctx, filter)

	if result.Err() != nil {
		return entity.Article{}, fmt.Errorf("failed to FindOne Article by id: %s due to error: %v", id, result.Err())
	}

	var article entity.Article

	if err := result.Decode(article); err != nil {
		return entity.Article{}, fmt.Errorf("failde to decode article id:%s due to error: %v ", id, err)
	}

	return article, nil
}

func (r *ArticleRepo) FindAll(ctx context.Context) ([]entity.Article, error) {

	cursor, err := r.collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, fmt.Errorf("failed to find all users due to error: %v", err)
	}

	var articles []entity.Article
	if err := cursor.All(ctx, &articles); err != nil {
		return nil, fmt.Errorf("failed to read all documents from cursor. error: %v", err)
	}

	return articles, nil
}

func (r *ArticleRepo) Delete(ctx context.Context, id string, aid string) error {
	// TODO
	panic("Implement me")
}
