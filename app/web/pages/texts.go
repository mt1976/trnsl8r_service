package pages

import (
	"github.com/mt1976/frantic-plum/logger"
	"github.com/mt1976/trnsl8r_service/app/dao/textStore"
)

func TextView(title, action, id string) (*Page, error) {
	logger.InfoLogger.Printf("Page: TextView %+v", id)
	p := New(title, action)

	txt, err := textStore.GetBySignature(id)
	if err != nil {
		return p, err
	}
	p.TextItem = txt

	//fmt.Printf("Page: Settings %+v", p)
	return p, nil
}

func TextList(title, action string) (*Page, error) {

	p := New(title, action)

	TextList, err := textStore.GetAll()
	if err != nil {
		return p, err
	}
	p.TextList = TextList
	//fmt.Printf("Page: Settings %+v", p)
	return p, nil
}

func TextEdit(title, action, id string) (*Page, error) {
	p := New(title, action)

	txt, err := textStore.GetBySignature(id)
	if err != nil {
		return p, err
	}
	p.TextItem = txt

	//fmt.Printf("Page: Settings %+v", p)
	return p, nil
}

func TextUpdate(title, action, id string) (*Page, error) {
	p := New(title, action)

	txt, err := textStore.GetBySignature(id)
	if err != nil {
		return p, err
	}
	p.TextItem = txt

	//fmt.Printf("Page: Settings %+v", p)
	return p, nil
}
