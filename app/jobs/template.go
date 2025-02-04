package jobs

import (
	"github.com/mt1976/frantic-plum/timing"
	"github.com/mt1976/trnsl8r_service/app/business/domains"
	"github.com/mt1976/trnsl8r_service/app/business/translation"
)

type template struct {
}

func (job template) Run() error {
	jobNotifications()
	NextRun(job)
	return nil
}

func (job template) Service() func() {
	return func() {
		job.Run()
	}
}

func (job template) Schedule() string {
	return "10 7 * * *"
}

func (job template) Name() string {
	return translation.Get("Template Job", "")
}

func jobNotifications() {
	// Do something every day at midnight

	j := timing.Start(domains.JOBS.String(), "Send", "Service")

	j.Stop(0)
}
