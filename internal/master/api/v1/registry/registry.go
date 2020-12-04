package registry

import (
	"github.com/gin-gonic/gin"
	"shortLink/internal/master/api/v1/handle"
)

var Router = gin.Default()

func init() {
	h := new(handle.Handle)
	Router.GET("/:id", h.Get)
	v1 := Router.Group("/v1/api")
	{
		v1.POST("/", h.AddLink)
	}
}
