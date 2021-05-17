package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) createOrder(c *gin.Context) {
	id, _ := c.Get(userCtx)
	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

func (h *Handler) getAllOrders(c *gin.Context) {

}

func (h *Handler) getOrdersById(c *gin.Context) {

}

func (h *Handler) updateOrder(c *gin.Context) {

}

func (h *Handler) deleteOrder(c *gin.Context) {

}