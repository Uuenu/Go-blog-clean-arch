package v1

import (
	"fmt"
	"go-blog-ca/internal/usecases"
	"go-blog-ca/pkg/logging"
	"net/http"

	"github.com/gin-gonic/gin"
)

type authorRoutes struct {
	auth usecases.Author
	l    logging.Logger
}

func newAuthorRoutes(handler *gin.RouterGroup, auth usecases.Author, l logging.Logger) {
	r := &authorRoutes{
		auth: auth,
		l:    l,
	}

	h := handler.Group("/author")
	{
		authenticated := h.Group("/") // add sessionMiddleware()
		{
			authenticated.PUT("/:author_id")
			authenticated.DELETE("/:autrhor_id")
			authenticated.GET("/:author_id", r.ArticleByID) // get by id
			authenticated.GET("")                           //get all
			//authenticated.

		}

		h.POST("signup")                    // create
		h.POST("signin")                    // auth
		h.GET("/:author_id", r.ArticleByID) // get by id
		h.GET("")                           //get all

	}

}

func (r *authorRoutes) ArticleByID(c *gin.Context) {
	aid := c.Param("author_id")

	acc, err := r.auth.GetByID(c.Request.Context(), aid)
	if err != nil {
		r.l.Error(fmt.Errorf("http - v1 - auth - get: %w", err))
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, acc)

}