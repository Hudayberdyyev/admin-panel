package handler

import (
	"context"
	"errors"
	"io"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/Hudayberdyyev/admin-panel-backend/storage"
	"github.com/gin-gonic/gin"
	"github.com/minio/minio-go/v7"
)

func (h *Handler) getCategories(c *gin.Context) {
	lang := c.GetString(langCtx)

	categories, err := h.services.News.GetAllCategories(lang)

	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, CategoriesResponse{
		Data: categories,
	})
}

func (h *Handler) getAuthors(c *gin.Context) {
	lang := c.GetString(langCtx)

	authors, err := h.services.News.GetAllAuthors(lang)

	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, AuthorResponse{
		Data: authors,
	})
}

func getBucketAndFilename(filename string) (string, string, error) {
	var bucketName, pattern string

	if strings.Contains(filename, storage.NewsBucketPattern) {
		bucketName = "news"
		pattern = storage.NewsBucketPattern
	} else {
		bucketName = "content"
		pattern = storage.ContentBucketPattern
	}

	entryPos := strings.Index(filename, pattern)
	if entryPos < 0 {
		return "", "", errors.New("invalid image filename " + filename)
	}

	filename = filename[entryPos+len(pattern):]
	return bucketName, filename, nil
}

func (h *Handler) getImage(c *gin.Context) {

	filename := c.Param("filename")

	bucketname, filename, err := getBucketAndFilename(filename)

	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	object, err := storage.ImageStorage.Client.GetObject(context.Background(), bucketname, filename, minio.GetObjectOptions{})

	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	if _, err = io.Copy(c.Writer, object); err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"ok": "ok",
	})
}

func (h *Handler) getLogo(c *gin.Context) {

	filename := c.Param("filename")

	if filename == "" {
		newErrorResponse(c, http.StatusInternalServerError, "empty filename")
		return
	}

	filepath := "logo/" + filename
	object, err := os.Open(filepath)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	if _, err = io.Copy(c.Writer, object); err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"ok": "ok",
	})
}

func (h *Handler) getTagsByNewsTextId(c *gin.Context) {
	newsTextId, convErr := strconv.Atoi(c.Param("id"))

	if convErr != nil {
		newErrorResponse(c, http.StatusBadRequest, convErr.Error())
		return
	}

	lang := c.GetString(langCtx)

	tags, err := h.services.News.GetTagsByNewsTextId(newsTextId, lang)

	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, TagsResponse{
		Data: tags,
	})
}

func (h *Handler) getContentByNewsTextId(c *gin.Context) {
	newsTextId, convErr := strconv.Atoi(c.Param("id"))

	if convErr != nil {
		newErrorResponse(c, http.StatusBadRequest, convErr.Error())
		return
	}

	contents, err := h.services.News.GetContentByNewsTextId(newsTextId)

	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, ContentResponse{
		Data: contents,
	})
}

func (h *Handler) getNewsCountForAllAuthors(c *gin.Context) {
	Data, err := h.services.News.GetNewsCountForAllAuthors()
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"Data": Data,
	})
}

func (h *Handler) getNewsCountForAllCategories(c *gin.Context) {
	lang := c.GetString(langCtx)

	data, err := h.services.News.GetNewsCountForAllCategories(lang)

	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"Data": data,
	})
}
