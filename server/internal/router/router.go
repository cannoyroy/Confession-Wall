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

	r.GET("/blacklist/manage", func(c *gin.Context) {
		services.BlackGet(c)
	})

	r.DELETE("/blacklist/delete", func(c *gin.Context) {
		services.BlackDele(c)
	})

	r.POST("/blacklist/new", func(c *gin.Context) {
		services.BlacklistPost(c)
	})

	r.GET("/main", func(c *gin.Context) {
		services.GetPost(c)
	})

	r.GET("/profile", func(c *gin.Context) {
		services.ProfileGet(c)
	})

	r.PUT("/profile/edit", func(c *gin.Context) {
		services.ProfileEdit(c)
	})

	r.POST("/profile/new", func(c *gin.Context) {
		services.ProfilePost(c)
	})

	r.POST("/report", func(c *gin.Context) {
		services.ReportPost(c)
	})

	r.GET("/admin/reports", func(c *gin.Context) {
		services.GetReport(c)
	})

	r.POST("/admin/reports/handle", func(c *gin.Context) {
		services.ReportHandle(c)
	})

	r.GET("/admin/reports/history", func(c *gin.Context) {
		services.GetHistory(c)
	})

}
