package routes

import (
	"html/template"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/mt1976/frantic-plum/common"
	"github.com/mt1976/frantic-plum/logger"
	"github.com/mt1976/frantic-plum/paths"
	"github.com/mt1976/trnsl8r_service/app/business/domains"
	"github.com/mt1976/trnsl8r_service/app/business/translation"
	"github.com/mt1976/trnsl8r_service/app/web/pages"
)

func Fail(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	logger.EventLogger.Printf("[%v] [FAIL] [View] Fail", domains.ROUTE.String())

	settings := common.Get()

	title := "Fail"
	action := "Message"

	trace(r)

	messageType := r.URL.Query().Get(settings.MessageTypeKey())
	messageTitle := r.URL.Query().Get(settings.MessageTitleKey())
	messageContent := r.URL.Query().Get(settings.MessageContentKey())
	messageAction := r.URL.Query().Get(settings.MessageActionKey())

	logger.ErrorLogger.Printf("[%v] [FAIL] message Type: [%v]=[%v]\n", domains.ROUTE.String(), settings.MessageTypeKey(), messageType)
	logger.ErrorLogger.Printf("[%v] [FAIL] message title: [%v]=[%v]\n", domains.ROUTE.String(), settings.MessageTitleKey(), messageTitle)
	logger.ErrorLogger.Printf("[%v] [FAIL] message content: [%v]=[%v]\n", domains.ROUTE.String(), settings.MessageContentKey(), messageContent)
	logger.ErrorLogger.Printf("[%v] [FAIL] message action: [%v]=[%v]\n", domains.ROUTE.String(), settings.MessageActionKey(), messageAction)

	t := template.Must(template.ParseFiles(getTemplate(title, action), paths.HTMLTemplate())) // Create a template.

	w.Header().Set("Content-Type", "text/html")
	w.Header().Add("Application", settings.ApplicationName())

	pg := pages.Generic(title, action)

	pg.MessageType = "error"
	if messageType != "" {
		pg.MessageType = messageType
	}

	pg.PageTitle = "Error"
	if messageTitle != "" {
		pg.PageTitle = messageTitle
	}

	pg.Message = "An error has occurred"
	if messageContent != "" {
		pg.Message = messageContent
	}

	pg.PageAction = "Error"
	if messageAction != "" {
		pg.PageAction = messageAction
	}

	pg.MessageType = translation.Get(pg.MessageType)
	pg.PageTitle = translation.Get(pg.PageTitle)
	pg.Message = translation.Get(pg.Message)
	pg.PageAction = translation.Get(pg.PageAction)

	err := t.Execute(w, pg) // merge.
	if err != nil {
		logger.ErrorLogger.Print(err.Error())
	}
}

func buildFailPS(msg string, title string, content string, action string) httprouter.Params {
	ps := httprouter.Params{}
	ps = append(ps, httprouter.Param{Key: settings.MessageTitleKey(), Value: translation.Get(msg)})
	ps = append(ps, httprouter.Param{Key: settings.MessageTitleKey(), Value: translation.Get(title)})
	ps = append(ps, httprouter.Param{Key: settings.MessageContentKey(), Value: translation.Get(content)})
	ps = append(ps, httprouter.Param{Key: settings.MessageActionKey(), Value: translation.Get(action)})
	return ps
}
