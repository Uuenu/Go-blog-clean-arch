package v1

import (
	"errors"
	"fmt"
	"go-blog-ca/config"
	"go-blog-ca/internal/usecases"
	"go-blog-ca/pkg/apperrors"
	"go-blog-ca/pkg/logging"
	"net/http"

	"github.com/gin-gonic/gin"
)

type authRoutes struct {
	auth usecases.Auth
	l    logging.Logger
	cfg  *config.Config
}

func newAuthRoutes(handler *gin.RouterGroup, a usecases.Auth, l logging.Logger, c *config.Config) {
	r := authRoutes{
		auth: a,
		l:    l,
		cfg:  c,
	}

	h := handler.Group("/auth")
	{
		h.GET("/logout", r.logout) // sessionMiddleware
		h.POST("/singin", r.signin)
	}

}

func (r *authRoutes) logout(c *gin.Context) {

	sid, err := sessionID(c)
	if err != nil {
		r.l.Error(fmt.Errorf("authRoutes - logout - sessionID. error: %v", err))
		c.AbortWithStatus(http.StatusInternalServerError)
	}

	err2 := r.auth.Logout(c.Request.Context(), sid)
	if err2 != nil {
		r.l.Error(fmt.Errorf("authRoutes - logout - r.auth.Logout. error: %v", err2))
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.SetCookie(
		r.cfg.Session.CookieKey,
		"",
		-1,
		apiPath,
		r.cfg.Session.CookieDomain,
		r.cfg.Session.CookieSecure,
		r.cfg.Session.CookieHTTPOnly,
	)

	c.Status(http.StatusNoContent)

}

type doLoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (r *authRoutes) signin(c *gin.Context) {
	var logReq doLoginRequest
	if err := c.ShouldBindJSON(&logReq); err != nil {
		r.l.Error(fmt.Errorf("authRoutes - signin - c.ShouldBindJSON. error: %v", err))
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	sess, err := r.auth.EmailLogin(c.Request.Context(), logReq.Email, logReq.Password)
	if err != nil {
		r.l.Error(fmt.Errorf("authRoutes - signin - r.auth.EmailLogin. error: %v", err))

		if errors.Is(err, apperrors.ErrAuthorIncorrectPassword) ||
			errors.Is(err, apperrors.ErrAuthorNotFound) {
			//abortWithError(c, http.StatusUnauthorized, apperrors.ErrAuthorIncorrectEmailOrPassword)
			return
		}
	}

	c.SetCookie(
		r.cfg.Session.CookieKey,
		sess.ID,
		sess.TTL,
		apiPath,
		r.cfg.Session.CookieDomain,
		r.cfg.Session.CookieSecure,
		r.cfg.Session.CookieHTTPOnly,
	)

	c.Status(http.StatusOK)
}
