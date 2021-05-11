package handler

import (
	"github.com/gin-gonic/gin"
	"pragency/pkg/service"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	auth := router.Group("/auth")
	{
		auth.POST("/sign-up", h.signUp)
		auth.POST("/sign-in", h.signIn)
	}

	api := router.Group("/api")
	{
		order := api.Group("/order")
		{
			order.POST("/", h.createOrder)
			order.GET("/", h.getAllOrders)
			order.GET("/:id", h.getOrderById)
			order.PUT("/:id", h.updateOrder)
			order.DELETE("/:id", h.deleteOrder)
		}
	}

	return router
}
