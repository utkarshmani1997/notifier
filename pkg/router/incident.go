package router

import (
	"github.com/gin-gonic/gin"
	"github.com/utkarshmani1997/notifier/pkg/incident"
	"github.com/utkarshmani1997/notifier/pkg/sqlite3"
)

// SetIncidentRoutes sets the routes for CRUD operations
func SetIncidentRoutes(engine *gin.Engine, db sqlite3.DataStore) *gin.Engine {

	incidentStore := sqlite3.IncidentStore{Store: db}
	// create table incident
	db.Db.AutoMigrate(&incident.Incident{})
	incidentController := incident.Controller{Store: incidentStore}
	engine.POST("/incident", incidentController.Create)
	engine.GET("/incident/:id", incidentController.Get)
	engine.PATCH("/incident/:id", incidentController.Update)
	engine.DELETE("/incident/:id", incidentController.Delete)

	return engine
}
