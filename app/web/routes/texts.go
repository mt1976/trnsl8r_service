package routes

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/mt1976/frantic-plum/logger"
	"github.com/mt1976/frantic-plum/paths"
	"github.com/mt1976/trnsl8r_service/app/business/translation"
	"github.com/mt1976/trnsl8r_service/app/dao/textStore"
	"github.com/mt1976/trnsl8r_service/app/web/pages"
)

func TextList(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	title := "Texts"
	action := "List"

	trace(r)

	t := template.Must(template.ParseFiles(getTemplate(title, action), paths.HTMLTemplate())) // Create a template.
	w.Header().Set("Content-Type", "text/html")

	page, err := pages.TextList(title, action)
	if err != nil {
		logger.ErrorLogger.Printf("Error=[%v]", err.Error())
		oops(w, r, ps, page.MessageType, page.Message)
		return
	}

	err = t.Execute(w, page) // merge.
	if err != nil {
		logger.ErrorLogger.Printf("Error=[%v]", err.Error())
	}
}

func TextView(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	title := "Texts"
	action := "View"

	trace(r)

	t := template.Must(template.ParseFiles(getTemplate(title, action), paths.HTMLTemplate())) // Create a template.
	w.Header().Set("Content-Type", "text/html")

	id := ps.ByName("id")

	page, err := pages.TextView(title, action, id)
	if err != nil {
		logger.ErrorLogger.Printf("Error=[%v]", err.Error())
		oops(w, r, ps, page.MessageType, page.Message)
		return
	}

	err = t.Execute(w, page) // merge.
	if err != nil {
		logger.ErrorLogger.Printf("Error=[%v]", err.Error())
	}
}

func TextEdit(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	trace(r)

	title := "Texts"
	action := "Edit"

	t := template.Must(template.ParseFiles(getTemplate(title, action), paths.HTMLTemplate())) // Create a template.
	w.Header().Set("Content-Type", "text/html")

	id := ps.ByName("id")

	page, err := pages.TextEdit(title, action, id)
	if err != nil {
		logger.ErrorLogger.Printf("Error=[%v]", err.Error())
		oops(w, r, ps, page.MessageType, page.Message)
		return
	}

	err = t.Execute(w, page) // merge.
	if err != nil {
		logger.ErrorLogger.Printf("Error=[%v]", err.Error())
	}
}

func TextUpdate(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	//title := "Texts"
	//action := "Update"

	trace(r)
	logger.TraceLogger.Printf("Params=%+v", ps)
	logger.TraceLogger.Printf("Request=%+v", r)
	logger.TraceLogger.Printf("r.Form: %+v %v\n", r.Form, len(r.Form))
	logger.TraceLogger.Printf("r.Body: %+v\n", r.Body)

	//id := r.FormValue("entity")
	//fmt.Printf("entity: %v\n", id)
	id, err := getIDString(ps)
	if err != nil {
		logger.ErrorLogger.Printf("Error=[%v]", err.Error())
		oops(w, r, ps, translation.Get("error", ""), err.Error())
		return
	}
	if id == "" {
		msg := translation.Get("Invalid ID: ID is required", "")
		logger.InfoLogger.Print(msg)
		oops(w, r, ps, translation.Get("error", ""), msg)
		return
	}

	t, err := textStore.GetBySignature(id)
	if err != nil {
		logger.ErrorLogger.Printf("Error=[%v]", err.Error())
		oops(w, r, ps, translation.Get("error", ""), err.Error())
		return
	}

	//newMessage := r.FormValue("message")
	newMessage := r.FormValue("Message")
	oldMessage := t.Message

	if newMessage == "" {
		msg := translation.Get("Invalid Name: Message is required", "")
		logger.InfoLogger.Print(msg)
		oops(w, r, ps, translation.Get("error", ""), msg)
		return
	}

	if newMessage == oldMessage {
		msg := translation.Get("No Change: Message is the same", "")
		logger.InfoLogger.Print(msg)
		oops(w, r, ps, translation.Get("error", ""), msg)
		return
	}

	t.Message = newMessage

	msg := "Message updated from [%v]->[%v]"
	msg = translation.Get(msg, "")
	msg = fmt.Sprintf(msg, oldMessage, newMessage)
	msg2 := msg
	logmsg := "[TEXT] " + msg
	logger.InfoLogger.Println(logmsg)

	err = t.Update(nil, msg)
	if err != nil {
		oops(w, r, ps, translation.Get("error", ""), err.Error())
		return
	}
	//winmsg := "Zone %v"
	winmsg := fmt.Sprintf(translation.Get("Text %v : ", ""), t.Signature) + msg2
	successMessage(w, r, ps, winmsg)
}
