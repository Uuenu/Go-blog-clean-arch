package repo

import (
	"context"
	"fmt"
	"go-blog-ca/internal/domain/entity"
	"go-blog-ca/pkg/postgres"
)

const _defaultEntityCap = 64

type AuthorPostgresRepo struct {
	*postgres.Postgres
}

func New(pg *postgres.Postgres) *AuthorPostgresRepo {
	return &AuthorPostgresRepo{pg}
}

func (r AuthorPostgresRepo) GetAll(ctx context.Context) ([]entity.Author, error) {
	sql, _, err := r.Builder.
		Select("id, username, email, password_hash, salt").
		From("author").
		ToSql()
	if err != nil {
		return nil, fmt.Errorf("AuthorRepo - GetAll - r.Builder. error: %v", err)
	}

	rows, err := r.Pool.Query(ctx, sql)
	if err != nil {
		return nil, fmt.Errorf("AuthorRepo - GetAll - r.Pool.Query, error: %v", err)
	}
	defer rows.Close()

	authors := make([]entity.Author, 0, _defaultEntityCap)

	for rows.Next() {
		ath := entity.Author{}

		err := rows.Scan(ath.ID, ath.Username, ath.Email, ath.PasswordHash, ath.Salt)
		if err != nil {
			return nil, fmt.Errorf("AuthorRepo - GetAll - rows.Scan. error: %v", err)
		}

		authors = append(authors, ath)
	}

	return authors, nil
}
