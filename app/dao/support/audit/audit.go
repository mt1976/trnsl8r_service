package audit

import (
	"context"
	"os"
	"strings"
	"time"

	"github.com/mt1976/trnsl8r_service/app/dao/support/database"
	"github.com/mt1976/trnsl8r_service/app/support/application"
	"github.com/mt1976/trnsl8r_service/app/support/config"
	"github.com/mt1976/trnsl8r_service/app/support/date"
	"github.com/mt1976/trnsl8r_service/app/support/logger"
)

var name = "Audit"
var cfg *config.Configuration

type Action struct {
	code    string
	message string
	silent  bool
}

func init() {
	cfg = config.Get()
}

func getDBVersion() int {
	// Implement the logic to get the DB version without importing the dao package
	return database.Version
}

type Audit struct {
	CreatedAt        time.Time
	CreatedBy        string
	CreatedOn        string
	CreatedAtDisplay string
	Updates          []AuditUpdateInfo
	DeletedAt        time.Time
	DeletedBy        string
	DeletedOn        string
	DeletedAtDisplay string
	AuditSequence    int
	DBVersion        int
	//Empty     time.Time // Convience Field - Used to avoid erros with dates.
}

type AuditUpdateInfo struct {
	UpdatedAt        time.Time
	UpdateAction     string
	UpdatedBy        string
	UpdatedOn        string
	UpdatedAtDisplay string
	UpdateNotes      string
}

const (
	AUDITMSG = "[%v] Action=[%s] At=[%v] By=[%v] On=[%v] Notes=[%v]"
)

var (
	CREATE       Action
	DELETE       Action
	UPDATE       Action
	ERASE        Action
	CLONE        Action
	NOTIFICATION Action
	SERVICE      Action
	SILENT       Action
	GRANT        Action
	REVOKE       Action
	PROCESS      Action
	IMPORT       Action

// REPAIR       Action
)

func init() {
	CREATE = Action{code: "NEW", message: "New Record", silent: false}
	DELETE = Action{code: "DEL", message: "Delete Record", silent: false}
	UPDATE = Action{code: "UPD", message: "Update Record", silent: false}
	ERASE = Action{code: "ERS", message: "Erase Record", silent: false}
	CLONE = Action{code: "CLN", message: "Clone Record", silent: false}
	NOTIFICATION = Action{code: "NTE", message: "Notification Sent", silent: false}
	SERVICE = Action{code: "SVC", message: "Service Action", silent: false}
	SILENT = Action{code: "SIL", message: "Silent Action", silent: true}
	GRANT = Action{code: "GNT", message: "Grant", silent: false}
	REVOKE = Action{code: "RVK", message: "Revoke", silent: false}
	PROCESS = Action{code: "PRC", message: "Process", silent: false}
	IMPORT = Action{code: "IMP", message: "Import", silent: false}

	// REPAIR = Action{text: "REP", message: "Repaired"}
}

func (a *Action) WithMessage(in string) Action {
	a.message = in
	return *a
}

func (a *Action) popMessage() string {
	message := a.message
	a.message = ""
	return message
}

func (a *Audit) Action(ctx context.Context, action Action) error {

	message := action.popMessage()

	//	start := timing.Start("Audit", action.text, message)

	auditTime := time.Now()
	auditDisplay := date.FormatAudit(auditTime)
	// auditUser := support.GetActiveUserCode()
	auditUser := getUser()
	auditHost := application.HostName()

	if auditUser == "" {
		logger.ErrorLogger.Printf("[%v] Error=[%v]", strings.ToUpper(name), "No Active User")
		logger.InfoLogger.Printf("[%v] Action=[%v] Message=[%v]", strings.ToUpper(name), action.code, message)
		os.Exit(0)
	}
	//updateAction := action

	if action.Is(CREATE) {
		a.CreatedAt = auditTime
		a.CreatedBy = auditUser
		a.CreatedOn = auditHost
		a.CreatedAtDisplay = auditDisplay
	}

	if action.Is(DELETE) {
		a.DeletedAt = auditTime
		a.DeletedBy = auditUser
		a.DeletedOn = auditHost
		a.DeletedAtDisplay = auditDisplay
	}

	if a.AuditSequence == 0 {
		a.AuditSequence = 1
	} else {
		a.AuditSequence++
	}

	update := AuditUpdateInfo{}

	update.UpdatedAt = auditTime
	update.UpdatedBy = auditUser
	update.UpdatedOn = auditHost
	update.UpdatedAtDisplay = auditDisplay
	update.UpdateAction = action.code
	update.UpdateNotes = message
	// a.DBVersion = dao.Version
	a.DBVersion = getDBVersion()
	if !(action.Is(SERVICE) || action.Is(SILENT) || action.IsSilent()) {
		a.Updates = append(a.Updates, update)
	}

	a.DBVersion = getDBVersion()

	logger.AuditLogger.Printf(AUDITMSG, strings.ToUpper(name), action.code, auditDisplay, auditUser, auditHost, message)
	//	start.Stop(1)
	return nil
}

func (a *Audit) Spew() error {
	// Spew the Audit Data
	noAudit := len(a.Updates)
	//logger.InfoLogger.Printf(" No Updates=[%v]", noAudit)
	if noAudit > 0 {
		for i := 0; i < noAudit; i++ {
			upd := a.Updates[i]
			logger.TraceLogger.Printf(AUDITMSG, strings.ToUpper(name), upd.UpdateAction, upd.UpdatedAtDisplay, upd.UpdatedBy, upd.UpdatedOn, upd.UpdateNotes)
		}
	}
	return nil
}

func (a *Action) Is(inAction Action) bool {
	return a.code == inAction.code
}

func (a *Action) IsSilent() bool {
	return a.silent
}

func (a *Action) Message() string {
	return a.message
}

func (a *Action) Text() string {
	return a.code
}

func (a *Action) SetMessage(in string) {
	a.message = in
}

func (a *Action) SetText(in string) {
	a.code = in
}

func (a *Action) Code() string {
	return a.code
}

func getUser() string {
	// Implement the logic to get the user without importing the dao package
	return "sys" + "_" + "service"
}
