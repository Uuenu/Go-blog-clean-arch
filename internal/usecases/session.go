package usecases

import (
	"context"
	"go-blog-ca/internal/domain/entity"
)

type SessionUseCase struct {
	repo SessionRepo
}

func NewSessionUseCase(r SessionRepo) *SessionUseCase {
	return &SessionUseCase{
		repo: r,
	}
}

func (uc *SessionUseCase) Create(ctx context.Context, aid string) (entity.Session, error) {
	panic("implement me")
}

func (uc *SessionUseCase) GetByID(ctx context.Context, sid string) (entity.Session, error) {
	panic("implement me")
}

func (uc *SessionUseCase) GetAll(ctx context.Context) []entity.Session {
	panic("implement me")
}

func (uc *SessionUseCase) Terminate(ctx context.Context, sid, currSid string) error {
	panic("implement me")
}

func (uc *SessionUseCase) TerminateAll(ctx context.Context, aid, sid string) error {
	panic("implement me")
}
