package handler

import (
	"github.com/gin-gonic/gin"
	"pragency/pkg/service"
)

// Структура хэндлеров

type Handler struct {
	services *service.Service
}

// Функция передачи в сервисы

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

// Функция работы с CORS для приема http запросов без строгой защиты

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Credentials", "true")
		c.Header("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Header("Access-Control-Allow-Methods", "POST, GET, PUT, DELETE, HEAD, PATCH, OPTIONS")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}

// Функция роутов описывающих запросы по ссылкам

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()
	router.Use(CORSMiddleware())

	auth := router.Group("/auth")
	{
		auth.POST("/sign-up", h.signUp)
		auth.POST("/sign-in", h.signIn)
	}

	api := router.Group("/api", h.userIdentity)
	{
		order := api.Group("/order")
		{
			order.POST("", h.createOrder)
			order.GET("", h.getAllOrders)
			order.GET("/:id", h.getOrdersById)
		}

		admin := api.Group("/admin")
		{
			admin.GET("", h.getAllOrdersAdmin)
			admin.GET("/:id", h.getOrderByIdAdmin)
			admin.PUT("/:id", h.updateOrderAdmin)
			admin.DELETE("/:id", h.deleteOrderAdmin)
		}
	}

	return router
}
