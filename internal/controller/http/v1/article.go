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
		h.GET("/id", r.getByID) // get by id
		h.GET("/all", r.all)    //get all

		authorized := h.Group("", sessionMiddleware(l, s))
		{
			authorized.GET("/hello", r.hello)
			authorized.POST("/create", r.create)
			authorized.PUT("", r.update)    //update article
			authorized.DELETE("", r.delete) //delete by id
		}

	}
}

func (r *articleRoutes) hello(c *gin.Context) {
	c.JSON(http.StatusOK, "Hello")
}

func (r *articleRoutes) getByID(c *gin.Context) {
	articleID := c.Param("id")

	acc, err := r.artcl.GetByID(c.Request.Context(), articleID)
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

func (r *articleRoutes) create(c *gin.Context) {
	var request doCreateRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		r.l.Error(fmt.Errorf("articleRoutes - Create - ShouldBindJSON. error: %v", err))
		c.AbortWithStatus(http.StatusBadRequest) // TODO check status
		return
	}

	aid, err := authorID(c)
	if err != nil {
		r.l.Error(fmt.Errorf("articleRoutes - create - sessionID. error: %v", err))
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	// request.AuthorID = aid

	if aid != request.AuthorID {
		r.l.Error(fmt.Errorf("articleRoutes - create - aid != request.AuthorID. "))
		c.AbortWithStatus(http.StatusForbidden)
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
		r.l.Error(fmt.Errorf("aritcleRoutes - Create - r.artcl.Create. error: %v", err))
		c.AbortWithStatus(http.StatusBadRequest) // TODO check status
		return
	}

	c.JSON(http.StatusCreated, article) // TODO check status

}

func (r *articleRoutes) all(c *gin.Context) {
	arcles, err := r.artcl.GetAll(c.Request.Context())

	if err != nil {
		r.l.Error(fmt.Errorf("articleRoutes - GetAll - r.artcl.GetAll. error: %v", err))
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, arcles)
}

func (r *articleRoutes) update(c *gin.Context) {
	panic("implement me")
}

type doDeleteRequest struct {
	//AuthorID  string `json:"article_id"` // aid from article == aid from session
	ArticleID string `json:"articleID"`
	SessionID string `json:"sid"`
	AuthorID  string `json:"aid"`
}

func (r *articleRoutes) delete(c *gin.Context) {
	var request doDeleteRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		// TODO logger and error response
		r.l.Error(fmt.Errorf("articleRoutes - Delete - ShouldBindJSON. error: %v", err))
		c.AbortWithStatus(http.StatusBadRequest) // TODO check status
		return
	}

	aid, err := authorID(c)
	if err != nil {
		r.l.Error(fmt.Errorf("articleRoutes - Delete - authorID. error: %v", err))
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	article, err := r.artcl.GetByID(c.Request.Context(), request.ArticleID)
	if err != nil {
		r.l.Error(fmt.Errorf("articleRoutes - Delete - r.artcl.GetByID. error: %v", err))
		c.AbortWithStatus(http.StatusInternalServerError) // TODO check status
		return
	}

	session, err := r.s.GetByID(c.Request.Context(), request.SessionID)
	if err != nil {
		r.l.Error(fmt.Errorf("articleRoutes - Delete - r.s.GetByID. error: %v", err))
		c.AbortWithStatus(http.StatusInternalServerError) // TODO check status
		return
	}

	if session.AuthorID != article.AuthorID {
		c.AbortWithStatus(http.StatusForbidden)
		return
	}

	if err := r.artcl.Delete(c.Request.Context(), request.ArticleID, aid); err != nil {
		r.l.Error(fmt.Errorf("articleRoutes - Delete - r.artcl.Delete. error: %v", err))
		c.AbortWithStatus(http.StatusInternalServerError) // TODO check status
		return
	}

	c.JSON(http.StatusOK, nil)
}
