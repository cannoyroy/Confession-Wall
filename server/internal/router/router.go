package router

import (
	"server/internal/midwares"

	"server/internal/services"

	"github.com/gin-gonic/gin"
)

func Init(r *gin.Engine) {

	r.Use(midwares.Cors())

	r.POST("/login", func(c *gin.Context) {
		services.Login(c)
	})

	r.POST("/reg", func(c *gin.Context) {
		services.Reg(c)
	})

	r.POST("/confessions/new", func(c *gin.Context) {
		services.ConfessionPost(c)
	})

	r.GET("/confessions/manage", func(c *gin.Context) {
		services.ConfessionGet(c)
	})

	r.PUT("/confessions/manage/edit", func(c *gin.Context) {
		services.ConfessionEdit(c)
	})

	r.DELETE("/confessions/manage/delete", func(c *gin.Context) {
		services.ConfessionDele(c)
	})

	r.GET("/main", func(c *gin.Context) {
		services.GetPost(c)
	})

}
