package routes

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
	logger "github.com/mt1976/frantic-core/logHandler"
	"github.com/mt1976/trnsl8r_service/app/jobs"
)

func Backup(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	logger.EventLogger.Println("Backup")

	trace(r)

	err := jobs.DatabaseBackup.Run()
	if err != nil {
		logger.ErrorLogger.Print(err.Error())
	}

	w.Header().Set("Content-Type", "text/html")
	_, err = w.Write([]byte("done"))
	if err != nil {
		logger.ErrorLogger.Print(err.Error())
	}
}
