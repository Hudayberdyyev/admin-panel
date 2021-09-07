package handler

import (
	"github.com/Hudayberdyyev/admin-panel-backend/pkg/service"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	router.Use(cors.New(cors.Config{
		AllowMethods:     []string{"GET", "POST", "OPTIONS", "PUT", "DELETE"},
		AllowHeaders:     []string{"Origin", "Content-Length", "Content-Type", "User-Agent", "Referrer", "Host", "Token", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		AllowAllOrigins:  true,
		MaxAge:           86400,
	}))

	router.NoRoute(func(c *gin.Context) {
		c.JSON(404, map[string]interface{}{"code": "PAGE_NOT_FOUND", "message": "Page not found"})
	})
	router.MaxMultipartMemory = 10 << 20 // 10 MiB

	auth := router.Group("/auth")
	{
		auth.POST("/sign-up", h.signUp)
		auth.POST("/sign-in", h.signIn)
	}

	router.GET("/image/:filename", h.getImage)
	router.GET("/logo/:filename", h.getLogo)

	api := router.Group("/api", h.userIdentity)
	{
		lang := api.Group("/:lang", h.getLocale)
		{
			lists := lang.Group("/news")
			{
				lists.GET("/categories", h.getCategories)
				lists.GET("/authors", h.getAuthors)
				lists.GET("/count/authors", h.getNewsCountForAllAuthors)
				lists.GET("/count/categories", h.getNewsCountForAllCategories)
				lists.GET("/all/author/:authId/category/:catId", h.getAllNewsByCategoryAndAuthorId)
				content := lists.Group("/:id")
				{
					content.GET("/tags", h.getTagsByNewsTextId)
					content.GET("/content", h.getContentByNewsTextId)
				}

			}
		}
		extra := api.Group("/extra")
		{
			extra.GET("/all", h.getAllExtraMessages)
			extra.POST("/", h.createExtraMessage)
			extra.GET("/delete/:id", h.deleteExtraMessage)
		}
	}

	return router
}
