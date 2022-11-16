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
	//"github.com/google/uuid"
)

const (
	_defaultEntityCap = 64
	_athTable         = "author"
)

type AuthorPostgresRepo struct {
	*postgres.Postgres
}

func New(pg *postgres.Postgres) *AuthorPostgresRepo {
	return &AuthorPostgresRepo{pg}
}

func (r AuthorPostgresRepo) Create(ctx context.Context, author entity.Author) (string, error) {
	sql, args, err := r.Builder.
		Insert(_athTable).
		Columns("username, email, password_hash, salt").
		Values(author.Username, author.Email, author.PasswordHash, author.Salt).
		Suffix("RETURNING id").
		ToSql()

	if err != nil {
		return "", fmt.Errorf("AuthorRepo - Create - r.Builder.Insert(). error: %v", err)
	}

	var aid string

	err = r.Pool.QueryRow(ctx, sql, args...).Scan(&aid)
	if err != nil {
		var pgErr *pgconn.PgError

		if errors.As(err, &pgErr) {
			if pgErr.Code == pgerrcode.UniqueViolation {
				return "", fmt.Errorf("r.Pool.Exec: %w", apperrors.ErrAuthorAlreadyExist)
			}
		}

		return "", fmt.Errorf("r.Pool.Exec: %w", err)
	}

	return aid, nil
}

func (r AuthorPostgresRepo) FindAll(ctx context.Context) ([]entity.Author, error) {
	sql, _, err := r.Builder.
		Select("id, username, email").
		From(_athTable).
		ToSql()
	if err != nil {
		return nil, fmt.Errorf("AuthorRepo - GetAll - r.Builder. error: %v", err)
	}

	rows, err := r.Pool.Query(ctx, sql, nil...)
	if err != nil {
		return nil, fmt.Errorf("AuthorRepo - GetAll - r.Pool.Query, error: %v", err)
	}
	defer rows.Close()

	authors := make([]entity.Author, 0, _defaultEntityCap)

	for rows.Next() {
		ath := entity.Author{}

		err := rows.Scan(ath.ID, ath.Username, ath.Email)
		if err != nil {
			return nil, fmt.Errorf("AuthorRepo - GetAll - rows.Scan. error: %v", err)
		}

		authors = append(authors, ath)
	}

	return authors, nil
}

func (r AuthorPostgresRepo) FindByID(ctx context.Context, aid string) (entity.Author, error) {

	//oid, _ := primitive.ObjectIDFromHex(aid)
	//uaid, _ := uuid.FromString(aid)

	sql, args, err := r.Builder.Select("id, username, email, password_hash, salt").
		From(_athTable).
		Where(squirrel.Eq{"id": aid}).
		ToSql()

	if err != nil {
		return entity.Author{}, fmt.Errorf("AuthorRepo - FindByID - r.Builder.Select. error: %v", err)
	}

	author := entity.Author{}

	err = r.Pool.QueryRow(ctx, sql, args...).Scan(
		&author.ID,
		&author.Username,
		&author.Email,
		&author.PasswordHash,
		&author.Salt,
	)
	if err != nil {
		if err == pgx.ErrNoRows {
			return entity.Author{}, fmt.Errorf("AuthorRepo - FindByID - r.Pool.QueryRow. error: %v", apperrors.ErrAuthorNotFound)
		}

		return entity.Author{}, fmt.Errorf("AuthorRepo - FindByID - r.Pool.QueryRow. error: %v", err)
	}

	return author, nil
}

func (r AuthorPostgresRepo) FindByEmail(ctx context.Context, email string) (entity.Author, error) {
	sql, args, err := r.Builder.Select("id, username, email, password_hash, salt").
		From(_athTable).
		Where(squirrel.Eq{"email": email}).
		ToSql()

	if err != nil {
		return entity.Author{}, fmt.Errorf("AuthorRepo - FindByID - r.Builder.Select. error: %v", err)
	}

	author := entity.Author{}

	err = r.Pool.QueryRow(ctx, sql, args...).Scan(
		&author.ID,
		&author.Username,
		&author.Email,
		&author.PasswordHash,
		&author.Salt,
	)
	if err != nil {
		if err == pgx.ErrNoRows {
			return entity.Author{}, fmt.Errorf("AuthorRepo - FindByID - r.Pool.QueryRow. error: %v", apperrors.ErrAuthorNotFound)
		}

		return entity.Author{}, fmt.Errorf("AuthorRepo - FindByID - r.Pool.QueryRow. error: %v", err)
	}

	return author, nil
}

func (r AuthorPostgresRepo) Delete(ctx context.Context, id string) error {
	sql, args, err := r.Builder.
		Delete("").
		From(_athTable).
		Where(squirrel.Eq{"id": id}).
		ToSql()

	if err != nil {
		return fmt.Errorf("AuthorRepo - Delete - r.Builder.Delete. error: %v", err)
	}

	ct, err := r.Pool.Exec(ctx, sql, args...)
	if err != nil {
		return fmt.Errorf("AuthorRepo - Delete - r.Pool.Exec. error: %v", err)
	}

	if !ct.Delete() {
		return fmt.Errorf("AuthorRepo - Delete - r.Pool.Exec. error: %v", apperrors.ErrAuthorNotFound)
	}

	return nil
}
