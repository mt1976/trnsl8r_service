package routes

import (
	"html/template"
	"net/http"

	"github.com/julienschmidt/httprouter"
	logger "github.com/mt1976/frantic-core/logHandler"
	"github.com/mt1976/frantic-core/paths"
	"github.com/mt1976/trnsl8r_service/app/web/pages"
)

func Success(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	successMessage(w, r, ps, "")
}

func successMessage(w http.ResponseWriter, r *http.Request, ps httprouter.Params, msg string) {

	title := "Success"
	action := "Message"

	trace(r)

	t := template.Must(template.ParseFiles(getTemplate(title, action), paths.HTMLTemplate())) // Create a template.
	w.Header().Set("Content-Type", "text/html")
	p := pages.Message(title, action, "success", msg)
	//p.Message = msg

	err := t.Execute(w, p) // merge.
	if err != nil {
		logger.ErrorLogger.Print(err.Error())
	}
}
