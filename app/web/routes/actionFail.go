package routes

import (
	"html/template"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/mt1976/frantic-core/commonConfig"
	"github.com/mt1976/frantic-core/logHandler"
	"github.com/mt1976/frantic-core/paths"
	"github.com/mt1976/trnsl8r_service/app/business/domains"
	"github.com/mt1976/trnsl8r_service/app/business/translate"
	"github.com/mt1976/trnsl8r_service/app/web/pages"
)

func Fail(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	logHandler.EventLogger.Printf("[%v] [FAIL] [View] Fail", domains.ROUTE.String())

	settings := commonConfig.Get()

	title := "Fail"
	action := "Message"

	trace(r)

	messageType := r.URL.Query().Get(msgTypeKey)
	messageTitle := r.URL.Query().Get(msgTitleKey)
	messageContent := r.URL.Query().Get(msgContentKey)
	messageAction := r.URL.Query().Get(msgActionKey)

	logHandler.ErrorLogger.Printf("[%v] [FAIL] message Type: [%v]=[%v]\n", domains.ROUTE.String(), msgTypeKey, messageType)
	logHandler.ErrorLogger.Printf("[%v] [FAIL] message title: [%v]=[%v]\n", domains.ROUTE.String(), msgTitleKey, messageTitle)
	logHandler.ErrorLogger.Printf("[%v] [FAIL] message content: [%v]=[%v]\n", domains.ROUTE.String(), msgContentKey, messageContent)
	logHandler.ErrorLogger.Printf("[%v] [FAIL] message action: [%v]=[%v]\n", domains.ROUTE.String(), msgActionKey, messageAction)

	t := template.Must(template.ParseFiles(getTemplate(title, action), paths.HTMLTemplate())) // Create a template.

	w.Header().Set("Content-Type", "text/html")
	w.Header().Add("Application", settings.GetApplication_Name())

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

	pg.MessageType = translate.Get(pg.MessageType, "")
	pg.PageTitle = translate.Get(pg.PageTitle, "")
	pg.Message = translate.Get(pg.Message, "")
	pg.PageAction = translate.Get(pg.PageAction, "")

	err := t.Execute(w, pg) // merge.
	if err != nil {
		logHandler.ErrorLogger.Print(err.Error())
	}
}

func buildFailPS(msg string, title string, content string, action string) httprouter.Params {
	ps := httprouter.Params{}
	ps = append(ps, httprouter.Param{Key: msgTitleKey, Value: translate.Get(msg, "")})
	ps = append(ps, httprouter.Param{Key: msgTitleKey, Value: translate.Get(title, "")})
	ps = append(ps, httprouter.Param{Key: msgContentKey, Value: translate.Get(content, "")})
	ps = append(ps, httprouter.Param{Key: msgActionKey, Value: translate.Get(action, "")})
	return ps
}
