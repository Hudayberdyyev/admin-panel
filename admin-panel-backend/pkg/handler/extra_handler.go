package handler

import (
	"net/http"
	"strconv"

	"github.com/Hudayberdyyev/admin-panel-backend/models"
	"github.com/gin-gonic/gin"
)

func (h *Handler) createExtraMessage(c *gin.Context) {
	var input models.ExtraMessages

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.services.ExtraMessages.Create(input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"id": id,
	})
}

func (h *Handler) deleteExtraMessage(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	if err = h.services.ExtraMessages.Delete(id); err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"ok": "ok",
	})
}

func (h *Handler) getAllExtraMessages(c *gin.Context) {
	data, err := h.services.ExtraMessages.GetAll()
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"Data": data,
	})
}
