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
	logHandler.Event.Println(name)
	ctx := r.Context()
	trace(r)

	jobs.DatabasePrune.Run()

	msg := name + " " + translate.Get(ctx, "Complete", "")
	msg = translate.Get(ctx, msg, "")

	successMessage(w, r, ps, msg)
}
