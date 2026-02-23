package pages

import (
	"context"

	"github.com/mt1976/trnsl8r_service/app/business/translate"
)

func Generic(ctx context.Context, title, action string) *Page {
	p := New(ctx, title, action)
	p.PageAction = translate.Get(ctx, action, "")
	p.PageTitle = translate.Get(ctx, title, "")

	////spew.Dump(p
	return p
}
