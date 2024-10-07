package services

import (
	"myapp/database"
	"myapp/models"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Register(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	// username�����пո�
	if strings.Contains(user.Username, " ") {
		c.JSON(http.StatusOK, gin.H{"code": 200502, "data": nil, "msg": "�û��������пո�"})
		return
	}
	// ���볤�ȴ���8λС��16λ
	if len(user.Password) <= 8 || len(user.Password) >= 16 {
		c.JSON(http.StatusOK, gin.H{"code": 200503, "data": nil, "msg": "���볤�ȱ�����9-15λ"})
		return
	}

	// �û������ȴ���3λС��20λ
	if len(user.Username) <= 3 || len(user.Username) >= 20 {
		c.JSON(http.StatusOK, gin.H{"code": 200504, "data": nil, "msg": "�û������ȱ�����4-19λ"})
		return
	}

	// �û����ܴ���
	var users models.Accounts
	if err := database.Where("username = ?", user.Username).First(&users).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			// ����ע��
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"code": 500, "msg": "Internal server error"})
			return
		}
	} else {
		// �û��Ѵ���
		c.JSON(http.StatusOK, gin.H{"code": 200505, "data": nil, "msg": "�û����Ѵ���"})
		return
	}
	err := database.SaveUser(&users)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save user"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "User registered successfully"})
}

func Login(c *gin.Context) {
	var loginData struct {
		Username string
		Password string
	}
	if err := c.ShouldBindJSON(&loginData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	user, err := database.GetUserByUsername(loginData.Username)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not found"})
		return
	}

	var account models.Accounts
	if user.Password != database.QueryRow("SELECT * FROM account WHERE Name =? LIMIT 1", "Password").Scan(&account.Password) {
		c.JSON(http.StatusOK, gin.H{"code": 200507, "data": nil, "msg": "Invalid password"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Login successful", "user": user})
}
