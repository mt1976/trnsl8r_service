package routes

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/julienschmidt/httprouter"
	"github.com/mt1976/frantic-core/commonConfig"
	"github.com/mt1976/frantic-core/logHandler"
)

func Setup(router *httprouter.Router) *httprouter.Router {

	//sessionID := "sessionID"

	logHandler.EventLogger.Println("Declare Behaviours ")

	logHandler.EventLogger.Println("Setup Routes ")

	// router.GET(announce("/"), Launch)
	router.ServeFiles(announceInsecure("/static/*filepath"), http.Dir("res/html/static"))
	router.ServeFiles(announceInsecure("/img/*filepath"), http.Dir("res/img"))
	router.ServeFiles(announceInsecure("/css/*filepath"), http.Dir("res/css"))
	router.ServeFiles(announceInsecure("/js/*filepath"), http.Dir("res/js"))
	router.ServeFiles(announceInsecure("/res/*filepath"), http.Dir("res/"))

	router.GET(announceInsecure("/"), TextList)
	router.GET(announceInsecure("/home"), TextList)

	router.GET(announceInsecure("/texts"), TextList) // List Text
	router.GET(announceInsecure("/text/vw/:id"), TextView)
	router.GET(announceInsecure("/text/ed/:id"), TextEdit)
	router.GET(announceInsecure("/text/up/:id"), TextUpdate)

	router.GET(announceInsecure("/health"), Health)
	router.GET(announceInsecure("/fail"), Fail)
	router.GET(announceInsecure("/export"), Trnsl8r_Export) // Not Required

	router.GET(announceInsecure("/trnsl8r/:origin/:message"), Trnsl8r)
	router.GET(announceInsecure("/testit"), Trnsl8r_Test)
	router.GET(announceInsecure("/refresh"), Trnsl8r_Refresh)
	router.GET(announceInsecure("/rebuild"), Trnsl8r_Rebuild)

	router.GET(announceInsecure("/locales"), Locales)
	router.GET(announceInsecure("/localestest"), LocaleTest)

	// Special Routes
	if settings.IsApplicationMode(commonConfig.MODE_DEVELOPMENT) {
		router.GET(announceInsecure("/test"), Test)
	}

	return router

}

func announceInsecure(route string) string {

	prefix := fmt.Sprintf("[ROUTE] Path=[%v]", route)
	padTo := 50
	if len(prefix) < padTo {
		prefix = prefix + strings.Repeat(" ", padTo-len(prefix))
	}
	logHandler.ApiLogger.Printf("%s %v://%v:%v%v", prefix, serverProtocol, serverHost, serverPort, route)
	return route
}

// func announceSecure(route string) string {
// 	route = route + "/:" + settings.SecuritySessionKey()
// 	return announceInsecure(route)
// }

func Initialise(cfg *commonConfig.Settings) {
	logHandler.EventLogger.Println("Initialise Routes")
}
