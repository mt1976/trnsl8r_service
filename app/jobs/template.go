package jobs

import (
	"github.com/mt1976/frantic-core/dao/database"
	"github.com/mt1976/frantic-core/jobs"
	"github.com/mt1976/frantic-core/timing"
	"github.com/mt1976/trnsl8r_service/app/business/domains"
	"github.com/mt1976/trnsl8r_service/app/business/translate"
)

type template struct {
}

func (job *template) Run() error {
	jobNotifications()
	jobs.NextRun(job.Name(), job.Schedule())
	return nil
}

func (job *template) Service() func() {
	return func() {
		job.Run()
	}
}

func (job *template) Schedule() string {
	return "10 7 * * *"
}

func (job *template) Name() string {
	return translate.Get("Template Job", "")
}

func jobNotifications() {
	// Do something every day at midnight

	j := timing.Start(domains.JOBS.String(), "Send", "Service")

	j.Stop(0)
}

func (job *template) AddDatabaseAccessFunctions(fn func() ([]*database.DB, error)) {
	// do nothing
}

func (t *template) Description() string {
	return "Template Job"
}
