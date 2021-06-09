package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	backend "pragency"
	"strconv"
)

// Функция создания заказа

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

// Структура для получения json объекта из всех заказов

type getAllOrdersResponse struct {
	Data []backend.Order `json:"data"`
}

// Функция получения всех заказов пользователя

func (h *Handler) getAllOrders(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		return
	}

	orders, err := h.services.Order.GetAll(userId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, getAllOrdersResponse{
		Data: orders,
	})
}

// Функция получения всех заказов для админа

func (h *Handler) getAllOrdersAdmin(c *gin.Context) {
	orders, err := h.services.Order.GetAllAdmin()
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, getAllOrdersResponse{
		Data: orders,
	})
}

// Функция получения одного заказа для админа

func (h *Handler) getOrderByIdAdmin(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}

	order, err := h.services.Order.GetOneAdmin(id)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, order)
}

// Функция получения одного заказа для пользователя

func (h *Handler) getOrdersById(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		return
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}

	order, err := h.services.Order.GetById(userId, id)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, order)
}

// Функция обновления заказа для админа

func (h *Handler) updateOrderAdmin (c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}

	var input backend.UpdateOrderInput
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	if err := h.services.UpdateAdmin(id, input); err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, statusResponse{"ok"})
}

// Функция удаления заказа для админа

func (h *Handler) deleteOrderAdmin(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}

	err = h.services.Order.DeleteAdmin(id)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, statusResponse{
		Status: "ok",
	})
}