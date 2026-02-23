package routes

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/julienschmidt/httprouter"
	"github.com/mt1976/frantic-core/htmlHelpers"
	"github.com/mt1976/frantic-core/logHandler"
	"github.com/mt1976/frantic-core/paths"
	"github.com/mt1976/trnsl8r_service/app/business/translate"
)

func trace(r *http.Request) {
	mesg := translate.Get(r.Context(), "Method=[%s] URI=[%s] Header[%v] Context[%+v] RequestURI=[%s]", "")
	logHandler.Trace.Printf(mesg, r.Method, r.URL, r.Header, r.Context(), r.RequestURI)
}

func getID(ps httprouter.Params, ctx context.Context) (int, error) {
	id, err := iconv(ps.ByName("id"), ctx)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func getIDString(ps httprouter.Params, ctx context.Context) (string, error) {
	id := ps.ByName("id")
	return id, nil
}

func getItemName(ps httprouter.Params, ctx context.Context) (string, error) {
	itemName := ps.ByName("item")
	if itemName == "" {
		msg := translate.Get(ctx, "Invalid Item Name [%s]", "")
		return "", fmt.Errorf(msg, itemName)
	}
	return itemName, nil
}

func iconv(arg interface{}, ctx context.Context) (int, error) {
	switch x := arg.(type) {
	case int:
		return x, nil
	case string:
		i, er := strconv.Atoi(x)
		if er != nil {
			msg := translate.Get(ctx, "Invalid ID [%s]", "")
			return 0, fmt.Errorf(msg, arg.(string))
		}
		return i, er
	}
	msg := translate.Get(ctx, "Invalid trip argument [%s]", "")
	return 0, fmt.Errorf(msg, arg.(string))
}

// func templatedHTML() string {
// 	where := paths.HTMLTemplates().String() + "templates.html"
// 	logger.Info.Printf("[TEMPLATE] Template Loc=[%v]", where)
// 	return where
// }

// htmlToBool converts a string representation of a boolean value to its corresponding boolean value.
// It takes a string s as input and returns a boolean value.
// The string s should be either "true" or "false".
// If s is "true", the function returns true.
// If s is "false", the function returns false.
// If s is neither "true" nor "false", the function behavior is undefined.
func htmlToBool(s string) bool {
	return htmlHelpers.ValueToBool(s)
}

func htmlToInt(s string) int {
	return htmlHelpers.ValueToInt(s)
}

func up(in string) string {
	return strings.ToUpper(in)
}

func Upper(in string) string {
	return strings.ToUpper(in)
}

func getTemplate(title, action string) string {
	tmpl := paths.HTMLTemplates().String() + "templates" + string(os.PathSeparator) + strings.ToLower(title) + "_" + strings.ToLower(action) + ".html"
	return tmpl
}
