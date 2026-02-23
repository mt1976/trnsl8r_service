package routes

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/mt1976/frantic-core/commonConfig"
	"github.com/mt1976/frantic-core/logHandler"
	"github.com/mt1976/frantic-core/paths"
	"github.com/mt1976/trnsl8r_service/app/business/translate"
	"github.com/mt1976/trnsl8r_service/app/dao/textStore"
	"github.com/mt1976/trnsl8r_service/app/web/pages"
)

func TextList(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	title := "Texts"
	action := "List"

	trace(r)

	t := template.Must(template.ParseFiles(getTemplate(title, action), paths.HTMLTemplate())) // Create a template.
	w.Header().Set("Content-Type", "text/html")

	page, err := pages.TextList(r.Context(), title, action)
	if err != nil {
		logHandler.Error.Printf("Error=[%v]", err.Error())
		oops(w, r, ps, page.MessageType, page.Message)
		return
	}

	err = t.Execute(w, page) // merge.
	if err != nil {
		logHandler.Error.Printf("Error=[%v]", err.Error())
	}
}

func TextView(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	title := "Texts"
	action := "View"

	trace(r)

	t := template.Must(template.ParseFiles(getTemplate(title, action), paths.HTMLTemplate())) // Create a template.
	w.Header().Set("Content-Type", "text/html")

	id := ps.ByName("id")

	page, err := pages.TextView(r.Context(), title, action, id)
	if err != nil {
		logHandler.Error.Printf("Error=[%v]", err.Error())
		oops(w, r, ps, page.MessageType, page.Message)
		return
	}

	err = t.Execute(w, page) // merge.
	if err != nil {
		logHandler.Error.Printf("Error=[%v]", err.Error())
	}
}

func TextEdit(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	trace(r)

	title := "Texts"
	action := "Edit"

	t := template.Must(template.ParseFiles(getTemplate(title, action), paths.HTMLTemplate())) // Create a template.
	w.Header().Set("Content-Type", "text/html")

	id := ps.ByName("id")

	page, err := pages.TextEdit(r.Context(), title, action, id)
	if err != nil {
		logHandler.Error.Printf("Error=[%v]", err.Error())
		oops(w, r, ps, page.MessageType, page.Message)
		return
	}

	err = t.Execute(w, page) // merge.
	if err != nil {
		logHandler.Error.Printf("Error=[%v]", err.Error())
	}
}

func TextUpdate(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	set := commonConfig.Get()
	// title := "Texts"
	// action := "Update"

	trace(r)
	logHandler.Trace.Printf("Params=%+v", ps)
	logHandler.Trace.Printf("Request=%+v", r)
	logHandler.Trace.Printf("r.Form: %+v %v\n", r.Form, len(r.Form))
	logHandler.Trace.Printf("r.Body: %+v\n", r.Body)

	// id := r.FormValue("entity")
	// fmt.Printf("entity: %v\n", id)
	id, err := getIDString(ps, r.Context())
	if err != nil {
		logHandler.Error.Printf("Error=[%v]", err.Error())
		oops(w, r, ps, translate.Get(r.Context(), "error", ""), err.Error())
		return
	}
	if id == "" {
		msg := translate.Get(r.Context(), "invalid ID: ID is required", "")
		logHandler.Info.Print(msg)
		oops(w, r, ps, translate.Get(r.Context(), "error", ""), msg)
		return
	}

	t, err := textStore.GetBy(textStore.Fields.Signature, id)
	if err != nil {
		logHandler.Error.Printf("Error=[%v]", err.Error())
		oops(w, r, ps, translate.Get(r.Context(), "error", ""), err.Error())
		return
	}

	// newMessage := r.FormValue("message")
	newMessage := r.FormValue("Message")
	oldMessage := t.Message

	if newMessage == "" {
		msg := translate.Get(r.Context(), "invalid Name: Message is required", "")
		logHandler.Info.Print(msg)
		oops(w, r, ps, translate.Get(r.Context(), "error", ""), msg)
		return
	}

	// Deal with any possible localised translations
	//
	// Get the current valid locales
	msgUpdated := false
	locales := getLocalesList(set)
	for _, locale := range locales {
		if locale == "" {
			continue
		}

		localText := r.FormValue(locale)

		logHandler.Event.Printf("Update Locale=[%v] Text=[%v] Original=[%v]", locale, localText, t.Localised[locale])

		if t.Localised[locale] != localText {
			t.Localised[locale] = localText
			msgUpdated = true
		}

	}

	logHandler.Event.Printf("newMessage=[%v] oldMessage=[%v] msgUpdated=[%v]", newMessage, oldMessage, msgUpdated)

	if !msgUpdated {
		if newMessage == oldMessage {
			msg := translate.Get(r.Context(), "no change: Message is the same", "")
			logHandler.Info.Print(msg)
			oops(w, r, ps, translate.Get(r.Context(), "error", ""), msg)
			return
		}
	}

	t.Message = newMessage

	// Get the current valid locales

	msg := "Message updated from [%v]->[%v]"
	msg = translate.Get(r.Context(), msg, "")
	msg = fmt.Sprintf(msg, oldMessage, newMessage)
	msg2 := msg
	logmsg := "[TEXT] " + msg
	logHandler.Info.Println(logmsg)

	err = t.Update(nil, msg)
	if err != nil {
		oops(w, r, ps, translate.Get(r.Context(), "error", ""), err.Error())
		return
	}
	// winmsg := "Zone %v"
	winmsg := fmt.Sprintf(translate.Get(r.Context(), "Text %v : ", ""), t.Signature) + msg2
	successMessage(w, r, ps, winmsg)
}
