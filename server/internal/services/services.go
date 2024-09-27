package services

import (
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"

	"server/internal/models"
)

func GetPost(c *gin.Context) {
	UserID := c.DefaultQuery("user_id", "0")

	var PostAll []models.Posts
	if err := db.Where("report_status != ?", 2).Find(&PostAll).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"code": 500, "msg": "Internal server error"})
		return
	}
	// log.Println(PostAll)
	// log.Println("--------")
	var BlacklistAll []models.Blacklist
	if err := db.Where("user_id = ?", UserID).Find(&BlacklistAll).Error; err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"code": 500, "msg": "Internal s!!erver error"})
		return
	}
	// log.Println(BlacklistAll)
	type Post_1 struct {
		PostID       int
		SenderID     int
		Content      string
		CreatedAt    time.Time
		ReportStatus int
		ScheduledAt  time.Time
	}
	// for _, k := range BlacklistAll {
	// 	log.Println(k.BlockedID)
	// }
	var OUTposts []Post_1
	for _, v := range PostAll {
		var a = 1
		for _, k := range BlacklistAll {
			if v.SenderID == k.BlockedID {
				a = 0
				break
			}
		}
		if a == 1 {
			var post = Post_1{
				PostID:       v.PostID,
				SenderID:     v.SenderID,
				Content:      v.Content,
				CreatedAt:    v.CreatedAt,
				ReportStatus: v.ReportStatus,
				ScheduledAt:  v.ScheduledAt,
			}
			OUTposts = append(OUTposts, post)
		}
	}
	// log.Println(PostAll)
	c.JSON(http.StatusOK, gin.H{"code": 200, "data": gin.H{"post_list": OUTposts}, "msg": "success"})
}
