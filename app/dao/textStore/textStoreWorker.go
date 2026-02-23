// Data Access Object for the TextStore table
// Template Version: 0.6.00 - 2026-02-14
// Generated 
// Date: 23/02/2026 & 12:36
// Who : matttownsend (orion)

package textStore

import (
	"github.com/mt1976/frantic-amphora/dao/database"
	"github.com/mt1976/frantic-amphora/jobs"
	"github.com/mt1976/frantic-core/logHandler"
	"github.com/mt1976/frantic-core/timing"
)

// Worker is a job that is scheduled to run at a predefined interval.
func Worker(j jobs.Job, db *database.DB) {
	clock := timing.Start(jobs.CodedName(j), "Initialise", j.Description())
	oldDB := activeDBConnection
	dbSwitched := false

	if db != nil {
		if activeDBConnection.Name != db.Name {
			logHandler.Event.Printf("Switching to %v.db", db.Name)
			activeDBConnection = db
			dbSwitched = true
		}
	}

	if worker != nil {
		worker(jobs.CodedName(j), j.Description())
	}

	if dbSwitched {
		logHandler.Event.Printf("Switching back to %v.db from %v.db", oldDB.Name, activeDBConnection.Name)
		activeDBConnection = oldDB
	}
	clock.Stop(1)
}
