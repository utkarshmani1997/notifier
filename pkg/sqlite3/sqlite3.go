package sqlite3

import (
	"github.com/davecgh/go-spew/spew"
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
	"github.com/pkg/errors"

	"github.com/utkarshmani1997/notify/pkg/incident"
	"github.com/utkarshmani1997/notify/pkg/logger"
)

var log = logger.Log

// Config holds the configuration used for instantiating a new DataStore.
type Config struct {
	Path string
}

type DataStore struct {
	Db  *gorm.DB
	cfg Config
}

// New returns a DataStore instance with the sql.DB set with the postgres
func New(cfg Config) (DataStore, error) {
	var err error
	var store DataStore

	if cfg.Path == "" {
		err = errors.Errorf(
			"Path must be set (%s)",
			spew.Sdump(cfg))
		return store, err
	}
	store.cfg = cfg
	db, err := gorm.Open("sqlite3", cfg.Path)

	if err != nil {
		err = errors.Wrapf(err,
			"Couldn't open connection to sqlite3 database (%s)",
			spew.Sdump(cfg)) // Sdump returns a string with the passed arguments formatted exactly the same as Dump.
		return store, err
	}

	db.SetLogger(log)

	store.Db = db
	return store, nil
}

func (ds *DataStore) Create(inc *incident.Incident) error {
	db := ds.Db.Create(inc)
	return db.Error
}

func (ds *DataStore) Delete(id string) (uint, error) {
	inc, err := ds.Get(id)
	if err != nil {
		return 0, err
	}

	db := ds.Db.Delete(&inc)
	return inc.ID, db.Error
}

func (ds *DataStore) Get(id string) (incident.Incident, error) {
	var inc incident.Incident
	db := ds.Db.First(&inc, id)
	return inc, db.Error
}

func (ds *DataStore) Update(id, issue string) (incident.Incident, error) {
	inc, err := ds.Get(id)
	if err != nil {
		return inc, err
	}
	db := ds.Db.Model(&inc).Update("Report", issue)
	return inc, db.Error
}
