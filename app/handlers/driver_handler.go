package handlers

import (
	"net/http"

	"github.com/badfan/inno-taxi-driver-service/app/models"
	"github.com/gin-gonic/gin"
)

func (h *Handler) SignUp(c *gin.Context) {
	var input models.Driver

	if err := c.ShouldBindJSON(&input); err != nil {
		h.newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.driverService.SignUp(c.Request.Context(), &input)
	if err != nil {
		h.newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

func (h *Handler) SignIn(c *gin.Context) {
	var input models.Driver

	if err := c.ShouldBindJSON(&input); err != nil {
		h.newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	token, err := h.driverService.SignIn(c.Request.Context(), input.PhoneNumber, input.Password)
	if err != nil {
		h.newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"token": token,
	})
}

func (h *Handler) GetDriverRating(c *gin.Context) {
	id, ok := c.Get("driverID")
	if !ok {
		h.newErrorResponse(c, http.StatusInternalServerError, "driver id not found")
		return
	}

	conID, ok := id.(int)
	if !ok {
		h.newErrorResponse(c, http.StatusInternalServerError, "cannot convert id to type int")
		return
	}

	rating, err := h.driverService.GetDriverRating(c.Request.Context(), conID)
	if err != nil {
		h.newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"rating": rating,
	})
}

func (h *Handler) GetDriverStatus(c *gin.Context) {
	id, ok := c.Get("driverID")
	if !ok {
		h.newErrorResponse(c, http.StatusInternalServerError, "driver id not found")
		return
	}

	conID, ok := id.(int)
	if !ok {
		h.newErrorResponse(c, http.StatusInternalServerError, "cannot convert id to type int")
		return
	}

	status, err := h.driverService.GetDriverStatus(c.Request.Context(), conID)
	if err != nil {
		h.newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"is busy": status,
	})
}
