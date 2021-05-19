package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	backend "pragency"
)

func (h *Handler) createOrder(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		return
	}

	var input backend.Order
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.services.Order.Create(userId, input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
	}

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