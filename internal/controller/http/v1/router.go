package v1

import (
	"go-blog-ca/internal/config"
	"go-blog-ca/internal/usecases"
	"go-blog-ca/pkg/logging"

	"github.com/gin-gonic/gin"
)

const apiPath = "/v1"

func NewRouter(handler *gin.Engine, l logging.Logger, author usecases.Author,
	artcl usecases.Article, auth usecases.Auth, session usecases.Session, c *config.Config) {

	handler.Use(gin.Logger())
	handler.Use(gin.Recovery())

	h := handler.Group("/v1")
	{
		newArticleRoutes(h, artcl, session, l)
		newAuthorRoutes(h, author, l)
		newSessionRoutes(h, session, author, l, c)
		newAuthRoutes(h, auth, l, c)
	}
}
