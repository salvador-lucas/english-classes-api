package models

import (
	"errors"
	"fmt"
	"os"
	"time"

	//mysql connection
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

func ObtainDbConnection() (*gorm.DB, error) {
	user := os.Getenv("DB_USERNAME")
	pass := os.Getenv("DB_PASSWORD")
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")

	if user == "" || pass == "" || host == "" {
		return nil, errors.New("invalid user/pass/host")
	}

	connectionURLgo := fmt.Sprintf("%s:%s@(%s:%s)/%s?parseTime=true", user, pass, host, port, dbName)

	db, err := gorm.Open("mysql", connectionURLgo)
	if err != nil {
		return nil, err
	}

	db.AutoMigrate(&Student{})
	db.AutoMigrate(&Classes{})
	// db.Model(&Classes{}).AddForeignKey("student_id", "students(id)", "RESTRICT", "RESTRICT")

	db.DB().SetMaxIdleConns(10)
	db.DB().SetMaxOpenConns(100)
	db.DB().SetConnMaxLifetime(time.Hour)

	return db, nil
}
