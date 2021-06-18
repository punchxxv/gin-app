package controllers

import (
	"gin-app/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetAllUsers(c *gin.Context) {
	user := models.GetAllUsers()
	c.JSON(http.StatusOK, user)
}

func GetUserById(c *gin.Context) {
	id := c.Params.ByName("id")
	user := models.GetUserById(id)
	c.JSON(http.StatusOK, user)
}

func CreateUser(c *gin.Context) {
	err := models.CreateUser(c)
	c.JSON(http.StatusOK, err)
}

func UpdateUser(c *gin.Context) {
	id := c.Params.ByName("id")
	err := models.UpdateUser(c, id)
	c.JSON(http.StatusOK, err)
}

func DeleteUser(c *gin.Context) {
	id := c.Params.ByName("id")
	err := models.DeleteUser(id)
	c.JSON(http.StatusOK, err)
}
