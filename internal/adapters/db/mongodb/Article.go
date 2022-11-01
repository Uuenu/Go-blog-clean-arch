package mongodb

import (
	"go-blog-ca/internal/domain/entity"

	"go.mongodb.org/mongo-driver/mongo"
)

type articleStorage struct {
	db *mongo.Database
}

func NewArticleStorage(db *mongo.Database) *articleStorage {
	return &articleStorage{
		db: db,
	}
}

func (a *articleStorage) Create(article entity.Article) error {
	panic("implement me")
}

func (a *articleStorage) GetOne(ID string) (entity.Article, error) {
	panic("Implement me")
}

func (a *articleStorage) GetAll() ([]entity.Article, error) {
	panic("Implement me")
}

func (a *articleStorage) Update(ID string, artile entity.Article) {
	panic("implement me")
}
