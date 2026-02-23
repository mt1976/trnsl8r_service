package jobs

import (
	"github.com/mt1976/frantic-amphora/dao/maintenance"
	"github.com/mt1976/frantic-amphora/jobs"
	"github.com/mt1976/frantic-core/commonConfig"
	"github.com/mt1976/frantic-core/logHandler"
	"github.com/mt1976/trnsl8r_service/app/business/domains"
	"github.com/mt1976/trnsl8r_service/app/dao/textStore"
)

var Template jobs.Job = &template{} // This is a template for other jobs.

var (
	DatabaseBackup jobs.Job = &maintenance.DatabaseBackupJob{}
	DatabasePrune  jobs.Job = &maintenance.DatabaseBackupCleanerJob{}
	LocaleUpdate   jobs.Job = &localeUpdate{}
)

var domain = domains.JOBS

func init() {
	cfg := commonConfig.Get()
	err := jobs.Initialise(cfg)
	if err != nil {
		logHandler.Service.Fatal(err.Error())
	}
}

func Start() {
	// Start the job
	logHandler.Service.Printf("[%v] Queue - Starting", domain.String())
	// Add the functions to the jobs, one for each table/domain that required a backup
	DatabaseBackup.AddDatabaseAccessFunctions(textStore.GetDatabaseConnections())
	// Database Backup§
	jobs.AddJobToScheduler(DatabaseBackup)
	// Prune the archive of backups
	jobs.AddJobToScheduler(DatabasePrune)
	// Start all the background jobs
	jobs.StartScheduler()
	logHandler.Service.Printf("[%v] Queue - Started", domain.String())
}
