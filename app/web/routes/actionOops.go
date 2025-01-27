package routes

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/mt1976/trnsl8r_service/app/support/logger"
	"github.com/mt1976/trnsl8r_service/app/support/paths"
	page "github.com/mt1976/trnsl8r_service/app/web/pages"
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

	title := "Oops"
	action := "Message"

	//OOPS.Spew()
	fmt.Printf("paths.HTMLTemplate(): %v\n", paths.HTMLTemplate())
	fmt.Printf("msgType: %v\n", msgType)
	fmt.Printf("msg: %v\n", msg)

	t := template.Must(template.ParseFiles(getTemplate(title, action), paths.HTMLTemplate())) // Create a template.
	w.Header().Set("Content-Type", "text/html")

	err := t.Execute(w, page.Message(title, action, msgType, msg)) // merge.
	if err != nil {
		logger.ErrorLogger.Print(err.Error())
	}
}
