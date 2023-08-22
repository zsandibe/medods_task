package delivery

import (
	"log"

	"github.com/gin-gonic/gin"
)

func (h *Handler) GenerateToken(c *gin.Context) {
	userGuid := c.Param("guid")
	token, err := h.service.CreateToken(userGuid)
	if err != nil {
		log.Printf("Error don't generate pair of tokens: %v\n", err)
		c.AbortWithStatusJSON(401, nil)
		return
	}
	// fmt.Println(token)
	h.SetCookie(c, token)
}
