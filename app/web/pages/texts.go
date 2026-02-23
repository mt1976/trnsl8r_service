package pages

import (
	"context"

	"github.com/mt1976/frantic-core/logHandler"
	"github.com/mt1976/trnsl8r_service/app/dao/textStore"
)

func TextView(ctx context.Context, title, action, id string) (*Page, error) {
	logHandler.Info.Printf("Page: TextView %+v", id)
	p := New(ctx, title, action)

	txt, err := textStore.GetBy(textStore.Fields.Signature, id)
	if err != nil {
		return p, err
	}
	p.TextItem = txt

	// fmt.Printf("Page: Settings %+v", p)
	return p, nil
}

func TextList(ctx context.Context, title, action string) (*Page, error) {
	p := New(ctx, title, action)

	TextList, err := textStore.GetAll()
	if err != nil {
		return p, err
	}
	p.TextList = TextList
	// fmt.Printf("Page: Settings %+v", p)
	return p, nil
}

func TextEdit(ctx context.Context, title, action, id string) (*Page, error) {
	p := New(ctx, title, action)

	txt, err := textStore.GetBy(textStore.Fields.Signature, id)
	if err != nil {
		return p, err
	}
	p.TextItem = txt

	// fmt.Printf("Page: Settings %+v", p)
	return p, nil
}

func TextUpdate(ctx context.Context, title, action, id string) (*Page, error) {
	p := New(ctx, title, action)

	txt, err := textStore.GetBy(textStore.Fields.Signature, id)
	if err != nil {
		return p, err
	}
	p.TextItem = txt

	// fmt.Printf("Page: Settings %+v", p)
	return p, nil
}
