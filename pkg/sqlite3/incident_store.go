package sqlite3

import (
	"github.com/pkg/errors"
	"github.com/utkarshmani1997/notifier/pkg/incident"
)

var (
	// compile time check if IncidentStore implements the contract
	_ incident.IncidentStore = IncidentStore{}
)

// IncidentStore provides persistence logic for "incident" table
type IncidentStore struct {
	Store DataStore
}

// Create creates a new Incident
func (is IncidentStore) Create(email, issue string) (uint, error) {
	inc := incident.New(incident.WithEmail(email), incident.WithReport(issue))
	if err := is.Store.Create(inc); err != nil {
		return 0, errors.Wrap(err, "Failed to create incident")
	}
	log.Infof("Created an incident with id: %d", inc.ID)
	return inc.ID, nil
}

// Delete deletes the Incident with given id
func (is IncidentStore) Delete(id string) (uint, error) {
	uid, err := is.Store.Delete(id)
	if err != nil {
		return 0, errors.Wrap(err, "Failed to delete incident")
	}
	log.Infof("Deleted an incident with id: %d", uid)
	return uid, nil
}

// Update updates issue of the Incident for a given id
func (is IncidentStore) Update(id string, issue string) (incident.Incident, error) {
	inc, err := is.Store.Update(id, issue)
	if err != nil {
		return inc, errors.Wrap(err, "Failed to update incident")
	}
	log.Infof("Updated an incident with id: %d", inc.ID)
	return inc, nil
}

// Get returns the Incident
func (is IncidentStore) Get(id string) (incident.Incident, error) {
	// ToDo: Write the code here
	inc, err := is.Store.Get(id)
	if err != nil {
		return inc, errors.Wrap(err, "Failed to get incident")
	}
	log.Infof("Get an incident with id: %d", inc.ID)
	return inc, nil
}
