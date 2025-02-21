package routes

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/mt1976/frantic-core/logHandler"
	"github.com/mt1976/trnsl8r_service/app/jobs"
)

func Backup(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	logHandler.EventLogger.Println("Backup")

	trace(r)

	err := jobs.DatabaseBackup.Run()
	if err != nil {
		logHandler.ErrorLogger.Print(err.Error())
	}

	w.Header().Set("Content-Type", "text/html")
	_, err = w.Write([]byte("done"))
	if err != nil {
		logHandler.ErrorLogger.Print(err.Error())
	}
}
