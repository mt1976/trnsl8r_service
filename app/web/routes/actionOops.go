package routes

import (
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/mt1976/frantic-core/logHandler"
)

func Oops(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	trace(r)
	oops(w, r, ps, "", "")
}

func oops(w http.ResponseWriter, r *http.Request, ps httprouter.Params, msgType, msg string) {
	//	logger.EventLogger.Printf("[ACTION] [View] Oops - %s %v", msgType, msg)
	//	logger.InfoLogger.Println("Oops " + msgType + " - " + msg)
	//msg = text.Get(msg)

	trace(r)

	w.Header().Set("Content-Type", "application/json")
	w.Header().Add("Application", appName)
	w.WriteHeader(http.StatusNotFound)
	fmt.Fprintf(w, "{\"message\":\"%v\"}", msg)
	logHandler.ErrorLogger.Printf("[ACTION] Oops - %s %v", msgType, msg)
}
