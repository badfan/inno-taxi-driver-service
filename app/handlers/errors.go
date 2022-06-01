package handlers

import "github.com/gin-gonic/gin"

type ErrorMessage struct {
	Message string
}

func (h *Handler) newErrorResponse(c *gin.Context, statusCode int, message string) {
	h.logger.Error(message)
	c.AbortWithStatusJSON(statusCode, ErrorMessage{Message: message})
}
