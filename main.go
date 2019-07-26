package main

import (
	"fmt"
	"github.com/richguo0615/mini-authsys/conf"
	"github.com/richguo0615/mini-authsys/router"
)

func main() {
	fmt.Println("execute main.go")
	conf.InitDB()
	router.InitRoute()
}
