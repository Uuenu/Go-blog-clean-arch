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
	ath  usecases.Author
	auth usecases.Auth
	s    usecases.Session
	l    logging.Logger
}

func newAuthorRoutes(handler *gin.RouterGroup, ath usecases.Author, auth usecases.Auth, s usecases.Session, l logging.Logger) {
	r := &authorRoutes{
		ath:  ath,
		auth: auth,
		s:    s,
		l:    l,
	}

	h := handler.Group("/author")
	{
		authenticated := h.Group("", sessionMiddleware(l, s)) // add sessionMiddleware()
		{
			authenticated.DELETE("/:id", r.Archive)

		}
		h.GET("/email", r.AuthorByEmail)
		h.POST("/signup", r.Singup)
		h.GET("/:id", r.AuthorByID) // get by id
		h.GET("", r.Authors)        //get all

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

type doResponseAuthor struct {
	ID       string `bson:"id,omitempty"`
	Username string `bson:"username,omitempty"`
	Email    string `bson:"email,omitempty"`
}

func (r *authorRoutes) AuthorByID(c *gin.Context) {
	aid := c.Param("id")

	acc, err := r.ath.GetByID(c.Request.Context(), aid)
	if err != nil {
		r.l.Error(fmt.Errorf("http - v1 - ath - get: %w", err))
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	response := doResponseAuthor{
		ID:       acc.ID,
		Username: acc.Username,
		Email:    acc.Email,
	}

	c.JSON(http.StatusOK, response)

}

func (r *authorRoutes) AuthorByEmail(c *gin.Context) {
	email := "codyvangoth@gmail.com"

	author, err := r.ath.GetByEmail(c.Request.Context(), email)
	if err != nil {
		r.l.Error(fmt.Errorf("http - v1 - ath - get by email: %v", err))
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	response := doResponseAuthor{
		ID:       author.ID,
		Username: author.Username,
		Email:    author.Email,
	}
	c.JSON(http.StatusOK, response)
}

func (r *authorRoutes) Authors(c *gin.Context) {
	authors, err := r.ath.GetAll(c.Request.Context())
	if err != nil {
		r.l.Error(fmt.Errorf("http - v1 - ath - get: %w", err))
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	var response []doResponseAuthor
	for _, ath := range authors {
		response = append(response, doResponseAuthor{
			ID:       ath.ID,
			Username: ath.Username,
			Email:    ath.Email,
		})
	}

	c.JSON(http.StatusOK, response)
}

func (r *authorRoutes) Archive(c *gin.Context) {
	sid, err := sessionID(c)
	if err != nil {
		r.l.Error(fmt.Errorf("http - v1 - ath - get: %v", err))
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	sess, err := r.s.GetByID(c.Request.Context(), sid)
	if err != nil {
		r.l.Error(fmt.Errorf("http - v1 - ath - get: %v", err))
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	aid := c.Param("id")
	r.l.Infof("Session.AuthorID: %v, aid: %v", sess.AuthorID, aid)

	if sess.AuthorID != aid {
		r.l.Infof("Session.AuthorID: %v, aid: %v", sess.AuthorID, aid)
		c.AbortWithStatus(http.StatusForbidden)
		return
	}

	if err := r.ath.Delete(c.Request.Context(), aid, sid); err != nil {
		r.l.Error(fmt.Errorf("http - v1 - ath - get: %w", err))
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	c.JSON(http.StatusOK, nil)
}
