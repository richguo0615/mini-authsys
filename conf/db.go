package conf

import (
	"fmt"
	"github.com/go-redis/redis"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/richguo0615/mini-authsys/model/db"
)

var DB *gorm.DB
var RedisClient *redis.Client

func InitDB() {
	var host string = "localhost"
	var port string = "5432"
	var user string = "postgres"
	var dbName string = "auth_system"
	var password string = "qianostgres"

	sysDB, err := gorm.Open("postgres", fmt.Sprint("host=", host, " port=", port, " user=", user, " dbname=", dbName, " password=", password, " sslmode=disable"))
	if err != nil {
		panic(err)
	}

	sysDB.AutoMigrate(
		&db.User{},
		&db.Transaction{},
	)

	DB = sysDB
	fmt.Println("db init ...")
}

func InitRedis() {
	RedisClient = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	pong, err := RedisClient.Ping().Result()
	fmt.Println("redis init: ", pong, err, "...")
}
