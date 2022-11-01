package service

import (
	"go-blog-ca/internal/domain/entity"
)

type ArticleStorage interface {
	GetOne(ID string) (entity.Article, error)
	GetAll() ([]entity.Article, error)
}
