package usecases

import (
	"context"
	"fmt"
	"go-blog-ca/internal/domain/entity"
	"go-blog-ca/pkg/apperrors"
	"go-blog-ca/pkg/utils"
)

type AuthUseCase struct {
	Author  Author
	Session Session
}

func NewAuthUseCase(a Author, s Session) *AuthUseCase {
	return &AuthUseCase{
		Author:  a,
		Session: s,
	}
}

func (uc *AuthUseCase) EmailLogin(ctx context.Context, email, password string) (entity.Session, error) {
	a, err := uc.Author.GetByEmail(ctx, email)
	if err != nil {
		return entity.Session{}, fmt.Errorf("AuthUseCase - EmailLogin - GetByEmail: %w", apperrors.ErrAuthorIncorrectEmailOrPassword)
	}

	a.Password = ""

	currPasswordHash, err := utils.PasswordHash(password, a.Salt)
	if err != nil {
		return entity.Session{}, fmt.Errorf("AuthUseCase - EmailLogin - PasswordHash: %v", err)
	}

	if currPasswordHash != a.PasswordHash {
		return entity.Session{}, fmt.Errorf("AuthUseCase - EmailLogin: %v", apperrors.ErrAuthorIncorrectEmailOrPassword)
	}

	sess, err := uc.Session.Create(ctx, a.ID)
	if err != nil {
		return entity.Session{}, fmt.Errorf("AuthUseCase - EmailLogin - uc.Session.Create: %w", err)
	}

	return sess, nil

}

func (uc *AuthUseCase) Logout(ctx context.Context, sid string) error {
	if err := uc.Session.Terminate(ctx, sid, ""); err != nil {
		return fmt.Errorf("AuthUseCase - Logout - uc.Session.Terminate: %w", err)
	}

	return nil
}
