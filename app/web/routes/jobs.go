package routes

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/mt1976/frantic-core/logHandler"
	"github.com/mt1976/trnsl8r_service/app/business/translate"
	"github.com/mt1976/trnsl8r_service/app/jobs"
)

func PruneBackups(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	name := "Database Backup Housekeeping"
	logHandler.EventLogger.Println(name)

	trace(r)

	jobs.DatabasePrune.Run()

	msg := name + " " + translate.Get("Complete", "")
	msg = translate.Get(msg, "")

	successMessage(w, r, ps, msg)
}
