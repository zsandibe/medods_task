package delivery

import (
	"task/internal/service"

	"github.com/gin-gonic/gin"
)

type HandlerInterface interface {
	Routes(router *gin.Engine)
}

type Handler struct {
	service service.AuthService
}

func NewHandlers(auth service.AuthService) HandlerInterface {
	return &Handler{
		service: auth,
	}
}

func (h *Handler) Routes(router *gin.Engine) {
	router.GET("/login/:guid", h.GenerateToken)
	router.GET("/refresh/:guid", h.Refresh)
}
