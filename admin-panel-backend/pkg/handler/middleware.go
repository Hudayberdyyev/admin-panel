package handler

import (
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

const (
	authorizationHeader = "Authorization"
	userContext = "user_id"
)

func (h *Handler) userIdentity(c *gin.Context) {
	header := c.GetHeader(authorizationHeader)

	if header == "" {
		newErrorResponse(c, http.StatusUnauthorized, "empty auth header")
		return
	}

	headerParts := strings.Split(header, " ")
	if cap(headerParts) != 2 {
		newErrorResponse(c, http.StatusUnauthorized, "invalid auth header")
		return
	}

	userId, err := h.services.Authorization.Parse(headerParts[1])
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.Set(userContext, userId)
}

func getUserId(c *gin.Context) (int, error) {
	ctxId, ok := c.Get(userContext)
	if !ok {
		newErrorResponse(c, http.StatusInternalServerError, "userId not found from context")
		return 0, errors.New("userId not found from context")
	}

	userId, ok := ctxId.(int)
	if !ok {
		newErrorResponse(c, http.StatusInternalServerError, "userId is of invalid type")
		return 0, errors.New("userId is of invalid type")
	}

	return userId, nil
}
