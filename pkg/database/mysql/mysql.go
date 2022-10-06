package mysql

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func Conn() *gorm.DB {
	var (
		host = os.Getenv("SIA_DB_HOST")
		port = os.Getenv("SIA_DB_PORT")
		user = os.Getenv("SIA_DB_USER")
		pass = os.Getenv("SIA_DB_PASS")
		name = os.Getenv("SIA_DB_NAME")
	)
	mysqlConnString := fmt.Sprintf("%s:%s@(%s:%s)/%s", user, pass, host, port, name)

	db, err := gorm.Open(mysql.Open(mysqlConnString), &gorm.Config{})

	log.Printf("Connection string %v", mysqlConnString)
	if err != nil {
		log.Println("Error connecting to database: ", err.Error())
		return nil
	}

	return db
}

func ConnNative() *sql.DB {
	var (
		host = os.Getenv("SIA_DB_HOST")
		port = os.Getenv("SIA_DB_PORT")
		user = os.Getenv("SIA_DB_USER")
		pass = os.Getenv("SIA_DB_PASS")
		name = os.Getenv("SIA_DB_NAME")
	)

	mysqlConnString := fmt.Sprintf("%s:%s@(%s:%s)/%s", user, pass, host, port, name)
	db, err := sql.Open("mysql", mysqlConnString)
	if err != nil {
		panic(err)
	}
	// See "Important settings" section.
	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)

	return db
}
