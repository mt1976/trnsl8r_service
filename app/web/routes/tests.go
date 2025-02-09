package routes

import (
	"html/template"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/mt1976/frantic-plum/logger"
	"github.com/mt1976/frantic-plum/paths"
	"github.com/mt1976/trnsl8r_service/app/web/pages"
)

func Test(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

	trace(r)
	title := "Tests"
	action := "View"

	t := template.Must(template.ParseFiles(getTemplate(title, action), paths.HTMLTemplate())) // Create a template.
	w.Header().Set("Content-Type", "text/html")
	w.Header().Add("Application", appName)

	err := t.Execute(w, pages.Generic(title, action)) // merge.
	if err != nil {
		logger.ErrorLogger.Print(err.Error())
	}
}
