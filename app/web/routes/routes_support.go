package routes

import (
	"fmt"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/julienschmidt/httprouter"
	html "github.com/mt1976/frantic-core/htmlHelpers"
	logger "github.com/mt1976/frantic-core/logHandler"
	"github.com/mt1976/frantic-core/paths"
	"github.com/mt1976/trnsl8r_service/app/business/translation"
)

func trace(r *http.Request) {
	mesg := translation.Get("Method=[%s] URI=[%s] Header[%v] Context[%+v] RequestURI=[%s]", "")
	logger.TraceLogger.Printf(mesg, r.Method, r.URL, r.Header, r.Context(), r.RequestURI)
}

func getID(ps httprouter.Params) (int, error) {
	id, err := iconv(ps.ByName("id"))
	if err != nil {
		return 0, err
	}
	return id, nil
}

func getIDString(ps httprouter.Params) (string, error) {
	id := ps.ByName("id")
	return id, nil
}

func getItemName(ps httprouter.Params) (string, error) {
	itemName := ps.ByName("item")
	if itemName == "" {
		msg := translation.Get("Invalid Item Name [%s]", "")
		return "", fmt.Errorf(msg, itemName)
	}
	return itemName, nil
}

func iconv(arg interface{}) (int, error) {
	switch x := arg.(type) {
	case int:
		return x, nil
	case string:
		i, er := strconv.Atoi(x)
		if er != nil {
			msg := translation.Get("Invalid ID [%s]", "")
			return 0, fmt.Errorf(msg, arg.(string))
		}
		return i, er
	}
	msg := translation.Get("Invalid trip argument [%s]", "")
	return 0, fmt.Errorf(msg, arg.(string))
}

// func templatedHTML() string {
// 	where := paths.HTMLTemplates().String() + "templates.html"
// 	logger.InfoLogger.Printf("[TEMPLATE] Template Loc=[%v]", where)
// 	return where
// }

// htmlToBool converts a string representation of a boolean value to its corresponding boolean value.
// It takes a string s as input and returns a boolean value.
// The string s should be either "true" or "false".
// If s is "true", the function returns true.
// If s is "false", the function returns false.
// If s is neither "true" nor "false", the function behavior is undefined.
func htmlToBool(s string) bool {
	return html.ValueToBool(s)
}

func htmlToInt(s string) int {
	return html.ValueToInt(s)
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
