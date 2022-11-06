package v1

import (
	"fmt"
	"go-blog-ca/internal/usecases"
	"go-blog-ca/pkg/logging"
	"net/http"

	"github.com/gin-gonic/gin"
)

type sessionRoutes struct {
	s usecases.Session
	a usecases.Author
	l logging.Logger
}

func newSessionRoutes(handler *gin.RouterGroup, s usecases.Session, a usecases.Author, l logging.Logger) {
	r := sessionRoutes{s, a, l}

	h := handler.Group("session") // sessionMiddleware()
	{
		h.GET("/:id")
		h.DELETE("/:id", r.Terminate)
		h.DELETE("", r.TerminateAll)
	}

}

func (r *sessionRoutes) SessionByID(c *gin.Context) { // change on get
	sid := c.Param("id")

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
	sid := c.Param("id")
	c.GetString("id")
	currID, err := c.Cookie("id") // current session id
	if err != nil {
		r.l.Error(fmt.Errorf("http - v1 - terminate - Cookie: %w", err))
		c.AbortWithStatus(http.StatusInternalServerError) //Change
		return
	}

	err = r.s.Terminate(c.Request.Context(), sid, currID)
	if err != nil {
		r.l.Error(fmt.Errorf("http - v1 - terminate. error: %v", err))
		c.AbortWithStatus(http.StatusInternalServerError) // change
		return
	}

	c.JSON(http.StatusOK, nil)
}

func (r *sessionRoutes) TerminateAll(c *gin.Context) {
	aid := c.Param("id")          // author id
	currID, err := c.Cookie("id") // current session id
	if err != nil {
		r.l.Error(fmt.Errorf("http - v1 - terminate - Cookie. error: %v", err))
		c.AbortWithStatus(http.StatusInternalServerError) // change
		return
	}

	err = r.s.TerminateAll(c.Request.Context(), aid, currID)
	if err != nil {
		r.l.Error(fmt.Errorf("http - v1 - terminateAll. error: %v", err))
		return
	}

	c.JSON(http.StatusOK, nil)
}
