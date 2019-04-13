package config

import (
	"fmt"
	"log"

	//driver mysql database
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

var db *gorm.DB

func connectionMysql(connString string) *gorm.DB {
	db, err := gorm.Open("mysql", connString)
	if err != nil {
		log.Println(err.Error())
		return nil
	}

	err = db.DB().Ping()
	if err != nil {
		log.Println("database not connect")
		return nil
	}

	db.DB().SetMaxIdleConns(10)
	db.DB().SetMaxOpenConns(10)
	return db
}

//GetMySQLDB - get connection mysql
func GetMySQLDB() *gorm.DB {
	return db
}

//InitConnectDB - initial connect database
func InitConnectDB() {
	dbHost := "localhost"
	dbUser := "root"
	dbPass := ""
	dbName := "golang_project"
	connString := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True", dbUser, dbPass, dbHost, dbName)

	db = connectionMysql(connString)
}
