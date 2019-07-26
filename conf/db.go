package conf

import (
	"fmt"
	"github.com/richguo0615/mini-authsys/model/db"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var DB *gorm.DB

func InitDB() {
	var host string = "localhost"
	var port string = "5432"
	var user string = "postgres"
	var dbName string = "auth_system"
	var password string = "qianostgres"

	sysDB, err := gorm.Open("postgres", fmt.Sprint("host=",host," port=",port," user=",user," dbname=",dbName," password=",password, " sslmode=disable"))
	if err != nil {
		panic(err)
	}

	sysDB.AutoMigrate(
		&db.User{},
		&db.Transaction{},
	)

	DB = sysDB
}
