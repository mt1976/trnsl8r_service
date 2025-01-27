package routes

import (
	"html/template"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/mt1976/trnsl8r_service/app/dao/textStore"
	"github.com/mt1976/trnsl8r_service/app/support/logger"
	"github.com/mt1976/trnsl8r_service/app/support/paths"
	"github.com/mt1976/trnsl8r_service/app/web/pages"
)

func Test(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

	trace(r)
	title := "Tests"
	action := "View"

	t := template.Must(template.ParseFiles(getTemplate(title, action), paths.HTMLTemplate())) // Create a template.
	w.Header().Set("Content-Type", "text/html")
	w.Header().Add("Application", cfg.ApplicationName())

	err := t.Execute(w, pages.Generic(title, action)) // merge.
	if err != nil {
		logger.ErrorLogger.Print(err.Error())
	}
}

func ExportTranslations(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	//logger.EventLogger.Printf("[TEST] [View]")

	trace(r)

	err := textStore.ExportCSV()
	if err != nil {
		logger.ErrorLogger.Print(err.Error())
		oops(w, r, nil, "error", err.Error())
	}
	//successMessage(w, r, nil, "success - translations exported")
	logger.EventLogger.Printf("[TEST] [Export] [Translations] [Success]")
	http.Redirect(w, r, "/test", http.StatusSeeOther)
}
