package config

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB = Connect()

type DBConfig struct {
	Host     string
	Port     int
	UserName string
	Password string
	DBName   string
}

func BuildDBConfig() *DBConfig {
	config := DBConfig{
		Host:     "go-dev2.cw4aowryfr4j.ap-southeast-1.rds.amazonaws.com",
		Port:     3306,
		UserName: "root",
		Password: "punchxxv25",
		DBName:   "demo",
	}

	return &config
}

func DBUrl(dbConfig *DBConfig) string {
	return fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local",
		dbConfig.UserName,
		dbConfig.Password,
		dbConfig.Host,
		dbConfig.Port,
		dbConfig.DBName,
	)

}

func Connect() *gorm.DB {
	db, err := gorm.Open(mysql.Open(DBUrl(BuildDBConfig())), &gorm.Config{})

	if err != nil {
		fmt.Println("Cloud not connect database")
	}

	return db
}
