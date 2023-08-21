package delivery

import (
	"net/http"
	"task/internal/repository"
	"time"

	"github.com/gin-gonic/gin"
)

func (h *Handler) SetCookie(c *gin.Context, tokens *repository.Jwt) {
	accessExpire := time.Now().Add(time.Minute * 100)
	refreshExpire := time.Now().AddDate(0, 0, 10)
	http.SetCookie(c.Writer, &http.Cookie{
		Name:     "access_token",
		Value:    tokens.AccessToken,
		HttpOnly: true,
		MaxAge:   10 * 24 * 60 * 60,
		Path:     "/",
		Expires:  accessExpire,
	})

	http.SetCookie(c.Writer, &http.Cookie{
		Name:     "refresh_token",
		Value:    tokens.RefreshToken,
		HttpOnly: true,
		MaxAge:   10 * 24 * 60 * 60,
		Path:     "/refresh",
		Expires:  refreshExpire,
	})
}
