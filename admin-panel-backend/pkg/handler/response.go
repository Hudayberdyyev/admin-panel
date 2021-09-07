package handler

import (
	"github.com/Hudayberdyyev/admin-panel-backend/models"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type AuthorResponse struct {
	Data []models.Author `json:"data"`
}

type CategoriesResponse struct {
	Data []models.Category `json:"data"`
}

type NewsListResponse struct {
	Data []models.News `json:"data"`
}

type ErrorResponse struct {
	Message string `json:"message"`
}

type TagsResponse struct {
	Data []models.Tag `json:"data"`
}

type ContentResponse struct {
	Data []models.NewsContent `json:"data"`
}

func newErrorResponse(ctx *gin.Context, statusCode int, message string) {
	logrus.Errorf(message)
	ctx.AbortWithStatusJSON(statusCode, ErrorResponse{message})
}
