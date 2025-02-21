package routes

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/mt1976/frantic-core/logHandler"
	"github.com/mt1976/frantic-core/timing"
)

func Health(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

	start := timing.Start("Health", "Monitor", "Health Check")
	logHandler.EventLogger.Println("Health")
	logHandler.ServiceLogger.Println("Health Check")

	trace(r)

	w.Header().Set("Content-Type", "text/html")
	_, err := w.Write([]byte("ok"))
	if err != nil {
		logHandler.ErrorLogger.Println(err.Error())
	}
	start.Stop(1)
}
