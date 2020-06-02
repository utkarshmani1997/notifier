package router

import (
	"github.com/gin-gonic/gin"
	"github.com/utkarshmani1997/notifier/pkg/sqlite3"
)

// InitRoutes initializes the routes for APIServer
func InitAPIServerRoutes(cfg sqlite3.Config) (sqlite3.DataStore, *gin.Engine) {
	router := gin.New()
	router.Use(Logger())
	db, err := sqlite3.New(cfg)
	if err != nil {
		panic(err)
	}
	router = SetIncidentRoutes(router, db)

	return db, router
}
