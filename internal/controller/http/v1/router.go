package v1

import (
	"go-blog-ca/internal/usecases"
	"go-blog-ca/pkg/logging"

	"github.com/gin-gonic/gin"
)

func NewRouter(handler *gin.Engine, l logging.Logger, auth usecases.Author, artcl usecases.Article) {
	handler.Use(gin.Logger())
	handler.Use(gin.Recovery())

	h := handler.Group("/v1")
	{
		newArticleRoutes(h, artcl, l)
		newAuthorRoutes(h, auth, l)
		//newSessionRoutes(h, session, l)
	}
}
