package services

import (
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"server/internal/models"
)

func Login(c *gin.Context) {
	var req models.LoginReq
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"code": 400, "msg": "Invalid request"})
		return
	}

	var user models.Accounts
	if err := db.Where("username = ?", req.Username).First(&user).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusOK, gin.H{"code": 200506, "data": nil, "msg": "用户不存在"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"code": 500, "msg": "Internal server error"})
		return
	}

	if user.Password != req.Password {
		c.JSON(http.StatusOK, gin.H{"code": 200507, "data": nil, "msg": "密码错误"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"code": 200, "data": gin.H{"user_id": user.UserID, "username": user.Username, "user_type": user.UserType}, "msg": "success"})
}

func Reg(c *gin.Context) {
	var req models.RegReq
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"code": 400, "msg": "Invalid request"})
		return
	}

	// username不能有空格
	if strings.Contains(req.Username, " ") {
		c.JSON(http.StatusOK, gin.H{"code": 200502, "data": nil, "msg": "用户名不得有空格"})
		return
	}
	// 密码长度大于8位小于16位
	if len(req.Password) <= 8 || len(req.Password) >= 16 {
		c.JSON(http.StatusOK, gin.H{"code": 200503, "data": nil, "msg": "密码长度必须在8-16位"})
		return
	}

	// 用户名长度大于3位小于20位
	if len(req.Username) <= 3 || len(req.Username) >= 20 {
		c.JSON(http.StatusOK, gin.H{"code": 200504, "data": nil, "msg": "用户名长度必须在3-20位"})
		return
	}

	// 用户不能存在
	var user models.Accounts
	if err := db.Where("username = ?", req.Username).First(&user).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			// 继续注册
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"code": 500, "msg": "Internal server error"})
			return
		}
	} else {
		// 用户已存在
		c.JSON(http.StatusOK, gin.H{"code": 200505, "data": nil, "msg": "用户名已存在"})
		return
	}

	newUser := models.Accounts{
		Username: req.Username,
		// UserID:    0,
		UserType:  1,
		Password:  req.Password,
		CreatedAt: time.Now(),
	}
	if err := db.Create(&newUser).Error; err != nil {
		println(err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"code": 500, "msg": "Internal server error"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"code": 200, "data": nil, "msg": "success"})
}

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
		c.JSON(http.StatusInternalServerError, gin.H{"code": 500, "msg": "Internal server error"})
		return
	}
	// log.Println(BlacklistAll)
	type Post_1 struct {
		PostID      int
		SenderID    int
		Username    string
		Content     string
		CreatedAt   time.Time
		ScheduledAt time.Time
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
			var point = v.SenderID
			var Accounts_a []models.Accounts
			db.Where("user_id = ?", point).Find(&Accounts_a)
			var post = Post_1{
				PostID:      v.PostID,
				SenderID:    v.SenderID,
				Username:    Accounts_a[0].Username,
				Content:     v.Content,
				CreatedAt:   v.CreatedAt,
				ScheduledAt: v.ScheduledAt,
			}
			OUTposts = append(OUTposts, post)
		}
	}
	// log.Println(PostAll)
	c.JSON(http.StatusOK, gin.H{"code": 200, "data": gin.H{"post_list": OUTposts}, "msg": "success"})
}

func ConfessionPost(c *gin.Context) {
	var req models.ConfessionNewReq

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"code": 400, "msg": "Invalid request"})
		return
	}

	var newPost models.Posts

	layout := "2006-01-02 15:04:05.000"
	var location, err = time.LoadLocation("Asia/Shanghai")
	parsedTime, err := time.ParseInLocation(layout, req.ScheduledAt+":00.000", location)
	if err != nil {
		log.Println("Error parsing time:", err)
		return
	}
	newPost.SenderID = req.UserID
	newPost.Content = req.Content
	newPost.CreatedAt = time.Now()
	newPost.ReportStatus = 0
	newPost.ScheduledAt = parsedTime
	newPost.Anonymous = req.Anonymous

	if err := db.Create(&newPost).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"code": 501, "msg": "Internal server error"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"code": 200, "data": nil, "msg": "success"})
}

func ConfessionGet(c *gin.Context) {
	UserID := c.DefaultQuery("user_id", "0")
	var SelfPosts []models.Posts
	if err := db.Where("sender_id = ?", UserID).Find(&SelfPosts).Error; err != nil {
		// log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"code": 500, "msg": "Internal server error"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"code": 200, "data": gin.H{"post_list": SelfPosts}, "msg": "success"})
}

func ConfessionEdit(c *gin.Context) {
	var req models.ConfessionEditReq
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"code": 400, "msg": "Invalid request"})
		return
	}

	var TarPost models.Posts
	if err := db.Where("post_id = ?", req.PostID).First(&TarPost).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"code": 404, "data": nil, "msg": "Post not found or does not belong to the user"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"code": 500, "msg": "Internal server error"})
		return
	}

	layout := "2006-01-02 15:04:05.000"
	var location, err = time.LoadLocation("Asia/Shanghai")
	parsedTime, err := time.ParseInLocation(layout, req.ScheduledTime+":00.000", location)
	if err != nil {
		log.Println("Error parsing time:", err)
		return
	}
	TarPost.Anonymous = req.Anonymous
	TarPost.Content = req.Content
	TarPost.ScheduledAt = parsedTime
	log.Println(TarPost.ScheduledAt)
	if err := db.Save(&TarPost).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"code": 500, "msg": "Internal server error"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"code": 200, "data": nil, "msg": "success"})
}

func ConfessionDele(c *gin.Context) {
	PostID := c.DefaultQuery("post_id", "0")

	var TarPost models.Posts
	if err := db.Where("post_id = ?", PostID).First(&TarPost).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"code": 404, "data": nil, "msg": "不存在该帖子或者该帖子不属于该用户"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"code": 500, "msg": "Internal server error"})
		return
	}

	if err := db.Delete(&TarPost).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"code": 500, "msg": "Internal server error"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"code": 200, "data": nil, "msg": "success"})
}
