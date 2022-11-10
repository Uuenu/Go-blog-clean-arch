package entity

import (
	"fmt"
	"time"

	"go-blog-ca/pkg/apperrors"
	"go-blog-ca/pkg/utils"
)

type Session struct {
	ID       string    `json:"id" bson:"_id"`
	AuthorID string    `json:"author_id" bson:"author_id"`
	IP       string    `json:"ip" bson:"ip"`
	TTL      int       `json:"ttl" bson:"ttl"`
	CreateAt time.Time `json:"createAt" bson:"create_at"`
}

func NewSession(aid, ip string, ttl time.Duration) (Session, error) {
	//id := utils.GenerateId() // replace on unique string
	id, err := utils.UniqueString(32)
	if err != nil {
		return Session{}, fmt.Errorf("Session - NewSession() - utils.UniqueString(). error: %v", apperrors.ErrSessionNotCreated)
	}

	now := time.Now()

	return Session{
		ID:       id,
		AuthorID: aid,
		IP:       ip,
		TTL:      int(ttl.Seconds()), // Time To Live
		CreateAt: now,
	}, nil
}
