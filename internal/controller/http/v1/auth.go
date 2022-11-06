package v1

import (
	"go-blog-ca/internal/usecases"
	"go-blog-ca/pkg/logging"

	"github.com/gin-gonic/gin"
)

type authRoutes struct {
	auth usecases.Auth
	l    logging.Logger
}

func newAuthRoutes(handler *gin.RouterGroup, a usecases.Auth, l logging.Logger) {
	r := authRoutes{
		auth: a,
		l:    l,
	}

	h := handler.Group("/") // sessionMiddleware
	{
		//logout
		h.GET("/logout", r.logout)
	}
	// signup
	// signin
}

func (r *authRoutes) logout(c *gin.Context) {
	//get sid
	r.auth.Logout(c.Request.Context(), "")
}

type doLoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
