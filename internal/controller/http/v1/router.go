package v1

import (
	"go-blog-ca/internal/usecases"
	"go-blog-ca/pkg/logging"

	"github.com/gin-gonic/gin"
)

func NewRouter(handler *gin.Engine, l logging.Logger, author usecases.Author,
	artcl usecases.Article, auth usecases.AuthUseCase, session usecases.Session) {

	handler.Use(gin.Logger())
	handler.Use(gin.Recovery())

	h := handler.Group("/v1")
	{
		newArticleRoutes(h, artcl, session, l)
		newAuthorRoutes(h, author, l)
		newSessionRoutes(h, session, author, l)

	}
}
