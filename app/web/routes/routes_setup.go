package routes

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/julienschmidt/httprouter"
	"github.com/mt1976/trnsl8r_service/app/support/config"
	"github.com/mt1976/trnsl8r_service/app/support/logger"
)

var cfg *config.Configuration

func init() {
	logger.EventLogger.Println("Loading Routes")
	cfg = config.Get()
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
	router.GET(announceInsecure("/ExportTranslations"), ExportTranslations) // Not Required

	router.GET(announceInsecure("/trnsl8r/translate/:message"), Trnsl8r)
	router.GET(announceInsecure("/trnsl8r/test"), Trnsl8r_Test)

	// Special Routes
	if cfg.ApplicationModeIs(config.MODE_DEVELOPMENT) {
		router.GET(announceInsecure("/test"), Test)
	}

	return router

}

func announceInsecure(route string) string {
	port := cfg.ApplicationPortString()
	hostMachine := "localhost"
	protocol := cfg.ServerProtocol()

	prefix := fmt.Sprintf("[ROUTE] Path=[%v]", route)
	padTo := 50
	if len(prefix) < padTo {
		prefix = prefix + strings.Repeat(" ", padTo-len(prefix))
	}
	logger.ApiLogger.Printf("%s %v://%v:%v%v", prefix, protocol, hostMachine, port, route)
	return route
}

// func announceSecure(route string) string {
// 	route = route + "/:" + cfg.SecuritySessionKey()
// 	return announceInsecure(route)
// }

func Initialise(cfg *config.Configuration) {
	logger.EventLogger.Println("Initialise Routes")
}
