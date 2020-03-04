package main

import (
	"github.com/gin-gonic/gin"
	"A-module/handler"
	"DAgent/src/routers"
	"A-module/setting"
	"A-module/log"
	"net/http"
	"time"
)

// ToDo: Use Lorgus => https://github.com/sirupsen/logrus
// Log level Redefine

func init() {
	setting.LoadConfig()
	gin.SetMode(gin.DebugMode)
	log.SetDebugMode()
}

func main() {
	go handler.ConnectToIBoFOS()
	startServer()
}

func startServer() {
	routersInit := routers.InitRouter()

	server := &http.Server{
		Addr:           ":" + setting.Config.Server.Dagent.Port,
		Handler:        routersInit,
		ReadTimeout:    30 * time.Second,
		WriteTimeout:   30 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	server.ListenAndServe()
}