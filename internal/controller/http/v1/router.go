package v1

import (
	"go-blog-ca/config"
	"go-blog-ca/internal/usecases"
	"go-blog-ca/pkg/logging"

	"github.com/gin-gonic/gin"
)

const apiPath = "/v1"

// NewRouter -.
// Swagger spec:
// @title       Blog based on Clean Architecture
// @description Using a blog service as an example
// @version     1.0
// @host        localhost:8080
// @BasePath    /v1
func NewRouter(handler *gin.Engine, l logging.Logger, author usecases.Author,
	artcl usecases.Article, auth usecases.Auth, session usecases.Session, c *config.Config) {

	handler.Use(gin.Logger())
	handler.Use(gin.Recovery())

	h := handler.Group("/v1")
	{
		newArticleRoutes(h, artcl, session, l)
		newAuthorRoutes(h, author, auth, session, l)
		newSessionRoutes(h, session, author, l, c)
		newAuthRoutes(h, auth, session, l, c)
	}
}
