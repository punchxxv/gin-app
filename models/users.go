package models

import (
	"gin-app/config"

	"github.com/gin-gonic/gin"
)

type Users struct {
	ID        int
	Username  string
	Password  string
	FirstName string
	LastName  string
}

func GetAllUsers() []Users {
	var user []Users
	config.DB.Raw("select id, username, password, first_name, last_name from users").Scan(&user)

	return user
}

func GetUserById(id string) Users {
	var user Users
	config.DB.Raw("select id, username, password, first_name, last_name from users where id = ?", id).First(&user)

	return user
}

func CreateUser(c *gin.Context) string {
	var user Users
	c.BindJSON(&user)
	config.DB.Create(&user)

	return "Success"
}

func UpdateUser(c *gin.Context, id string) string {
	var user Users
	user = GetUserById(id)
	c.BindJSON(&user)
	config.DB.Save(&user)

	return "Success"
}

func DeleteUser(id string) string {
	config.DB.Delete(&Users{}, id)

	return "Success"
}
