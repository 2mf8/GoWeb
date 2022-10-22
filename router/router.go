package router

import (
	"github.com/2mf8/GoWeb/api"
	"github.com/2mf8/GoWeb/middleware"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	r := gin.New()
	r.Use(middleware.CORS())
	r.GET("/", api.IndexApi)
	r.GET("/price/:item", api.PriceGetItemApi)
	r.GET("/prices/:key", api.PriceGetItemsApi)
	r.POST("/price", api.PriceAddAndUpdateByItemApi)
	r.POST("/price/:item", api.PriceAddAndUpdateByItemApi)
	r.DELETE("/price/:item", api.PriceDeleteByItemApi)
	return r
}
