package model

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"log"
)

var (
	userName = "root"
	password = "fenghaojiang97"
	database = "douban"
)

func init() {
	var DB *gorm.DB
	var err error
	DB, err = gorm.Open("mysql", fmt.Sprintf("%s:%s@tcp(:3306)/%s?charset=utf8&parseTime=True&loc=Local", username, password, dbName))
	if err != nil {
		log.Fatalf(" gorm.Open.err: %v", err)
	}
	fmt.Println("Init successfully")
	DB.SingularTable(true)
	gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
		return "sp_" + defaultTableName
	}
}
