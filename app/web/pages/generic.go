package pages

import (
	"github.com/mt1976/trnsl8r_service/app/business/translation"
)

func Generic(title, action string) *Page {

	p := New(title, action)
	p.PageAction = translation.Get(action)
	p.PageTitle = translation.Get(title)

	////spew.Dump(p
	return p
}
