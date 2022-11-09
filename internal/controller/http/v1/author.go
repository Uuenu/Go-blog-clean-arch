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
	s    usecases.Session
	l    logging.Logger
}

func newAuthorRoutes(handler *gin.RouterGroup, auth usecases.Author, s usecases.Session, l logging.Logger) {
	r := &authorRoutes{
		auth: auth,
		s:    s,
		l:    l,
	}

	h := handler.Group("/author")
	{
		authenticated := h.Group("/", sessionMiddleware(l, s)) // add sessionMiddleware()
		{
			authenticated.PUT("/:id")
			authenticated.DELETE("/:id")

		}

		h.GET("/:id", r.ArticleByID) // get by id
		h.GET("")                    //get all

	}

}

func (r *authorRoutes) ArticleByID(c *gin.Context) {
	aid := c.Param("id")

	acc, err := r.auth.GetByID(c.Request.Context(), aid)
	if err != nil {
		r.l.Error(fmt.Errorf("http - v1 - auth - get: %w", err))
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, acc)

}
