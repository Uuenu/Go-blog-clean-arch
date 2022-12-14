package usecases

import (
	"context"
	"fmt"
	"go-blog-ca/config"
	"go-blog-ca/internal/domain/entity"
	"go-blog-ca/pkg/apperrors"
	"go-blog-ca/pkg/logging"
	"time"
)

type SessionUseCase struct {
	repo   SessionRepo
	cfg    *config.Config
	logger logging.Logger
}

func NewSessionUseCase(r SessionRepo, cfg *config.Config, l logging.Logger) *SessionUseCase {
	return &SessionUseCase{
		cfg:    cfg,
		repo:   r,
		logger: l,
	}
}

func (uc *SessionUseCase) Create(ctx context.Context, aid string) (entity.Session, error) {
	sess, err := entity.NewSession(aid, "", uc.cfg.Session.TTL)
	if err != nil {
		return entity.Session{}, fmt.Errorf("SessionUseCase - Create - entity.NewSession: %w", err)
	}

	repoCtx, cancel := context.WithTimeout(ctx, time.Second*1)
	defer cancel()

	sid, err := uc.repo.Create(repoCtx, sess)

	if err != nil {
		return entity.Session{}, fmt.Errorf("SessionUseCase - Create - uc.repo.Create: %w", err)
	}

	sess.ID = sid

	return sess, nil
}

func (uc *SessionUseCase) GetByID(ctx context.Context, sid string) (entity.Session, error) {

	sess, err := uc.repo.FindByID(ctx, sid)

	if err != nil {
		return entity.Session{}, fmt.Errorf("SessionUseCase  - GetByID - uc.repo.FindByID: %w", err)
	}

	return sess, nil
}

func (uc *SessionUseCase) GetAll(ctx context.Context, aid string) ([]entity.Session, error) {
	sessions, err := uc.repo.FindAll(ctx, aid)
	if err != nil {
		return nil, fmt.Errorf("SessionUseCase - GetAll - uc.repo.GetAll: %w", err)
	}

	return sessions, nil
}

func (uc *SessionUseCase) Terminate(ctx context.Context, sid, currSid string) error {
	if sid == currSid {
		return fmt.Errorf("SessionUseCase - Terminate: %w", apperrors.ErrSessionNotTerminated)
	}

	if err := uc.repo.Delete(ctx, sid); err != nil {
		return fmt.Errorf("SessionUseCase - Terminate - s.repo.Delete: %w", err)
	}

	return nil
}

func (uc *SessionUseCase) TerminateAll(ctx context.Context, aid, currSid string) error {
	if err := uc.repo.DeleteAll(ctx, aid, currSid); err != nil {
		return fmt.Errorf("SessionUseCase - TerminateAll - s.repo.DeleteAll. error: %v", apperrors.ErrSessionNotTerminated)
	}
	return nil
}
