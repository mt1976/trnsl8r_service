package jobs

import (
	"strings"
	"time"

	"github.com/mt1976/trnsl8r_service/app/business/translation"
	"github.com/mt1976/trnsl8r_service/app/dao/support/database"
	"github.com/mt1976/trnsl8r_service/app/support/date"
	"github.com/mt1976/trnsl8r_service/app/support/io"
	"github.com/mt1976/trnsl8r_service/app/support/logger"
	"github.com/mt1976/trnsl8r_service/app/support/paths"
	"github.com/mt1976/trnsl8r_service/app/support/timing"
)

type databaseBackup struct {
}

func (job databaseBackup) Run() error {
	//name := "Database Backup"
	announceJob(job, "Started")
	j := timing.Start(job.Name(), "Backup", "All")

	dateTime := time.Now().Format(date.Format.BackupFolder)
	logger.ServiceLogger.Printf("[%v] [BACKUP] Date=[%v]", domain.String(), dateTime)

	destPath := paths.Backups().String() + paths.Seperator() + dateTime
	fullBackupPath := paths.Application().String() + destPath

	//create a folder
	err := io.MkDir(fullBackupPath)
	if err != nil {
		logger.ErrorLogger.Printf("[%v] [%v] Error=[%v]", domain.String(), strings.ToUpper(job.Name()), err.Error())
	}

	// texts.Backup(destPath)
	// status.Backup(destPath)
	// users.Backup(destPath)
	// settings.Backup(destPath)
	//hosts.Backup(destPath)
	database.Backup(destPath)

	j.Stop(6)
	//b.lastRun = time.Now()

	announceJob(job, "Completed")
	NextRun(job)
	return nil
}

func (job databaseBackup) Service() func() {
	return func() {
		job.Run()
	}
}

func (job databaseBackup) Schedule() string {
	return "55 11 * * *"
}

func (job databaseBackup) Name() string {
	return translation.Get("Scheduled Database Backup")
}
