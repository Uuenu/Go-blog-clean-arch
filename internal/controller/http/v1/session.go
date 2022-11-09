package v1

import (
	"fmt"
	"go-blog-ca/config"
	"go-blog-ca/internal/usecases"
	"go-blog-ca/pkg/logging"
	"net/http"

	"github.com/gin-gonic/gin"
)

type sessionRoutes struct {
	s   usecases.Session
	a   usecases.Author
	l   logging.Logger
	cfg *config.Config
}

func newSessionRoutes(handler *gin.RouterGroup, s usecases.Session, a usecases.Author, l logging.Logger, c *config.Config) {
	r := sessionRoutes{s, a, l, c}

	h := handler.Group("/session") // sessionMiddleware()
	{

		h.GET("/", r.SessionByID)
		h.DELETE("/:sessionID", r.Terminate)
		h.DELETE("/", r.TerminateAll)
	}

}

func (r *sessionRoutes) SessionByID(c *gin.Context) { // change on get
	sid, err := sessionID(c)
	if err != nil {
		r.l.Error(fmt.Errorf("http - v1 - sessionID. error: %v", err))
		c.AbortWithStatus(http.StatusInternalServerError) //TODO check status
		return
	}

	// check sid is exist for getbyid for author

	sess, err := r.s.GetByID(c.Request.Context(), sid)
	if err != nil {
		r.l.Error(fmt.Errorf("http - v1 - SessionByID. error: %v", err))
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, sess)
}

func (r *sessionRoutes) Terminate(c *gin.Context) {
	currID, err := sessionID(c)
	if err != nil {
		r.l.Error(fmt.Errorf("http - v1 - Terminate - sessionID. error: %v", err))
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	if err != nil {
		r.l.Error(fmt.Errorf("http - v1 - terminate - Cookie: %w", err))
		c.AbortWithStatus(http.StatusInternalServerError) //Change
		return
	}

	err = r.s.Terminate(c.Request.Context(), c.Param("sessionID"), currID)
	if err != nil {
		r.l.Error(fmt.Errorf("http - v1 - terminate. error: %v", err))
		c.AbortWithStatus(http.StatusInternalServerError) // change
		return
	}

	c.JSON(http.StatusOK, nil)
}

func (r *sessionRoutes) TerminateAll(c *gin.Context) {
	currID, err := sessionID(c)
	if err != nil {
		r.l.Error(fmt.Errorf("http - v1 - terminate - sessionID. error: %v", err))
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	aid, err := authorID(c)
	if err != nil {
		r.l.Error(fmt.Errorf("http - v1 - terminate - authorID. error: %v", err))
		c.AbortWithStatus(http.StatusUnauthorized)
	}

	err = r.s.TerminateAll(c.Request.Context(), aid, currID)
	if err != nil {
		r.l.Error(fmt.Errorf("http - v1 - terminateAll. error: %v", err))
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, nil)
}
