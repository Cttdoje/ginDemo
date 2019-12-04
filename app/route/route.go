package route

import (
	"ginDemo/app/controller/product"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Route struct {
	Engine *gin.Engine
}

func (route *Route) Init() {

	route.Engine.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	ProductRouter := route.Engine.Group("/v1")
	{
		ProductRouter.POST("/product", product.Add)
	}
}
