package usecases

import (
	"context"
	"go-blog-ca/internal/domain/entity"
)

type (

	// AuthorUseCase Interface
	Author interface {
		Create(ctx context.Context, author entity.Author) (string, error)

		GetByID(ctx context.Context, id string) (entity.Author, error)

		GetByEmail(ctx context.Context, email string) (entity.Author, error)

		GetAll(ctx context.Context) ([]entity.Author, error)

		Delete(ctx context.Context, id string) error
	}

	// AuthorRepo Interface
	AuthorRepo interface {
		Create(ctx context.Context, author entity.Author) (string, error)

		FindByID(ctx context.Context, id string) (entity.Author, error)

		FindByEmail(ctx context.Context, email string) (entity.Author, error)

		FindAll(ctx context.Context) ([]entity.Author, error)

		Delete(ctx context.Context, id string) error
	}

	// ArticleUseCase Interface  (change entity.Article to DTO)
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

	Auth interface {
		EmailLogin(ctx context.Context, email, password string) (entity.Session, error)

		Logout(ctx context.Context, sid string) error
	}

	// SessionUseCase interface
	Session interface {
		Create(ctx context.Context, aid string) (entity.Session, error)

		GetByID(ctx context.Context, sid string) (entity.Session, error)

		GetAll(ctx context.Context) []entity.Session

		Terminate(ctx context.Context, sid, currSid string) error

		TerminateAll(ctx context.Context, aid, sid string) error
	}

	// SessionRepo interface
	SessionRepo interface {
		Create(ctx context.Context, s entity.Session) error

		FindByID(ctx context.Context, sid string) (entity.Session, error)

		FindAll(ctx context.Context, aid string) ([]entity.Session, error)

		Delete(ctx context.Context, sid string) error

		// DeleteAll account sessions by provided account id excluding current session.
		DeleteAll(ctx context.Context, aid, currSid string) error
	}
)
