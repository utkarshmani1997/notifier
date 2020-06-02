package incident

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"github.com/utkarshmani1997/notify/pkg/logger"
)

var log = logger.Log

type Controller struct {
	Store IncidentStore // This gives extensibility and testability
}

// HTTP Post - /incident
func (handler Controller) Create(c *gin.Context) {
	var inc createIncident
	// Decode the incoming json data to Incident struct
	if err := c.ShouldBindJSON(&inc); err != nil {
		log.Error("Failed to decode create incident request body, err: ", err)
		c.String(http.StatusBadRequest, errors.Wrap(err, "Unable to decode JSON request body").Error())
		return
	}
	// Persistence
	id, err := handler.Store.Create(inc.Email, inc.Issue)
	if err != nil {
		log.Error(err)
		c.String(http.StatusInternalServerError, errors.Wrap(err, "Error on inserting incident").Error())
		return
	}
	c.JSON(http.StatusCreated, id)
}

// HTTP Get - /incident
func (handler Controller) Get(c *gin.Context) {
	id := c.Param("id")
	// Persistence
	newIncident, err := handler.Store.Get(id)
	if err != nil {
		log.Error(err)
		c.String(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, newIncident)
}

// HTTP Update - /incident
func (handler Controller) Update(c *gin.Context) {
	var inc updateIncident
	// Decode the incoming json data to Incident struct
	id := c.Param("id")
	if err := c.ShouldBindJSON(&inc); err != nil {
		log.Error("Failed to decode update incident request body, err: ", err)
		c.String(http.StatusBadRequest, errors.Wrap(err, "Unable to decode JSON request body").Error())
		return
	}

	// Persistence
	newIncident, err := handler.Store.Update(id, inc.Issue)
	if err != nil {
		log.Error(err)
		c.String(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, newIncident)
}

// HTTP Delete - /incident
func (handler Controller) Delete(c *gin.Context) {
	id := c.Param("id")
	// Persistence
	uid, err := handler.Store.Delete(id)
	if err != nil {
		log.Error(err)
		c.String(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, uid)
}
