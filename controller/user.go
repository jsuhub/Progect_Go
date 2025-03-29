package controller

import (
	"Progect/config"
	"Progect/model"
	"Progect/utils"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func PassUser(c *gin.Context) {
	var user model.User
	if err := c.ShouldBind(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": false,
			"error":  err,
		})
	}
	log.Println(user)
	password := user.Password
	result := config.DB.Where("name = ?", user.Name).First(&user)
	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": false,
			"error":  result.Error,
		})
	}
	if user.Password == password {
		token, err := utils.GenerateToken(user.Name)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status": false,
				"error":  err,
			})
		}
		c.JSON(http.StatusOK, gin.H{
			"status": true,
			"user":   user,
			"token":  token,
		})
	} else {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": "密码错误",
		})
	}

}

func CreateUser(c *gin.Context) {
	log.Println("接收到了请求")
	var user model.User
	if err := c.ShouldBindJSON(&user); err != nil {
		log.Println(user)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	log.Println(user)
	result := config.DB.Create(&user)
	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": false,
			"error":  result.Error,
		})
	}
	token, err := utils.GenerateToken(user.Name)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": false,
			"error":  err,
		})
	}
	c.JSON(http.StatusOK, gin.H{
		"status":  true,
		"message": "用户添加成功",
		"token":   token,
	})
}

func GetUsers(c *gin.Context) {
	var users []model.User
	result := config.DB.Find(&users) // 查询 users 表中所有行
	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": false,
			"error":  result.Error,
		})
	}
	c.JSON(200, gin.H{
		"status": true,
		"data":   users,
	})

}
