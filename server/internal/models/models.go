package models

import (
	"time"
)

type LoginReq struct {
	// 密码
	Password string `json:"password"`
	// 用户，具有唯一性
	Username string `json:"username"`
}

type RegReq struct {
	// 密码，8-16位
	Password string `json:"password"`
	// 用户名，3-20，不得有空格
	Username string `json:"username"`
}

type ConfessionNewReq struct {
	Anonymous   int    `json:"Anonymous"`
	Content     string `json:"content"`
	ScheduledAt string `json:"ScheduledAt"`
	UserID      int    `json:"user_id"`
}

type ConfessionEditReq struct {
	Anonymous     int    `json:"anonymous"`
	Content       string `json:"content"`
	PostID        int    `json:"post_id"`
	ScheduledTime string `json:"scheduled_time"`
}

type BlacklistRequest struct {
	UserID    string `json:"user_id"`
	BlockedID string `json:"blocked_id"`
}

type ProfileEditReq struct {
	ContactOther  string `json:"contact_other"`
	ContactQq     string `json:"contact_qq"`
	ContactTele   string `json:"contact_tele"`
	ContactWechat string `json:"contact_wechat"`
	Gender        string `json:"gender"`
	ImgURL        string `json:"img_url"`
	OtherInfo     string `json:"other_info"`
	Region        string `json:"region"`
	UserID        int    `json:"user_id"`
}

type Accounts struct {
	Username  string    `gorm:"column:username;size:20;not null;check:length(username) BETWEEN 3 AND 20"`               // 对应 username 列，长度为 20，不能为空
	UserID    int       `gorm:"column:user_id;primaryKey;autoIncrement"`                                                // 对应 user_id 列，设置为主键和自增
	UserType  int       `gorm:"column:user_type;not null;check:user_type IN (1, 2)"`                                    // 对应 user_type 列，不能为空，并限制其值为 1 或 2
	Password  string    `gorm:"column:password;size:255;not null;check:length(password) > 8 AND length(password) < 16"` // 对应 password 列，长度为 255，不能为空
	CreatedAt time.Time `gorm:"column:created_at;autoCreateTime"`                                                       // 对应 created_at 列，自动设置创建时间
}

type UserInfo struct {
	UserID        int       `gorm:"column:user_id;primaryKey;autoIncrement"`                                      // 对应 user_id 列，设置为主键和自增
	Username      string    `gorm:"column:username;size:20;not null;check:length(username) BETWEEN 3 AND 20"`     // 对应 username 列，长度为 20，不能为空
	Gender        string    `gorm:"column:gender;not null;type:ENUM('male', 'female', 'other', 'not_disclosed')"` // 对应 gender 列，不能为空，使用 ENUM 类型
	ContactTele   string    `gorm:"column:contact_tele;size:255;not null"`                                        // 对应 contact_tele 列，长度为 255，不能为空
	ContactQQ     string    `gorm:"column:contact_qq;size:255"`                                                   // 对应 contact_qq 列，长度为 255
	ContactWechat string    `gorm:"column:contact_wechat;size:255"`                                               // 对应 contact_wechat 列，长度为 255
	ContactOther  string    `gorm:"column:contact_other;size:255"`                                                // 对应 contact_other 列，长度为 255
	Region        string    `gorm:"column:region;size:255;not null"`                                              // 对应 region 列，长度为 255，不能为空
	ImgURL        string    `gorm:"type:varchar(255);not null;character_set:utf8mb3;collate:utf8mb3_general_ci"`
	OtherInfo     string    `gorm:"type:varchar(255);default:null"`
	CreatedAt     time.Time `db:"created_at"`
}

type Posts struct {
	PostID       int       `gorm:"column:post_id;primaryKey;autoIncrement"`                        // 帖子ID，主键，自增
	SenderID     int       `gorm:"column:sender_id;not null"`                                      // 发送者ID，不能为空
	Content      string    `gorm:"column:content;not null"`                                        // 帖子内容，不能为空
	CreatedAt    time.Time `gorm:"column:created_at;default:CURRENT_TIMESTAMP"`                    // 创建时间，默认当前时间
	ReportStatus int       `gorm:"column:report_status;not null;check:report_status IN (0, 1, 2)"` // 举报状态，0未举报，1举报审核中，2举报成功
	ScheduledAt  time.Time `gorm:"column:scheduled_time;default:CURRENT_TIMESTAMP"`                // 定时发送时间，默认当前时间
	Anonymous    int       `db:"anonymous"`                                                        // 使用标签来映射数据库列名
}
type Report struct {
	ReportID   int       `gorm:"column:report_id;primaryKey;autoIncrement"` // 举报ID，自增主键
	PostID     int       `gorm:"column:post_id;not null"`                   // 帖子ID，不能为空
	ReporterID int       `gorm:"column:reporter_id;not null"`               // 举报者ID，不能为空
	Reason     string    `gorm:"column:reason;size:255;not null"`           // 举报原因，不能为空
	CreatedAt  time.Time `gorm:"column:created_at;autoCreateTime"`          // 创建时间，自动设置为当前时间
}
type Trash struct {
	ReportID   int       `gorm:"column:report_id;primaryKey;autoIncrement"`     // 举报ID，主键，自增
	PostID     int       `gorm:"column:post_id;not null"`                       // 帖子ID，不能为空
	ReporterID int       `gorm:"column:reporter_id;not null"`                   // 举报者ID，不能为空
	Reason     string    `gorm:"column:reason;size:255;not null"`               // 举报原因，不能为空
	Status     int       `gorm:"column:status;not null;check:status IN (0, 1)"` // 处理结果：0 为举报失败，1 为举报成功
	CreatedAt  time.Time `gorm:"column:created_at;autoCreateTime"`              // 创建时间，默认当前时间
}
type Blacklist struct {
	BlackID         int       `gorm:"column:black_id;primaryKey;autoIncrement"`    // 黑名单ID，主键，自增
	UserID          int       `gorm:"column:user_id;not null"`                     // 用户ID，不能为空
	BlockedID       int       `gorm:"column:blocked_id;not null"`                  // 被拉黑的用户ID，不能为空
	BlockedUsername string    `gorm:"column:blocked_username;size:20;not null"`    // 被拉黑的用户名，不能为空
	CreatedAt       time.Time `gorm:"column:created_at;default:CURRENT_TIMESTAMP"` // 创建时间，默认当前时间
}
