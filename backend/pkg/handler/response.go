package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

// Структура ошибки json

type errorResponse struct {
	Message string `json:"message"`
}

// Структура статуса запроса json

type statusResponse struct {
	Status string `json:"status"`
}

// Функция создания ошибки через logrus

func newErrorResponse(c *gin.Context, statusCode int, message string) {
	logrus.Errorf(message)
	c.AbortWithStatusJSON(statusCode, errorResponse{message})
}