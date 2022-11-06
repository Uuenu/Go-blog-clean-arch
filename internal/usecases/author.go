package usecases

import (
	"context"
	"fmt"
	"go-blog-ca/internal/domain/entity"
)

type AuthorUseCase struct {
	repo    AuthorRepo
	session Session
}

func NewAuthorUseCase(r AuthorRepo, s Session) *AuthorUseCase {
	return &AuthorUseCase{
		repo:    r,
		session: s,
	}
}

func (uc *AuthorUseCase) Create(ctx context.Context, author entity.Author) (string, error) {

	//Generate password hash
	if err := author.GeneratePasswordHash(); err != nil {
		return "", fmt.Errorf("AuthorUseCase - Create - GeneratePasswordHash. error: %v", err)
	}

	aid, err := uc.repo.Create(ctx, author)
	if err != nil {
		return "", fmt.Errorf("AuthorUseCase - Create - us.repo.Create error: %v", err)
	}

	return aid, nil
}

func (uc *AuthorUseCase) GetByID(ctx context.Context, id string) (entity.Author, error) {
	author, err := uc.repo.FindByID(ctx, id)
	if err != nil {
		return entity.Author{}, fmt.Errorf("AuthorUseCase - GetByID - us.repo.FindByID error: %v", err)
	}

	return author, nil
}

func (uc *AuthorUseCase) GetByEmail(ctx context.Context, email string) (entity.Author, error) {
	author, err := uc.repo.FindByEmail(ctx, email)
	if err != nil {
		return entity.Author{}, fmt.Errorf("AuthorUseCase - GetByEmail - us.repo.FindByEmail error: %v", err)
	}

	return author, nil
}

func (uc *AuthorUseCase) GetAll(ctx context.Context) ([]entity.Author, error) {
	authors, err := uc.repo.FindAll(ctx)
	if err != nil {
		return nil, fmt.Errorf("AuthorUseCase - GetAll - us.repo.FindAll error: %v", err)
	}
	return authors, nil
}

func (uc *AuthorUseCase) Delete(ctx context.Context, id string) error {
	
	return nil
}
