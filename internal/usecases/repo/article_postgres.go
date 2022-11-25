package repo

import (
	"context"
	"errors"
	"fmt"
	"go-blog-ca/internal/domain/entity"
	"go-blog-ca/pkg/apperrors"
	"go-blog-ca/pkg/postgres"

	"github.com/Masterminds/squirrel"
	"github.com/jackc/pgconn"
	"github.com/jackc/pgerrcode"
	"github.com/jackc/pgx/v4"
)

const (
	//_defaultEntityCap = 64
	_articleTable = "article"
)

type ArticlePostgresRepo struct {
	*postgres.Postgres
}

func NewArticlePostgresRepo(pg *postgres.Postgres) *ArticlePostgresRepo {
	return &ArticlePostgresRepo{pg}
}

func (r ArticlePostgresRepo) Create(ctx context.Context, article entity.Article) (string, error) {
	sql, args, err := r.Builder.
		Insert(_articleTable).
		Columns("author_id, header, text").
		Values(article.AuthorID, article.Header, article.Text).
		Suffix("RETURNING id").
		ToSql()

	if err != nil {
		return "", fmt.Errorf("ArticleRepo - Create - r.Builder.Insert(). error: %v", err)
	}

	var articleID string

	err = r.Pool.QueryRow(context.Background(), sql, args...).Scan(&articleID)
	if err != nil {
		var pgErr *pgconn.PgError

		// TODO replace author on article
		if errors.As(err, &pgErr) {
			if pgErr.Code == pgerrcode.UniqueViolation {
				return "", fmt.Errorf("r.Pool.Exec: %w", apperrors.ErrAuthorAlreadyExist)
			}
		}

		return "", fmt.Errorf("r.Pool.Exec: %w", err)
	}

	return articleID, nil
}

func (r ArticlePostgresRepo) FindById(ctx context.Context, id string) (entity.Article, error) {
	sql, args, err := r.Builder.Select("id, author_id, header, text").
		From(_articleTable).
		Where(squirrel.Eq{"id": id}).
		ToSql()

	if err != nil {
		return entity.Article{}, fmt.Errorf("ArticleRepo - FindByID - r.Builder.Select(). error: %v", err)
	}

	article := entity.Article{}

	err = r.Pool.QueryRow(ctx, sql, args...).Scan(
		&article.ID,
		&article.AuthorID,
		&article.Header,
		&article.Text,
	)
	if err != nil {
		if err == pgx.ErrNoRows {
			// TODO replace author error on article
			return entity.Article{}, fmt.Errorf("ArticleRepo - FindByID - r.Pool.QueryRow. error: %v", apperrors.ErrAuthorNotFound)
		}

		return entity.Article{}, fmt.Errorf("ArticleRepo - FindByID - r.Pool.QueryRow. error: %v", err)
	}

	return article, nil
}

func (r ArticlePostgresRepo) FindAll(ctx context.Context) ([]entity.Article, error) {
	panic("Implement me")
}

func (r ArticlePostgresRepo) Delete(ctx context.Context, id string, aid string) error {
	panic("Implement me")
}
