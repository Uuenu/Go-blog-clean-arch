package v1

import (
	"go-blog-ca/internal/usecases"
	"go-blog-ca/pkg/logging"
)

type sessionRoutes struct {
	s usecases.Session
	l logging.Logger
}
