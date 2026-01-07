package textStore

// Data Access Object Text
// Version: 0.3.0
// Updated on: 2025-12-31

import (
	"context"

	"github.com/mt1976/frantic-core/dao/audit"
	"github.com/mt1976/frantic-core/dao/database"
	"github.com/mt1976/frantic-core/jobs"
	"github.com/mt1976/frantic-core/logHandler"
	"github.com/mt1976/frantic-core/timing"
)

// Worker is a job that is scheduled to run at a predefined interval
func Worker(j jobs.Job, db *database.DB) {
	clock := timing.Start(jobs.CodedName(j), "Initialise", j.Description())
	oldDB := activeDB
	dbSwitched := false
	// Overide the default database connection if one is passed

	if db != nil {
		if activeDB.Name != db.Name {
			logHandler.EventLogger.Printf("Switching to %v.db", db.Name)
			activeDB = db
			dbSwitched = true
		}
	}

	TextJobProcessor(j)

	if dbSwitched {
		logHandler.EventLogger.Printf("Switching back to %v.db from %v.db", oldDB.Name, activeDB.Name)
		activeDB = oldDB
	}
	clock.Stop(1)
}

// TextJobProcessor processes jobs related to the TextStore domain entity.
// This function is triggered by the job scheduler to perform specific operations on TextStore records.
func TextJobProcessor(j jobs.Job) {
	clock := timing.Start(jobs.CodedName(j), "Process", j.Description())
	count := 0

	//TODO: Add your job processing code here

	// Get all the sessions
	// For each session, check the expiry date
	// If the expiry date is less than now, then delete the session

	TextEntries, err := GetAll()
	if err != nil {
		logHandler.ErrorLogger.Printf("[%v] Error: '%v'", jobs.CodedName(j), err.Error())
		return
	}

	noTextEntries := len(TextEntries)
	if noTextEntries == 0 {
		logHandler.ServiceLogger.Printf("[%v] No %vs to process", jobs.CodedName(j), Domain)
		clock.Stop(0)
		return
	}

	for TextEntryIndex, TextRecord := range TextEntries {
		logHandler.ServiceLogger.Printf("[%v] %v(%v/%v) %v", jobs.CodedName(j), Domain, TextEntryIndex+1, noTextEntries, TextRecord.Raw)
		TextRecord.UpdateWithAction(context.Background(), audit.SERVICE, "Job Processing "+j.Name())
		count++
	}
	clock.Stop(count)
}
