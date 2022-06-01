package handlers

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

const (
	authorizationHeader = "Authorization"
)

func (h *Handler) Middleware(c *gin.Context) {
	header := c.GetHeader(authorizationHeader)
	if header == "" {
		h.newErrorResponse(c, http.StatusUnauthorized, "empty authorization header")
		return
	}

	headerParts := strings.Split(header, " ")
	if len(headerParts) != 2 {
		h.newErrorResponse(c, http.StatusUnauthorized, "invalid authorization header")
		return
	}

	id, err := h.authService.ParseToken(headerParts[1])
	if err != nil {
		h.newErrorResponse(c, http.StatusUnauthorized, err.Error())
		return
	}

	c.Set("driverID", id)
}
