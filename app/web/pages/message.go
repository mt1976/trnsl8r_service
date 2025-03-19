package pages

import (
	"github.com/mt1976/trnsl8r_service/app/business/translate"
)

func Message(title, action, msgType, msg string) *Page {
	p := New(title, action)

	p.Message = msg // Dont trranslate here as the message is already translated
	p.MessageType = translate.Get(msgType, "")

	p.SingleItem = true
	//p.PageAction = text.Get("Oops")
	//p.PageTitle = text.Get("Oops")
	return p
}
