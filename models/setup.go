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
	user := os.Getenv("DB_USER")
	pass := os.Getenv("DB_PASS")
	host := os.Getenv("DB_HOST")

	if user == "" || pass == "" || host == "" {
		return nil, errors.New("invalid user/pass/host")
	}

	connectionURLgo := fmt.Sprintf("%s:%s@(%s)/solar_system_db?parseTime=true", user, pass, host)

	db, err := gorm.Open("mysql", connectionURLgo)
	if err != nil {
		return nil, err
	}

	db.AutoMigrate(&Student{})
	db.Model(&Student{}).AddForeignKey("solar_system_id", "solar_systems(id)", "RESTRICT", "RESTRICT")

	db.DB().SetMaxIdleConns(10)
	db.DB().SetMaxOpenConns(100)
	db.DB().SetConnMaxLifetime(time.Hour)

	return db, nil
}
