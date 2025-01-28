package routes

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/mt1976/frantic-plum/logger"
	"github.com/mt1976/trnsl8r_service/app/business/translation"
	"github.com/mt1976/trnsl8r_service/app/jobs"
)

func PruneBackups(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	name := "Database Backup Housekeeping"
	logger.EventLogger.Println(name)

	trace(r)

	jobs.DatabasePrune.Run()

	msg := name + " " + translation.Get("Complete")
	msg = translation.Get(msg)

	successMessage(w, r, ps, msg)
}
