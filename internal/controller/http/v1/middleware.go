package v1

import (
	"fmt"
	"go-blog-ca/internal/usecases"
	"go-blog-ca/pkg/apperrors"
	"go-blog-ca/pkg/logging"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// TODO fix session

func sessionMiddleware(l logging.Logger, s usecases.Session) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		sid, err := ctx.Cookie("id")
		l.Infof("SessionID From Middleware: %s", sid)
		if err != nil {
			l.Error(fmt.Errorf("http - v1 - middleware sessionMiddleware - c.Cookie. error: %v", err))
			ctx.AbortWithStatus(http.StatusUnauthorized)
		}

		sess, err := s.GetByID(ctx, sid)
		if err != nil {
			l.Error(fmt.Errorf("http - v1 - middleware - sessionMiddleware - s.GetByID. error: %v", err))
			ctx.AbortWithStatus(http.StatusUnauthorized)
		}
		l.Infof("Session: %s", sess)

		ctx.Set("sid", sess.ID)
		ctx.Set("aid", sess.AuthorID)
		ctx.Next()
	}
}

func authorID(c *gin.Context) (string, error) {
	aid := c.GetString("aid")

	_, err := uuid.Parse(aid)
	if err != nil {
		return "", apperrors.ErrAuthorContextNotFound
	}

	return aid, nil
}

func sessionID(c *gin.Context) (string, error) {
	sid, err := c.Cookie("id")
	if err != nil {
		return "", apperrors.ErrSessionContextNotFound
	}

	if sid == "" {
		return "", apperrors.ErrSessionContextNotFound
	}
	return sid, nil
}
