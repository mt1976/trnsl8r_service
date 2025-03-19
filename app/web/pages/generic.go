package pages

import (
	"github.com/mt1976/trnsl8r_service/app/business/translate"
)

func Generic(title, action string) *Page {

	p := New(title, action)
	p.PageAction = translate.Get(action, "")
	p.PageTitle = translate.Get(title, "")

	////spew.Dump(p
	return p
}
