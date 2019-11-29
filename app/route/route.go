package route

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Route struct {
	Engine *gin.Engine
}

func (route *Route) Init()  {

	route.Engine.GET("/ping",func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	v1 :=route.Engine.Group("/v1")
	{
		v1.GET("/ping",func(c *gin.Context) {
			c.String(http.StatusOK, "pong")
		})
	}
}
