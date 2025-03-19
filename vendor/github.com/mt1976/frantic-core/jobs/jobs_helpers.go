package jobs

import (
	"fmt"
	"strings"
	"time"

	"github.com/jsuar/go-cron-descriptor/pkg/crondescriptor"
	"github.com/mt1976/frantic-core/dao/actions"
	"github.com/mt1976/frantic-core/dateHelpers"
	"github.com/mt1976/frantic-core/logHandler"
	"github.com/mt1976/frantic-core/stringHelpers"
	"github.com/mt1976/frantic-core/timing"
)

func StartOfDay(t time.Time) time.Time {
	// Purpose: To remove the time from a date
	return dateHelpers.StartOfDay(t)
}

func BeforeOrEqualTo(t1, t2 time.Time) bool {
	return dateHelpers.IsBeforeOrEqualTo(t1, t2)
}

func AfterOrEqualTo(t1, t2 time.Time) bool {
	return dateHelpers.IsAfterOrEqualTo(t1, t2)
}

func NextRun(name, schedule string) string {
	// Purpose: To determine the next run time of a job
	rtn := fmt.Sprintf("[%v] [%v] NextRun=[%v]", domain, name, GetHumanReadableCronFreq(schedule))
	logHandler.ServiceLogger.Println(rtn)
	return rtn
}

// Announce - Announce the start of a job to the log
// Deprecated: Use PreRun instead
func Announce(name, inAction string) {
	//logHandler.ServiceBanner(domain, name, inAction)
}

func GetHumanReadableCronFreq(freq string) string {
	bkHuman1, _ := crondescriptor.NewCronDescriptor(freq)
	bkHuman, _ := bkHuman1.GetDescription(crondescriptor.Full)
	return *bkHuman
}

func PreRun(job Job) {
	// Purpose: To log the start of a job
	logHandler.ServiceLogger.Printf("[%v] Job %v - Started", domain, stringHelpers.DQuote(job.Name()))
}

func PostRun(job Job) {
	// Purpose: To log the completion of a job
	nextRun := GetHumanReadableCronFreq(job.Schedule())
	logHandler.ServiceLogger.Printf("[%v] Job %v - Completed", domain, stringHelpers.DQuote(job.Name()))
	logHandler.ServiceLogger.Printf("[%v] Job %v Scheduled [%v] [%v]", domain, stringHelpers.DQuote(job.Name()), job.Schedule(), nextRun)
}

func CodedName(job Job) string {
	// Purpose: To return the coded name of a job
	name := job.Name()
	name = strings.Replace(name, " ", "", -1)
	name = strings.Replace(name, "-", "", -1)
	name = strings.Replace(name, "_", "", -1)
	name = stringHelpers.RemoveSpecialChars(name)
	return name
}

func AddJobToScheduler(j Job) {
	//logHandler.ServiceLogger.Printf("[%v] Scheduling Job [%v] [%v]", domain, j.Name(), j.Schedule())
	clock := timing.Start(domain, actions.SCHEDULE.GetCode(), j.Name())
	// Start the job
	jobID, err := scheduledTasks.AddFunc(j.Schedule(), j.Service())
	if err != nil {
		logHandler.ErrorLogger.Printf("[%v] Job %v Scheduling Error=[%v]", domain, stringHelpers.DQuote(j.Name()), err.Error())
		return
	}
	nextRun := GetHumanReadableCronFreq(j.Schedule())
	logHandler.ServiceLogger.Printf("[%v] Job %v Scheduled [%v] [%v] (id=%v)", domain, stringHelpers.DQuote(j.Name()), j.Schedule(), nextRun, jobID)
	clock.Stop(1)
}

func AddJobsToScheduler(jobs []Job) {
	clock := timing.Start(domain, actions.SCHEDULE.GetCode(), "Jobs")
	// Schedule a list of jobs
	for _, j := range jobs {
		AddJobToScheduler(j)
	}
	clock.Stop(len(jobs))
}

func StartScheduler() {
	clock := timing.Start(domain, actions.START.GetCode(), "Scheduler")
	logHandler.ServiceLogger.Printf("[%v] Scheduler - Starting", domain)
	// Start the scheduler
	scheduledTasks.Start()

	noEntries := len(scheduledTasks.Entries())
	// Log the scheduled tasks
	// for x, entry := range scheduledTasks.Entries() {
	// 	logHandler.ServiceLogger.Printf("(%v/%v) [%v] [%v] [%v]", x+1, noEntries, entry.ID, entry.Next, entry.Job)
	// }
	logHandler.ServiceLogger.Printf("[%v] Scheduler - Started", domain)
	clock.Stop(noEntries + 1)
}
