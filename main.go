package main

import (
	"net/http"
	"strings"

	"github.com/julienschmidt/httprouter"
	common "github.com/mt1976/frantic-core/commonConfig"
	dockerhelpers "github.com/mt1976/frantic-core/dockerHelpers"
	logger "github.com/mt1976/frantic-core/logHandler"
	"github.com/mt1976/frantic-core/stringHelpers"
	"github.com/mt1976/frantic-core/timing"
	"github.com/mt1976/trnsl8r_service/app/business/translation"
	"github.com/mt1976/trnsl8r_service/app/dao"
	"github.com/mt1976/trnsl8r_service/app/dao/textStore"
	"github.com/mt1976/trnsl8r_service/app/jobs"
	"github.com/mt1976/trnsl8r_service/app/web/routes"
)

var settings *common.Settings

func init() {
	settings = common.Get()
}

func main() {

	err := dockerhelpers.DeployDefaultsPayload()
	if err != nil {
		logger.ErrorLogger.Fatal(err.Error())
	}

	err = error(nil)
	appName := settings.GetApplicationName()
	logger.InfoLogger.Printf("[%v] Starting...", appName)
	logger.InfoLogger.Printf("[%v] Connecting...", appName)
	err = dao.Initialise(settings)
	if err != nil {
		logger.ErrorLogger.Fatal(err.Error())
	}
	logger.InfoLogger.Printf("[%v] Connected", appName)
	logger.ServiceLogger.Printf("[%v] Backup Starting...", appName)

	err = jobs.DatabaseBackup.Run()
	if err != nil {
		logger.ErrorLogger.Fatal(err.Error())
	}

	err = jobs.DatabasePrune.Run()
	if err != nil {
		logger.PanicLogger.Fatal(err.Error())
	}

	logger.ServiceLogger.Printf("[%v] Backup Done", appName)

	logger.InfoLogger.Printf("[%v] Starting...", appName)
	setupSystemUser()

	na := strings.ToUpper(appName)

	timer := timing.Start(na, "Initialise", "Service")

	logger.InfoBanner(na, "Initialise", "Start...")

	// Preload the text store
	logger.InfoBanner(na, "Texts", "Importing")
	err = textStore.ImportCSV()
	if err != nil {
		logger.ErrorLogger.Fatal(err.Error())
	}

	logger.InfoBanner(na, "Texts", "Imported")

	logger.InfoBanner(na, "Texts", "Upgrading")
	err = jobs.LocaleUpdate.Run()
	if err != nil {
		logger.ErrorLogger.Fatal(err.Error())
	}

	logger.InfoBanner(na, "Texts", "Upgraded")

	logger.InfoBanner(na, "Initialise", "Done")

	logger.InfoBanner(na, "Routes", "Setup")

	router := httprouter.New()
	router = routes.Setup(router)

	// ANNOUNCE ROUTES ABOVE
	router.NotFound = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, translation.Get("404 page not found", ""), http.StatusNotFound)
	})

	//logger.InfoLogger.Println("APP: Routes Setup")
	logger.InfoBanner(na, "Routes", "Done")

	// Start the job processor
	jobs.Start()

	// test get of "🙁 ERROR [%v]"

	msg := "🙁 ERROR [%v]"
	newFunction(msg)
	newFunction(settings.GetApplicationDescription())
	newFunction(appName)

	port := settings.GetServerPortAsString()
	hostMachine := "localhost"
	protocol := settings.GetServerProtocol()

	logger.InfoLogger.Printf("[%v] Starting Server Port=[%v]", na, port)

	timer.Stop(1)

	logger.InfoLogger.Printf("[%v] Listening on %v://%v:%v/", na, protocol, hostMachine, port)
	logger.ErrorLogger.Fatal(http.ListenAndServe(":"+port, router))
}

func newFunction(msg string) {
	logger.TranslationLogger.Println("Translating: ", stringHelpers.DChevrons(translation.Get(msg, "")))

	// Get a list of the locales
	localList := settings.GetLocales()
	for _, locale := range localList {
		logger.TranslationLogger.Println(locale.Name + " " + stringHelpers.SBracket(translation.Get(msg, locale.Key)))
	}
}

func setupSystemUser() {
	logger.InfoBanner("System", "Users", "Setup")
	// Create the system user

	sysUCode := "sys"
	sysUName := "service"

	logger.InfoLogger.Printf("System User [%v] [%v] Available", sysUName, sysUCode)
}
