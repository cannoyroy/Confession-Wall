package services

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
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
		c.JSON(http.StatusOK, gin.H{"code": 200503, "data": nil, "msg": "密码长度必须在9-15位"})
		return
	}

	// 用户名长度大于3位小于20位
	if len(req.Username) <= 3 || len(req.Username) >= 20 {
		c.JSON(http.StatusOK, gin.H{"code": 200504, "data": nil, "msg": "用户名长度必须在4-19位"})
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
	Page := c.DefaultQuery("page", "1")
	Limit := c.DefaultQuery("limit", "100")
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
	var cnt = 0

	limit, err := strconv.Atoi(Limit)
	if err != nil {
		fmt.Println("转换失败:", err)
		return
	}
	page, err := strconv.Atoi(Page)
	if err != nil {
		fmt.Println("转换失败:", err)
		return
	}
	var start = limit*(page-1) + 1
	var end = limit * page
	for _, v := range PostAll {
		var a = 1
		for _, k := range BlacklistAll {
			if v.SenderID == k.BlockedID {
				a = 0
				break
			}
		}
		now := time.Now()
		if now.After(v.ScheduledAt) {
			if a == 1 {
				cnt = cnt + 1
				if cnt < start {
					continue
				}
				if cnt > end {
					break
				}
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
				if v.Anonymous == 1 {
					post.Username = "DEAR"
				}
				OUTposts = append(OUTposts, post)

			}
		}
	}
	// log.Println(PostAll)
	c.JSON(http.StatusOK, gin.H{"code": 200, "data": gin.H{"post_list": OUTposts, "total": cnt}, "msg": "success"})
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

func BlackGet(c *gin.Context) {
	UserID := c.DefaultQuery("user_id", "0")
	var SelfBlack []models.Blacklist
	if err := db.Where("user_id = ?", UserID).Find(&SelfBlack).Error; err != nil {
		// log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"code": 500, "msg": "Internal server error"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"code": 200, "data": gin.H{"post_list": SelfBlack}, "msg": "success"})
}

func BlackDele(c *gin.Context) {
	BlockedID := c.DefaultQuery("blocked_id", "0")
	UserID := c.DefaultQuery("user_id", "0")

	var TarBlack models.Blacklist
	if err := db.Where("blocked_id = ? and user_id = ?", BlockedID, UserID).First(&TarBlack).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"code": 404, "data": nil, "msg": "不存在该黑名单"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"code": 500, "msg": "Internal server error"})
		return
	}

	if err := db.Delete(&TarBlack).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"code": 500, "msg": "Internal server error"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"code": 200, "data": nil, "msg": "success"})
}

func BlacklistPost(c *gin.Context) {
	// BlockedID := c.DefaultQuery("blocked_id", "0")
	// log.Println(BlockedID)
	// UserID := c.DefaultQuery("user_id", "0")

	var req models.BlacklistRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		// 如果解析失败，返回错误
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		log.Println(err.Error())
		return
	}
	BlockedID := req.BlockedID
	UserID := req.UserID

	var newBlack models.Blacklist
	var user models.Accounts

	if err := db.Where("user_id = ?", BlockedID).First(&user).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusOK, gin.H{"code": 200506, "data": nil, "msg": "用户不存在"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"code": 500, "msg": "Internal server error"})
		return
	}
	var err error
	newBlack.UserID, err = strconv.Atoi(UserID)
	if err != nil {
		// 错误处理
		log.Printf("Error converting UserID to int: %v\n", err)
		return
	}
	newBlack.BlockedID, err = strconv.Atoi(BlockedID)
	newBlack.BlockedUsername = user.Username
	if err != nil {
		// 错误处理
		log.Printf("Error converting UserID to int: %v\n", err)
		return
	}

	if err := db.Create(&newBlack).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"code": 501, "msg": "Internal server error"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"code": 200, "data": nil, "msg": "success"})
}

func ProfileGet(c *gin.Context) {
	UserID := c.DefaultQuery("user_id", "0")
	var SelfProfile []models.UserInfo
	if err := db.Where("user_id = ?", UserID).Find(&SelfProfile).Error; err != nil {
		// log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"code": 500, "msg": "Internal server error"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"code": 200, "data": gin.H{"post_list": SelfProfile}, "msg": "success"})
}

func ProfileEdit(c *gin.Context) {
	var req models.ProfileEditReq
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"code": 400, "msg": "Invalid request"})
		return
	}

	var TarProfile models.UserInfo
	if err := db.Where("user_id = ?", req.UserID).First(&TarProfile).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"code": 404, "data": nil, "msg": "Post not found or does not belong to the user"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"code": 500, "msg": "Internal server error"})
		return
	}

	// TarProfile.Gender = req.Gender
	// TarProfile.ContactTele = req.ContactTele
	// TarProfile.ContactQQ = req.ContactQq
	// TarProfile.ContactWechat = req.ContactWechat
	// TarProfile.ContactOther = req.ContactOther
	// TarProfile.Region = req.Region
	// TarProfile.ImgURL = req.ImgURL
	// TarProfile.OtherInfo = req.OtherInfo

	if req.Gender != "none" {
		TarProfile.Gender = req.Gender
	}
	if req.ContactTele != "none" {
		TarProfile.ContactTele = req.ContactTele
	}
	if req.ContactQq != "none" {
		TarProfile.ContactQQ = req.ContactQq
	}
	if req.ContactWechat != "none" {
		TarProfile.ContactWechat = req.ContactWechat
	}
	if req.ContactOther != "none" {
		TarProfile.ContactOther = req.ContactOther
	}
	if req.Region != "none" {
		TarProfile.Region = req.Region
	}
	if req.ImgURL != "none" {
		TarProfile.ImgURL = req.ImgURL
	}
	if req.OtherInfo != "none" {
		TarProfile.OtherInfo = req.OtherInfo
	}

	if err := db.Save(&TarProfile).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"code": 500, "msg": "Internal server error"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"code": 200, "data": nil, "msg": "success"})
}

func ProfilePost(c *gin.Context) {
	var req models.ProfileEditReq

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"code": 400, "msg": "Invalid request"})
		return
	}

	var newProfile models.UserInfo
	var user models.Accounts
	log.Println(55555)
	log.Println(req)
	if err := db.Where("user_id = ?", req.UserID).First(&user).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusOK, gin.H{"code": 200506, "data": nil, "msg": "用户不存在"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"code": 500, "msg": "Internal server error"})
		return
	}
	newProfile.UserID = req.UserID
	newProfile.Username = user.Username
	newProfile.Gender = req.Gender
	newProfile.ContactTele = req.ContactTele
	newProfile.ContactQQ = req.ContactQq
	newProfile.ContactWechat = req.ContactWechat
	newProfile.ContactOther = req.ContactOther
	newProfile.Region = req.Region
	newProfile.ImgURL = req.ImgURL
	newProfile.OtherInfo = req.OtherInfo
	newProfile.CreatedAt = time.Now()
	if err := db.Create(&newProfile).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"code": 501, "msg": "Internal server error"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"code": 200, "data": nil, "msg": "success"})
}

func ReportPost(c *gin.Context) {
	var req models.ReportReq
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
	TarPost.ReportStatus = 1

	if err := db.Save(&TarPost).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"code": 500, "msg": "Internal server error"})
		return
	}

	RepotFresh := models.Report{
		PostID:     req.PostID,
		ReporterID: req.ReporterID,
		Reason:     req.Reason,
		CreatedAt:  time.Now(),
	}
	if err := db.Create(&RepotFresh).Error; err != nil {
		println(err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"code": 501, "msg": "Internal server error"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"code": 200, "data": nil, "msg": "success"})
}

func GetReport(c *gin.Context) {
	var TarReport []models.Report
	if err := db.Find(&TarReport).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"code": 500, "msg": "Internal server error"})
		return
	}
	type out struct {
		ReportID         int
		PostID           int
		ReporterID       int
		ReporterUsername string
		ReportedContent  string
		Reason           string
		CreatedAt        time.Time
	}
	var OUTposts []out
	for _, v := range TarReport {
		var temp out
		temp.PostID = v.PostID
		temp.Reason = v.Reason
		temp.ReporterID = v.ReporterID
		temp.ReportID = v.ReportID

		var RawPost models.Posts
		if err := db.Where("post_id = ?", v.PostID).Find(&RawPost).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"code": 500, "msg": "Internal server error"})
			return
		}
		temp.ReportedContent = RawPost.Content

		var RawUser models.UserInfo
		if err := db.Where("user_id = ?", v.ReporterID).Find(&RawUser).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"code": 500, "msg": "Internal server error"})
			return
		}
		temp.ReporterUsername = RawUser.Username
		// temp.State = RawPost.State

		OUTposts = append(OUTposts, temp)
	}
	c.JSON(http.StatusOK, gin.H{"code": 200, "data": gin.H{"report_list": OUTposts}, "msg": "success"})
}

func ReportHandle(c *gin.Context) {
	var req models.HandleReportReq
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"code": 400, "msg": "Invalid request"})
		return
	}

	var TarReport []models.Report
	if err := db.Where("post_id = ?", req.PostID).Find(&TarReport).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"code": 501, "msg": "Internal server error"})
		return
	}

	var TarPost models.Posts
	if err := db.Where("post_id = ?", req.PostID).First(&TarPost).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"code": 404, "data": nil, "msg": "Post not found or does not belong to the user"})
			return
		}
		println(err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"code": 501, "msg": "Internal server error"})
		return
	}
	var trashst = 0
	if req.Approval == 2 {
		TarPost.ReportStatus = 0
		trashst = 0 //举报失败
	} else {
		TarPost.ReportStatus = 2
		trashst = 1 //举报成功
	}
	if err := db.Save(&TarPost).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"code": 502, "msg": "Internal server error"})
		return
	}

	// var TarPost_1 []models.Report
	// if err := db.Where("post_id = ?", TarReport.PostID).Find(&TarPost_1).Error; err != nil {
	// 	if err == gorm.ErrRecordNotFound {
	// 		c.JSON(http.StatusNotFound, gin.H{"code": 404, "data": nil, "msg": "Post not found or does not belong to the user"})
	// 		return
	// 	}
	// 	c.JSON(http.StatusInternalServerError, gin.H{"code": 503, "msg": "Internal server error"})
	// 	return
	// }

	var newTrash []models.Trash

	// 迁移 TarPost_1 中的数据到 Trash 表
	for _, report := range TarReport {
		trash := models.Trash{
			ReportID:   report.ReportID,
			PostID:     report.PostID,
			ReporterID: report.ReporterID,
			Reason:     report.Reason,
			Status:     trashst,
			CreatedAt:  time.Now(), // 设置为当前时间
		}
		newTrash = append(newTrash, trash)
	}

	if err := db.Create(&newTrash).Error; err != nil {
		println(err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"code": 550, "msg": "Internal server error"})
		return
	}

	if err := db.Delete(&TarReport).Error; err != nil {
		println(err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"code": 504, "msg": "Internal server error"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"code": 200, "data": nil, "msg": "success"})
}

func GetHistory(c *gin.Context) {
	var TarTrash []models.Trash
	if err := db.Find(&TarTrash).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"code": 500, "msg": "Internal server error"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"code": 200, "data": gin.H{"report_list": TarTrash}, "msg": "success"})
}
