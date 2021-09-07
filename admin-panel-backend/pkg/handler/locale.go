package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

const (
	nLANG   = 2
	langCtx = "lang"
)

var LANG [nLANG]string = [nLANG]string{
	"tm",
	"ru",
}

func (h *Handler) getLocale(c *gin.Context) {
	lang := c.Param("lang")
	if lang == "" {
		newErrorResponse(c, http.StatusBadRequest, "empty language")
		return
	}

	isPossibleLang := false
	for _, v := range LANG {
		if lang == v {
			isPossibleLang = true
			break
		}
	}
	if !isPossibleLang {
		newErrorResponse(c, http.StatusBadRequest, "invalid type of language")
		return
	}

	c.Set(langCtx, lang)
}
