package routes

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/mt1976/trnsl8r_service/app/support/logger"
	"github.com/mt1976/trnsl8r_service/app/support/timing"
)

func Health(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

	start := timing.Start("Health", "Monitor", "Health Check")
	logger.EventLogger.Println("Health")
	logger.ServiceLogger.Println("Health Check")

	trace(r)

	w.Header().Set("Content-Type", "text/html")
	_, err := w.Write([]byte("ok"))
	if err != nil {
		logger.ErrorLogger.Println(err.Error())
	}
	start.Stop(1)
}
