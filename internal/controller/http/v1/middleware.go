package v1

import (
	"go-blog-ca/pkg/apperrors"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func authorID(c *gin.Context) (string, error) {
	aid := c.GetString("aid")

	_, err := uuid.Parse(aid)
	if err != nil {
		return "", apperrors.ErrAuthorContextNotFound
	}

	return aid, nil
}

func sessionID(c *gin.Context) (string, error) {
	sid := c.GetString("sid")

	if sid == "" {
		return "", apperrors.ErrSessionContextNotFound
	}
	return sid, nil
}
