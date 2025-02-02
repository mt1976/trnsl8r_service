package routes

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/julienschmidt/httprouter"
	"github.com/mt1976/frantic-plum/common"
	"github.com/mt1976/frantic-plum/logger"
)

var settings *common.Settings

func init() {
	logger.EventLogger.Println("Loading Routes")
	settings = common.Get()
}

func Setup(router *httprouter.Router) *httprouter.Router {

	//sessionID := "sessionID"

	logger.EventLogger.Println("Declare Behaviours ")

	logger.EventLogger.Println("Setup Routes ")

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

	// Special Routes
	if settings.ApplicationModeIs(common.MODE_DEVELOPMENT) {
		router.GET(announceInsecure("/test"), Test)
	}

	return router

}

func announceInsecure(route string) string {
	port := settings.ApplicationPortString()
	hostMachine := "localhost"
	protocol := settings.ServerProtocol()

	prefix := fmt.Sprintf("[ROUTE] Path=[%v]", route)
	padTo := 50
	if len(prefix) < padTo {
		prefix = prefix + strings.Repeat(" ", padTo-len(prefix))
	}
	logger.ApiLogger.Printf("%s %v://%v:%v%v", prefix, protocol, hostMachine, port, route)
	return route
}

// func announceSecure(route string) string {
// 	route = route + "/:" + settings.SecuritySessionKey()
// 	return announceInsecure(route)
// }

func Initialise(cfg *common.Settings) {
	logger.EventLogger.Println("Initialise Routes")
}
