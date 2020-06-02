package main

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	// initialize sqlite3
	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm/dialects/sqlite"

	"github.com/utkarshmani1997/notify/pkg/logger"
	"github.com/utkarshmani1997/notify/pkg/router"
	"github.com/utkarshmani1997/notify/pkg/sqlite3"
)

var log = logger.Log

func main() {
	db, route := router.InitAPIServerRoutes(sqlite3.Config{Path: "incident.report"})
	defer db.Db.Close()
	route.Use(gin.Recovery())
	server := &http.Server{
		Addr:         ":8080",
		Handler:      route,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	go func() {
		server.ListenAndServe()
	}()

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, os.Kill, syscall.SIGTERM)
	interruptSignal := <-interrupt
	switch interruptSignal {
	case os.Kill:
		log.Warning("Got SIGKILL...")
	case os.Interrupt:
		log.Warning("Got SIGINT...")
	case syscall.SIGTERM:
		log.Warning("Got SIGTERM...")
	}

	log.Info("The service is shutting down...")
	server.Shutdown(context.Background())
	log.Info("Shut down is done")
}
