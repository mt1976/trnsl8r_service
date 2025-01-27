package jobs

import (
	"github.com/mt1976/trnsl8r_service/app/business/domains"
	"github.com/mt1976/trnsl8r_service/app/support/config"
	"github.com/mt1976/trnsl8r_service/app/support/logger"
	cron3 "github.com/robfig/cron/v3"
)

var Template Job = template{} // This is a template for other jobs.

var DatabaseBackup Job = databaseBackup{}
var DatabasePrune Job = databasePrune{}
var scheduledTasks *cron3.Cron
var cfg *config.Configuration
var domain = domains.JOBS

func init() {
	scheduledTasks = cron3.New()
	cfg = config.Get()
}

func Start() {
	// Start the job
	logger.ServiceLogger.Printf("[%v] Queue - Starting", domain.String())
	// Database Backup
	Schedule(DatabaseBackup)
	// Prune the archive of backups
	Schedule(DatabasePrune)
	// Start all the background jobs
	scheduledTasks.Start()
	logger.ServiceLogger.Printf("[%v] Queue - Started", domain.String())
}

func Schedule(j Job) {
	// Start the job
	scheduledTasks.AddFunc(j.Schedule(), j.Service())
	announceJob(j, "Scheduled")
	NextRun(j)
}
