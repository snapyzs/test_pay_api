package handler

import (
	"github.com/gin-gonic/gin"
	"test_project_sell/internal/service"
)

type Handler struct {
	service *service.Service
}

func NewHandler(s *service.Service) *Handler {
	return &Handler{service: s}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	api := router.Group("/api")
	{
		api.POST("/create_pay", h.CreatePay)
		paySystem := api.Group("/pay_system", h.checkAuth)
		{
			paySystem.POST("/edit_pay", h.EditStatusPay)
		}
		api.POST("/check_pay", h.CheckPay)
		api.POST("/check_pay_userid", h.GetAllPayUserById)
		api.POST("/check_pay_email", h.GetAllPayUserByEmail)
		api.POST("/cancel_pay_id", h.CancelPayById)
		api.POST("/generate_token", h.GenerateTokenForUse)
	}

	return router
}
