package jobs

import (
	"fmt"

	"github.com/mt1976/frantic-core/dao/maintenance"
	"github.com/mt1976/frantic-core/jobs"
	"github.com/mt1976/frantic-core/logHandler"
	"github.com/mt1976/trnsl8r_service/app/business/domains"
	"github.com/mt1976/trnsl8r_service/app/dao/textstore"
	cron3 "github.com/robfig/cron/v3"
)

var Template jobs.Job = &template{} // This is a template for other jobs.

var DatabaseBackup jobs.Job = &maintenance.DatabaseBackupJob{}
var DatabasePrune jobs.Job = &maintenance.DatabaseBackupCleanerJob{}
var LocaleUpdate jobs.Job = &localeUpdate{}
var scheduledTasks *cron3.Cron

var domain = domains.JOBS

func init() {
	scheduledTasks = cron3.New()
}

func Start() {
	// Start the job
	logHandler.ServiceLogger.Printf("[%v] Queue - Starting", domain.String())
	// Add the functions to the jobs
	DatabaseBackup.AddFunction(textstore.GetDB())
	fmt.Printf("DatabaseBackup: %+v\n", DatabaseBackup)
	// Database Backup
	Schedule(DatabaseBackup)
	// Prune the archive of backups
	Schedule(DatabasePrune)
	// Start all the background jobs
	scheduledTasks.Start()
	logHandler.ServiceLogger.Printf("[%v] Queue - Started", domain.String())
}

func Schedule(j jobs.Job) {
	// Start the job
	id, err := scheduledTasks.AddFunc(j.Schedule(), j.Service())
	if err != nil {
		logHandler.ErrorLogger.Printf("[%v] Job [%v] Schedule [%v] Error [%v]", domain.String(), j.Name(), j.Schedule(), err.Error())
		return
	}
	logHandler.ServiceLogger.Printf("[%v] Job [%v] Scheduled [%v] ID [%v]", domain.String(), j.Name(), j.Schedule(), id)
	jobs.Announce(j.Name(), "Scheduled")
	jobs.NextRun(j.Name(), j.Schedule())
}
