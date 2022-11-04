package v1

import (
	"fmt"
	"go-blog-ca/internal/domain/entity"
	"go-blog-ca/internal/usecases"
	"go-blog-ca/pkg/logging"
	"net/http"

	"github.com/gin-gonic/gin"
)

type articleRoutes struct {
	artcl usecases.Article
	s     usecases.Session
	l     logging.Logger
}

func newArticleRoutes(handler *gin.RouterGroup, artcl usecases.Article, s usecases.Session, l logging.Logger) {
	r := articleRoutes{artcl, s, l}

	h := handler.Group("/article")
	{
		h.GET("/:id", r.GetByID) // get by id
		h.POST("")               // create
		h.GET("")                //get all
		h.PUT("/:id")            // update article
		h.DELETE("/:id")         //delete by id
	}
}

func (r *articleRoutes) GetByID(c *gin.Context) {
	aid := c.Param("id")

	acc, err := r.artcl.GetByID(c.Request.Context(), aid)
	if err != nil {
		r.l.Error(fmt.Errorf("http - v1 - article - get: %w", err))
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, acc)

}

type doCreateRequest struct {
	AuthorID string `json:"author_id" binding:"required"`
	Header   string `json:"header" binding:"required"`
	Text     string `json:"text" binding:"required"`
}

func (r *articleRoutes) Create(c *gin.Context) {
	var request doCreateRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		// TODO logger and error response

		return
	}

	article, err := r.artcl.Create(
		c.Request.Context(),
		entity.Article{
			AuthorID: request.AuthorID,
			Header:   request.Header,
			Text:     request.Text,
		},
	)

	if err != nil {
		// TODO logger and error response

		return
	}

	c.JSON(http.StatusOK, article)

}

func (r *articleRoutes) GetAll(c *gin.Context) {
	arcles, err := r.artcl.GetAll(c.Request.Context())

	if err != nil {
		// TODO logger and error respose
		return
	}

	c.JSON(http.StatusOK, arcles)
}

func (r *articleRoutes) Update(c *gin.Context) {
	panic("implement me")
}

type doDeleteRequest struct {
	//AuthorID  string `json:"article_id"` // aid from article == aid from session
	ArticleID string `json:"id"`
	SessionID string `json:"session_id"`
}

func (r *articleRoutes) Delete(c *gin.Context) {
	var request doDeleteRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		// TODO logger and error response
		return
	}

	// article, err := r.artcl.GetByID(c.Request.Context(), request.ArticleID)
	// if err != nil {
	// 	//TODO error
	// 	return
	// }

}
