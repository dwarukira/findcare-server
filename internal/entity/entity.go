package entity

import (
	"fmt"
	"time"

	"github.com/dwarukira/findcare/internal/event"
	"github.com/jinzhu/gorm"
)

var log = event.Log

func logError(result *gorm.DB) {
	if result.Error != nil {
		log.Error(result.Error.Error())
	}
}

type Types map[string]interface{}

// List of database entities and their table names.
var Entities = Types{
	"errors":            &Error{},
	"providers":         &Provider{},
	"services":          &Service{},
	"provider_services": &ProviderService{},
}

type RowCount struct {
	Count int
}

// WaitForMigration waits for the database migration to be successful.
func (list Types) WaitForMigration() {
	attempts := 100
	for name := range list {
		for i := 0; i <= attempts; i++ {
			count := RowCount{}
			if err := Db().Raw(fmt.Sprintf("SELECT COUNT(*) AS count FROM %s", name)).Scan(&count).Error; err == nil {
				// log.Debugf("entity: table %s migrated", name)
				break
			} else {
				log.Debugf("entity: wait for migration %s (%s)", err.Error(), name)
			}

			if i == attempts {
				panic("migration failed")
			}

			time.Sleep(50 * time.Millisecond)
		}
	}
}

// Truncate removes all data from tables without dropping them.
func (list Types) Truncate() {
	for name := range list {
		if err := Db().Exec(fmt.Sprintf("DELETE FROM %s WHERE 1", name)).Error; err == nil {
			// log.Debugf("entity: removed all data from %s", name)
			break
		} else if err.Error() != "record not found" {
			log.Debugf("entity: %s in %s", err, name)
		}
	}
}

// Drop migrates all database tables of registered entities.
func (list Types) Migrate() {
	for _, entity := range list {
		if err := UnscopedDb().AutoMigrate(entity).Error; err != nil {
			log.Debugf("entity: migrate %s (waiting 1s)", err.Error())

			time.Sleep(time.Second)

			if err := UnscopedDb().AutoMigrate(entity).Error; err != nil {
				panic(err)
			}
		}
	}
}

// Drop drops all database tables of registered entities.
func (list Types) Drop() {
	for _, entity := range list {
		if err := UnscopedDb().DropTableIfExists(entity).Error; err != nil {
			panic(err)
		}
	}
}

func CreateDefaultFixtures() {
	CreateDefaultProviders()
	CreateDefaultServices()
}

// MigrateDb creates all tables and inserts default entities as needed.
func MigrateDb() {
	Entities.Migrate()
	Entities.WaitForMigration()

	CreateDefaultFixtures()
}

// InitTestDb connects to and completely initializes the test database incl fixtures.
func InitTestDb(driver, dsn string) *Gorm {
	if HasDbProvider() {
		return nil
	}

	if driver == "test" || driver == "sqlite" || driver == "" || dsn == "" {
		driver = "sqlite3"
		dsn = ".test.db"
	}

	log.Infof("initializing %s test db in %s", driver, dsn)

	db := &Gorm{
		Driver: driver,
		Dsn:    dsn,
	}

	SetDbProvider(db)
	// ResetTestFixtures()

	return db
}
