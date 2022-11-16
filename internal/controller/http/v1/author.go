package v1

import (
	"fmt"
	"go-blog-ca/internal/domain/entity"
	"go-blog-ca/internal/usecases"
	"go-blog-ca/pkg/logging"
	"net/http"

	"github.com/gin-gonic/gin"
)

type authorRoutes struct {
	ath usecases.Author
	s   usecases.Session
	l   logging.Logger
}

func newAuthorRoutes(handler *gin.RouterGroup, ath usecases.Author, s usecases.Session, l logging.Logger) {
	r := &authorRoutes{
		ath: ath,
		s:   s,
		l:   l,
	}

	h := handler.Group("/author")
	{
		authenticated := h.Group("", sessionMiddleware(l, s)) // add sessionMiddleware()
		{
			authenticated.DELETE("/:id")

		}
		h.POST("/signup", r.Singup)
		h.GET("/:id", r.ArticleByID) // get by id
		h.GET("")                    //get all

	}

}

type doSignupRequest struct {
	Username string `json:"username,omitempty"`
	Email    string `json:"email,omitempty"`
	Password string `json:"password,omitempty"`
}

func (r *authorRoutes) Singup(c *gin.Context) {

	var signup doSignupRequest
	if err := c.ShouldBindJSON(&signup); err != nil {
		r.l.Error(fmt.Errorf("http - v1 - ath - c.ShouldBindJSON: %v", err))
		c.AbortWithStatusJSON(http.StatusInternalServerError, nil)
		return
	}

	r.l.Infoln(signup)

	newAth := entity.Author{
		ID:           "",
		Username:     signup.Username,
		Email:        signup.Email,
		Password:     signup.Password,
		PasswordHash: "",
		Salt:         []byte{},
	}

	r.l.Infoln(newAth)

	if err := newAth.GeneratePasswordHash(); err != nil {
		r.l.Error(fmt.Errorf("http - v1 - author - GeneratePasswordHash: %v", err))
		c.AbortWithStatusJSON(http.StatusInternalServerError, nil)
		return
	}

	id, err := r.ath.Create(c.Request.Context(), newAth)
	if err != nil {
		r.l.Error(fmt.Errorf("http - v1 - author - r.ath.Create: %v", err))
		c.AbortWithStatusJSON(http.StatusInternalServerError, nil)
		return
	}

	r.l.Infof("New author with id: %s", id)

	c.JSON(http.StatusOK, nil)

}

func (r *authorRoutes) ArticleByID(c *gin.Context) {
	aid := c.Param("id")

	acc, err := r.ath.GetByID(c.Request.Context(), aid)
	if err != nil {
		r.l.Error(fmt.Errorf("http - v1 - ath - get: %w", err))
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, acc)

}
