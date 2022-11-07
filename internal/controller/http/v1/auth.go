package v1

import (
	"fmt"
	"go-blog-ca/internal/usecases"
	"go-blog-ca/pkg/logging"
	"net/http"

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

	h := handler.Group("") // sessionMiddleware
	{
		//logout
		h.GET("/logout", r.logout)
	}
	h.POST("/singin")
	h.POST("/signout")
}

func (r *authRoutes) logout(c *gin.Context) {
	//get sid
	sid := c.GetString("sid")
	err := r.auth.Logout(c.Request.Context(), sid)
	if err != nil {
		r.l.Error(fmt.Errorf("authRoutes - logout - r.auth.Logout. error: %v", err))
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, nil)

}

type doLoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
