package mysql

import (
	"fmt"
	"log"
	"os"

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

	if err != nil {
		log.Println("Error connecting to database: ", err.Error())
		return nil
	}

	return db
}
