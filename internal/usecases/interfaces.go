package usecases

import (
	"context"
	"go-blog-ca/internal/domain/entity"
)

type (

	// Author Interface
	Author interface {
		Create(ctx context.Context, author entity.Author) (string, error)

		GetByID(ctx context.Context, id string) (entity.Author, error)

		GetAll(ctx context.Context) ([]entity.Author, error)
	}

	// AuthorRepo Interface
	AuthorRepo interface {
		Create(ctx context.Context, author entity.Author) (string, error)

		FindByID(ctx context.Context, id string) (entity.Author, error)

		FindAll(ctx context.Context) ([]entity.Author, error)
	}

	// Article Interface  (change entity.Article to DTO)
	Article interface {
		Create(ctx context.Context, article entity.Article) (string, error)

		GetByID(ctx context.Context, id string) (entity.Article, error)

		GetAll(ctx context.Context) ([]entity.Article, error)

		//ctx, AuthorID, Article ID
		Delete(ctx context.Context, id string, aid string) error
	}

	// ArticleRepo Interface
	ArticleRepo interface {
		Create(ctx context.Context, article entity.Article) (string, error)

		FindById(ctx context.Context, id string) (entity.Article, error)

		FindAll(ctx context.Context) ([]entity.Article, error)

		Delete(ctx context.Context, id string, aid string) error
	}

	Session interface {
	}

	SessionRepo interface {
		Create(ctx context.Context, s entity.Session) error
		FindByID(ctx context.Context, sid string) (entity.Session, error)
		FindAll(ctx context.Context) ([]entity.Session, error)
		Delete(ctx context.Context, sid string) error
		//DeleteAll(ctx context.Context) error
	}
)
