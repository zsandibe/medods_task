package delivery

import (
	"log"
	"task/internal/repository"

	"github.com/gin-gonic/gin"
)

func (h *Handler) Refresh(c *gin.Context) {
	// fmt.Println("OK")
	userGuid := c.Param("guid")
	// fmt.Println(userGuid)
	refreshToken, err := c.Request.Cookie("refresh-token")
	// fmt.Println(refreshToken)
	// fmt.Println(err)
	if err != nil {
		log.Printf("Can`t find refresh token : %v\n", err)
		c.AbortWithStatusJSON(401, gin.H{"error": "Invalid api token"})
		return
	}

	accessToken, err := c.Request.Cookie("access-token")
	if err != nil {
		log.Printf("Can`t find access token : %v\n", err)
		c.AbortWithStatusJSON(401, gin.H{"error": "Invalid api token"})
		return
	}

	oldTokens := &repository.Jwt{
		UserGUID:     userGuid,
		RefreshToken: refreshToken.Value,
		AccessToken:  accessToken.Value,
	}

	tokens, err := h.service.UpdateToken(oldTokens)
	if err != nil {
		log.Printf("Can`t find refresh token in DB: %v\n", err)
		c.AbortWithStatusJSON(401, gin.H{"error": "error"})
		return
	}
	h.SetCookie(c, tokens)
}
