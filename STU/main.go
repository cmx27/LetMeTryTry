package main

import (
	"STU/app/midwares"
	"STU/config/database"
	"STU/config/router"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	database.Init()
	r := gin.Default()
	r.NoMethod(midwares.HandleNotFound)
	r.NoRoute(midwares.HandleNotFound)
	router.Init(r)
	router := gin.Default()
	router.SetTrustedProxies([]string{"127.0.0.1"})

	err := r.Run()
	if err != nil {
		log.Fatal("Server start failed: ", err)
	}

}
