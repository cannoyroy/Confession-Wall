# 基本信息

腾讯云MySQL主机密码为 -15d79-a

## 人员

我是健行学院理工的钱枷羽，暑期是完成了前后端的开发，但是学的没有很深入、系统。

【前端】我是土木学院的马逸轩，暑期完成了前端的大作业，学的也没有很深入系统。

【前端】我是环境学院的张梓煊，会一点算法，前后端零基础（但可以学）

【后端】我是健行学院理工的谢安邦，在前端和后端上算是小白，但是愿意花一点时间去学习这方面

## 官方要求

基础要求 0. 前端使用Vue全家桶，使用Route进行路由管理，ts进行类型规范，配置路由守卫.

1. 登陆与注册功能，前端要求实现一段时间内保留登录状态，后端要求使用常见的鉴 权方式。 *和论坛差不多*
2. 实现个人页面，能够修改昵称和设置自己上传的图片作为头像，账号，登陆密码等个人 信息
3. 可以发布一条想对别人说的表白，且能够选择是否匿名以及是否公开*和论坛差不多*
4. 能够管理自己的表白（修改、删除等）*和论坛一致*
5. 实现社区功能，能看到别人发的表白（注意实名和匿名）*和论坛一致*
6. 实现拉黑功能（看不到拉黑人所发的表白）*和论坛差不多*
7. 使用github进行版本管理，采用合理、规范的目录结构，使用易读、标准的变量命名 形式
8. 完成前后端联调，可以用网页展示项目的全部效果

提高要求

1. 实现表白消息的评论和回复评论的功能-https://github.com/readpage/undraw-ui
2. 可以自定义全局字体大小和主题色-https://theme-hope.vuejs.press/zh/guide/interface/theme-color.html#设置默认主题色
3. 后端成功部署到云端服务器，可以公网访问网站服务？
4. 后端进行使用全局错误统一处理中间件，并有设置统一错误码
5. 实现页面加载加速（前端cdn或后端缓存）
6. 可以发送自定义定时表白消息

## 其他通知

团队协作要用到git，要把项目传到github上，并在文档里填写前后端仓库地址，我知道你们每组都有人能上github的，10月2号前要把地址补充完全[表情]

你们可以先团队内部先把项目的具体需求定下来，比如项目基础要求要怎么实现，要不要考虑尝试哪个提高点，交流完后端先把apifox完善，前端使用接口时可以自己写mock或者调后端接口，但我们还是建议尽早进行前后端对接(自己写mock不如push后端)，因为如果前后端都差不多完成了，但前后端对接失败，再修改接口会相对比较麻烦



# 9.26

基础要求 

0. 前端使用Vue全家桶，使用Route进行路由管理，ts进行类型规范，配置路由守卫.

1. 登陆与注册功能，前端要求实现一段时间内保留登录状态，后端要求使用常见的鉴权方式。 
2. 实现个人页面，能够修改昵称和设置自己上传的图片作为头像，账号，登陆密码等个人 信息
3. 可以发布一条想对别人说的表白，且能够选择是否匿名以及是否公开
4. 能够管理自己的表白（修改、删除等）
5. 实现社区功能，能看到别人发的表白（注意实名和匿名）
6. 实现拉黑功能（看不到拉黑人所发的表白）
7. 使用github进行版本管理，采用合理、规范的目录结构，使用易读、标准的变量命名 形式
8. 完成前后端联调，可以用网页展示项目的全部效果



```
9.26-9.29
登陆与注册界面
基本框架

10.1
个人页面

10.2
看贴界面
发帖界面
管理界面
拉黑界面

10.3
评论界面
全局字体/主题色

————手机端适配


周六9.28 学完——代码都能看懂，能自己跑起来，自己敲一遍代码——https://apifox.com/apidoc/shared-e228cfbb-d012-45df-943f-268ea9ad8f60/api-195488732 ——POST，GET，DELETE
```



提高要求

1. 实现表白消息的评论和回复评论的功能-‣
2. 可以自定义全局字体大小和主题色-‣
3. 后端成功部署到云端服务器，可以公网访问网站服务？
4. 后端进行使用全局错误统一处理中间件，并有设置统一错误码
5. 实现页面加载加速（前端cdn或后端缓存）
6. 可以发送自定义定时表白消息



别的要求：

1. 点赞，分享【链接简单】
2. 个人信息 要包括 性别
3. 添加对象信息
4. 内容生成AI【暂定，放到最后去做】
5. 历史记录，给别人点赞的记录
6. 关键词搜索筛选（帖子关键词，人名匹配）【搜索怎么实现】
7. 审核机制：第一层（关键词屏蔽）第二层（举报删除）
8. 通知：成功发出……【？？放后面】
9. 个人界面与管理员界面
10. 帖子 标签（捞人、组队、爱情、友情、闲置、趣事、爆料、招新）
11. ~~AI2: 外接AI，对表白语言进行分析，然后评估成功概率或给出建议（这个可以根据表白对象的大致性格，但是终究只是幻想）（这个难搞，需要数据调教，不知道能不能直接抄，还有就是外接需要考虑api花费，不一定有米，但是这个真搞出来，我感觉真的强）~~
12. 用户排行榜：ZZX给个算法【简单】
13. 编辑草稿箱【简单，放后面】
14. 表白回顾：显示给自己看
15. 推荐，用户登陆后，可以来个 恋爱墙别的人（异性为主）发的帖子，搞个视觉冲击
16. 漂流瓶 随机抽
17. 熟悉后，共享身份信息
18. 情侣间游玩美食推荐，情侣贴士



# ALL



APIFOX文档

数据库

前后端交互逻辑

后端部署服务器？



所有的GO的路由和函数 丢给XAB

### 前后端交互

Reg：

{    "username": "延芳",    "password": "occ56ut" }

{    "code": 200,    "data": null,    "msg": "success" }

{    "code": 200502,    "data": null,    "msg": "用户名不得有空格" }

{    "code": 200503,    "data": null,    "msg": "密码长度必须在8-16位" }

{    "code": 200504,    "data": null,    "msg": "用户名长度必须在3-20位" }

{    "code": 200505,    "data": null,    "msg": "用户名已存在" }

Login：

{    "username": "123",    "password": "12345678" }

{"code":200,"data":{"user_id":2,"user_type":1},"msg":"success"}

{"code":200506,"data":null,"msg":"用户不存在"}

{"code":200507,"data":null,"msg":"密码错误"}



### 路由

以下是满足您需求的网站路由名称和结构设计：

### 路由名称和结构

- **/login**  — 登录页面
- **/register** — 注册页面



- **/main** — 主页面



- **/profile** — 个人信息页面
  - **/profile/edit** — 修改个人信息（昵称、头像、密码）



- **/confessions** — 表白发布页面
  - **/confessions/new** — 发布新表白
  - **/confessions/manage** — 管理自己的表白
  - **/confessions/manage/edit/:id** — 修改表白
  - **/confessions/manage/delete/:id** — 删除表白
  - **/confessions/scheduled-confessions** — 定时表白消息管理



- **/blacklist** — 管理拉黑列表
  - **/blacklist/manage**
  - **/blacklist/delete**




- **/confessions/:id/comments** — 查看表白的评论 ——暂时不做
  - **/confessions/:id/comments/new** — 发布新评论
  - **/confessions/:id/comments/:commentId/replies** — 回复评论



- **/settings** — 设置页面
  - **/settings/theme** — 主题和字体设置







- /admin/reports

   — 查看所有举报的帖子

  - /admin/reports/handle

     — 处理举报

  - /admin/reports/history

    —处理记录



- /admin/posts——暂时不做

   — 管理所有帖子

  - /admin/posts/

    — 查看和编辑特定帖子

```
/src
  /components
    /Auth
      Login.vue
      Register.vue
    /Admin
      Reports.vue
      PostManagement.vue
      ReportHistory.vue
      EditPost.vue
    /Profile
      Profile.vue
      EditProfile.vue
    /Confessions
      ConfessionList.vue
      NewConfession.vue
      ManageConfessions.vue
      EditConfession.vue
    /Community
      Community.vue
    /Comments
      CommentList.vue
      NewComment.vue
    /Settings
      Settings.vue
      ThemeSettings.vue
  /router
    index.ts
  /store
    index.ts
  /views
    Home.vue
    NotFound.vue
  App.vue
  main.ts
```



### 数据库

##### 账号表：

```sql
CREATE TABLE accounts (
    username VARCHAR(20) NOT NULL CHECK (LENGTH(username) BETWEEN 3 AND 20),
    user_id INT(8) PRIMARY KEY AUTO_INCREMENT,
    user_type TINYINT NOT NULL CHECK (user_type IN (1, 2)),
    password VARCHAR(255) NOT NULL CHECK (LENGTH(password) > 8 AND LENGTH(password) < 16),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
```

```go
type Accounts struct {
	Username  string    `gorm:"column:username;size:20;not null;check:length(username) BETWEEN 3 AND 20"` // 对应 username 列，长度为 20，不能为空
	UserID    int       `gorm:"column:user_id;primaryKey;autoIncrement"`                                         // 对应 user_id 列，设置为主键和自增
	UserType  int       `gorm:"column:user_type;not null;check:user_type IN (1, 2)"`                           // 对应 user_type 列，不能为空，并限制其值为 1 或 2
	Password  string    `gorm:"column:password;size:255;not null;check:length(password) > 8 AND length(password) < 16"` // 对应 password 列，长度为 255，不能为空
	CreatedAt time.Time `gorm:"column:created_at;autoCreateTime"`                                               // 对应 created_at 列，自动设置创建时间
}
```

##### 用户信息表：

用户名：

性别：

联系方式：

地区：

```sql
CREATE TABLE `user_infos` (
    `user_id` bigint NOT NULL AUTO_INCREMENT,
    `username` varchar(20) NOT NULL,
    `gender` enum(
        'male',
        'female',
        'other',
        'not_disclosed'
    ) NOT NULL,
    `contact_tele` varchar(255) NOT NULL,
    `contact_qq` varchar(255) DEFAULT NULL,
    `contact_wechat` varchar(255) DEFAULT NULL,
    `contact_other` varchar(255) DEFAULT NULL,
    `region` varchar(255) NOT NULL,
    `img_url` varchar(255) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL,
    `other_info` varchar(255) DEFAULT NULL,
    `CreatedAt` date NOT NULL,
    PRIMARY KEY (`user_id`),
    CONSTRAINT `chk_user_infos_username` CHECK (
        (
            length(`username`) between 3 and 20
        )
    )
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb3
```

```go
type UserInfo struct {
	UserID        int    `gorm:"column:user_id;primaryKey;autoIncrement"`                             // 对应 user_id 列，设置为主键和自增
	Username      string `gorm:"column:username;size:20;not null;check:length(username) BETWEEN 3 AND 20"` // 对应 username 列，长度为 20，不能为空
	Gender        string `gorm:"column:gender;not null;type:ENUM('male', 'female', 'other', 'not_disclosed')"` // 对应 gender 列，不能为空，使用 ENUM 类型
	ContactTele   string `gorm:"column:contact_tele;size:255;not null"`                               // 对应 contact_tele 列，长度为 255，不能为空
	ContactQQ     string `gorm:"column:contact_qq;size:255"`                                          // 对应 contact_qq 列，长度为 255
	ContactWechat  string `gorm:"column:contact_wechat;size:255"`                                      // 对应 contact_wechat 列，长度为 255
	ContactOther   string `gorm:"column:contact_other;size:255"`                                       // 对应 contact_other 列，长度为 255
	Region        string `gorm:"column:region;size:255;not null"`                                     // 对应 region 列，长度为 255，不能为空
}
```

##### 帖子表

帖子ID

发送者ID

帖子内容

创建时间 

举报状态：0未举报，1举报审核中，2举报成功

匿名状态：0不匿名，1匿名

```sql
CREATE TABLE posts (
    post_id INT(8) PRIMARY KEY AUTO_INCREMENT,                         -- 帖子ID，主键，自增
    sender_id INT(8) NOT NULL,                                        -- 发送者ID，不能为空
    content TEXT NOT NULL,                                            -- 帖子内容，不能为空
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,                   -- 创建时间，默认当前时间
    report_status TINYINT NOT NULL CHECK (report_status IN (0, 1, 2)) -- 举报状态，0未举报，1举报审核中，2举报成功
);
```

```go
type Post struct {
	PostID       int       `gorm:"column:post_id;primaryKey;autoIncrement"`                          // 帖子ID，主键，自增
	SenderID     int       `gorm:"column:sender_id;not null"`                                       // 发送者ID，不能为空
	Content      string    `gorm:"column:content;not null"`                                         // 帖子内容，不能为空
	CreatedAt    time.Time `gorm:"column:created_at;default:CURRENT_TIMESTAMP"`                     // 创建时间，默认当前时间
	ReportStatus int       `gorm:"column:report_status;not null;check:report_status IN (0, 1, 2)"` // 举报状态，0未举报，1举报审核中，2举报成功
    ScheduledAt    time.Time `gorm:"column:scheduled_time;default:CURRENT_TIMESTAMP"`                     // 定时发送时间，默认当前时间
}
```

##### 举报表

举报ID

帖子ID

举报者ID

举报原因

创建时间

```sql
CREATE TABLE report (
    report_id INT(8) PRIMARY KEY AUTO_INCREMENT,       -- 举报ID，自增主键
    post_id INT(8) NOT NULL,                           -- 帖子ID，不能为空
    reporter_id INT(8) NOT NULL,                       -- 举报者ID，不能为空
    reason VARCHAR(255) NOT NULL,                      -- 举报原因，不能为空
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP    -- 创建时间，默认当前时间
);
```

```go
type Report struct {
	ReportID   int       `gorm:"column:report_id;primaryKey;autoIncrement"`    // 举报ID，自增主键
	PostID     int       `gorm:"column:post_id;not null"`                      // 帖子ID，不能为空
	ReporterID  int       `gorm:"column:reporter_id;not null"`                  // 举报者ID，不能为空
	Reason     string    `gorm:"column:reason;size:255;not null"`             // 举报原因，不能为空
	CreatedAt  time.Time `gorm:"column:created_at;autoCreateTime"`             // 创建时间，自动设置为当前时间
}
```

##### 举报垃圾箱

举报ID

帖子ID

举报者ID

举报原因

处理结果：0 举报失败，1举报成功

创建时间

```sql
CREATE TABLE trash (
    report_id INT(8) PRIMARY KEY AUTO_INCREMENT,               -- 举报ID，主键，自增
    post_id INT(8) NOT NULL,                                   -- 帖子ID，不能为空
    reporter_id INT(8) NOT NULL,                               -- 举报者ID，不能为空
    reason VARCHAR(255) NOT NULL,                              -- 举报原因，不能为空
    status TINYINT NOT NULL CHECK (status IN (0, 1)),         -- 处理结果：0 为举报失败，1 为举报成功
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP             -- 创建时间，默认当前时间
);
```

```go
type Trash struct {
	ReportID    int       `gorm:"column:report_id;primaryKey;autoIncrement"`                 // 举报ID，主键，自增
	PostID      int       `gorm:"column:post_id;not null"`                                   // 帖子ID，不能为空
	ReporterID  int       `gorm:"column:reporter_id;not null"`                               // 举报者ID，不能为空
	Reason      string    `gorm:"column:reason;size:255;not null"`                          // 举报原因，不能为空
	Status      int       `gorm:"column:status;not null;check:status IN (0, 1)"`           // 处理结果：0 为举报失败，1 为举报成功
	CreatedAt   time.Time `gorm:"column:created_at;autoCreateTime"`                         // 创建时间，默认当前时间
}
```

##### 黑名单

用户ID

黑名ID

黑名Username

创建时间

```sql
CREATE TABLE blacklist (
    black_id INT(8) PRIMARY KEY AUTO_INCREMENT,                 -- 黑名ID，主键，自增
    user_id INT(8) NOT NULL,                                   -- 用户ID，不能为空
    blocked_id INT(8) NOT NULL,                                 -- 被拉黑的用户ID，不能为空
    blocked_username VARCHAR(20) NOT NULL,                     -- 被拉黑的用户名，不能为空
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP             -- 创建时间，默认当前时间
);
```

```go
type Blacklist struct {
	BlackID        int       `gorm:"column:black_id;primaryKey;autoIncrement"`          // 黑名单ID，主键，自增
	UserID         int       `gorm:"column:user_id;not null"`                           // 用户ID，不能为空
	BlockedID      int       `gorm:"column:blocked_id;not null"`                        // 被拉黑的用户ID，不能为空
	BlockedUsername string   `gorm:"column:blocked_username;size:20;not null"`          // 被拉黑的用户名，不能为空
	CreatedAt      time.Time `gorm:"column:created_at;default:CURRENT_TIMESTAMP"`       // 创建时间，默认当前时间
}
```

