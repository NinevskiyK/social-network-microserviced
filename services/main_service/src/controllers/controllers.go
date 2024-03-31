package controllers

import (
	"main_service/db_utils"
	"main_service/models"
	"main_service/utils"

	"github.com/gin-gonic/gin"
)

// Post /user/login
// Logs user into the system
func LoginUser(c *gin.Context) {
	var creds models.Credentionals
	if err := c.BindJSON(&creds); err != nil {
		c.JSON(400, gin.H{"error": "no creds"})
		return
	}

	is_right_creds, err := db_utils.CheckPassword(creds)
	if err != nil || !is_right_creds {
		c.JSON(400, gin.H{"error": "wrong creds"})
		return
	}

	id, err := db_utils.GetId(creds.UserName)
	if err != nil {
		c.JSON(500, gin.H{"error": "internal error"})
		return
	}

	token, err := utils.CreateJWT(id)
	if err != nil {
		c.JSON(500, gin.H{"error": "internal error"})
		return
	}

	c.SetCookie("token", token, 60*60, "/", "localhost", true, false)
	c.JSON(200, gin.H{"status": "OK"})
}

// Post /user/register
// Register a new user
func Register(c *gin.Context) {
	var creds models.Credentionals
	if err := c.BindJSON(&creds); err != nil {
		c.JSON(400, gin.H{"error": "no creds"})
		return
	}

	is_login_used, err := db_utils.IsLoginAlreadyInUse(creds.UserName)
	if err != nil {
		c.JSON(500, gin.H{"error": "internal error"})
	}
	if is_login_used {
		c.JSON(400, gin.H{"error": "login is already in use"})
		return
	}

	db_utils.AddNewUser(creds)
}

// Put /user/update_me
// Updates user info
func UpdateMe(c *gin.Context) {
	id := c.GetString("id")
	var info models.UserInfo
	if err := c.BindJSON(&info); err != nil {
		c.JSON(400, gin.H{"error": "no info"})
		return
	}
	err := db_utils.UpdateInfo(id, info)
	if err != nil {
		c.JSON(500, gin.H{"error": "internal error"})
		return
	}
	c.JSON(200, gin.H{"status": "OK"})
}
