package handler

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/Hudayberdyyev/admin-panel-backend/pkg/utils"
	"github.com/gin-gonic/gin"
)

func getCategoryAndAuthorQuery(catId, authId int) (string, []interface{}) {
	nArg := 0
	var argc []string
	var argv []interface{}
	if catId > 0 {
		nArg++
		argc = append(argc, fmt.Sprintf("n.category_id = $%d", nArg))
		argv = append(argv, catId)
	}
	if authId > 0 {
		nArg++
		argc = append(argc, fmt.Sprintf("n.author_id = $%d", nArg))
		argv = append(argv, authId)
	}
	query := strings.Join(argc, " and ")
	if len(query) > 0 {
		query = "and " + query
	}
	return query, argv
}

func (h *Handler) getAllNewsByCategoryAndAuthorId(c *gin.Context) {
	pagination := utils.GeneratePaginationFromRequest(c)
	lang := c.GetString(langCtx)

	categoryId, err := strconv.Atoi(c.Param("catId"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	authorId, err := strconv.Atoi(c.Param("authId"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	if categoryId < 0 {
		newErrorResponse(c, http.StatusBadRequest, "invalid category id")
		return
	}

	if authorId < 0 {
		newErrorResponse(c, http.StatusBadRequest, "invalid author id")
		return
	}

	selectorQuery, selectorParams := getCategoryAndAuthorQuery(categoryId, authorId)

	newsList, err := h.services.News.GetAllNewsByCategoryAndAuthorId(pagination, selectorQuery, selectorParams, lang)

	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, NewsListResponse{
		Data: newsList,
	})

}
