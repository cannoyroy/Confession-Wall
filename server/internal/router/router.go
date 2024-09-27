package router

import (
	"server/internal/midwares"

	"server/internal/services"

	"github.com/gin-gonic/gin"
)

func Init(r *gin.Engine) {

	r.Use(midwares.Cors())

	r.GET("/main", func(c *gin.Context) {
		services.GetPost(c)
	})

}
