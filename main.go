package main

import (
	"net/http"
	"strings"

	"github.com/julienschmidt/httprouter"
	"github.com/mt1976/frantic-core/commonConfig"
	"github.com/mt1976/frantic-core/dao/actions"
	dockerHelpers "github.com/mt1976/frantic-core/dockerHelpers"
	"github.com/mt1976/frantic-core/logHandler"
	"github.com/mt1976/frantic-core/stringHelpers"
	"github.com/mt1976/frantic-core/timing"
	"github.com/mt1976/trnsl8r_service/app/business/translate"
	"github.com/mt1976/trnsl8r_service/app/dao"
	"github.com/mt1976/trnsl8r_service/app/dao/textstore"
	"github.com/mt1976/trnsl8r_service/app/jobs"
	"github.com/mt1976/trnsl8r_service/app/web/routes"
)

var settings *commonConfig.Settings

func init() {
	settings = commonConfig.Get()
}

func main() {

	// Deploy the default payload
	err := dockerHelpers.DeployDefaultsPayload()
	if err != nil {
		logHandler.ErrorLogger.Fatal(err.Error())
	}

	err = error(nil)
	appName := settings.GetApplication_Name()
	logHandler.InfoLogger.Printf("[%v] Starting...", appName)

	logHandler.InfoLogger.Printf("[%v] Connecting...", appName)
	err = dao.Initialise(settings)
	if err != nil {
		logHandler.ErrorLogger.Fatal(err.Error())
	}

	// //textstore.Initialise(context.TODO())
	// xx0, err := textstore.Count()
	// if err != nil {
	// 	logHandler.ErrorLogger.Fatal(err.Error())
	// }
	// logHandler.InfoLogger.Printf("[%v] Texts [%v]", appName, xx0)

	logHandler.InfoLogger.Printf("[%v] Connected", appName)
	logHandler.ServiceLogger.Printf("[%v] Backup Starting...", appName)

	// Add the functions DBs to the job before running
	jobs.DatabaseBackup.AddDatabaseAccessFunctions(textstore.FetchDatabaseInstances())

	err = jobs.DatabaseBackup.Run()
	if err != nil {
		logHandler.ErrorLogger.Fatal(err.Error())
	}

	//textstore.Initialise(context.TODO())

	// xx, err := textstore.Count()
	// if err != nil {
	// 	logHandler.ErrorLogger.Fatal(err.Error())
	// }
	// logHandler.InfoLogger.Printf("[%v] Texts [%v]", appName, xx)

	err = jobs.DatabasePrune.Run()
	if err != nil {
		logHandler.PanicLogger.Fatal(err.Error())
	}

	logHandler.ServiceLogger.Printf("[%v] Backup Done", appName)

	logHandler.InfoLogger.Printf("[%v] Starting...", appName)
	setupSystemUser()

	na := strings.ToUpper(appName)

	timer := timing.Start("", actions.INITIALISE.GetCode(), appName)

	//textstore.Initialise(context.TODO())

	// xx2, err := textstore.Count()
	// if err != nil {
	// 	logHandler.ErrorLogger.Fatal(err.Error())
	// }
	// logHandler.InfoLogger.Printf("[%v] Texts [%v]", na, xx2)

	logHandler.InfoBanner(na, "Initialise", "Start...")

	// Preload the text store
	logHandler.InfoBanner(na, "Texts", "Importing")
	err = textstore.ImportCSV()
	if err != nil {
		logHandler.ErrorLogger.Fatal(err.Error())
	}

	logHandler.InfoBanner(na, "Texts", "Imported")

	logHandler.InfoBanner(na, "Texts", "Upgrading")
	err = jobs.LocaleUpdate.Run()
	if err != nil {
		logHandler.ErrorLogger.Fatal(err.Error())
	}

	logHandler.InfoBanner(na, "Texts", "Upgraded")

	logHandler.InfoBanner(na, "Initialise", "Done")

	logHandler.InfoBanner(na, "Routes", "Setup")

	router := httprouter.New()
	router = routes.Setup(router)

	// ANNOUNCE ROUTES ABOVE
	router.NotFound = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, translate.Get("404 page not found", ""), http.StatusNotFound)
	})

	//logger.InfoLogger.Println("APP: Routes Setup")
	logHandler.InfoBanner(na, "Routes", "Done")

	// Start the job processor
	jobs.Start()

	// test get of "🙁 ERROR [%v]"

	msg := "🙁 ERROR [%v]"
	newFunction(msg)
	newFunction(settings.GetApplication_Description())
	newFunction(appName)

	port := settings.GetServer_PortString()
	hostMachine := "localhost"
	protocol := settings.GetServer_Protocol()

	logHandler.InfoLogger.Printf("[%v] Starting Server Port=[%v]", na, port)

	timer.Stop(1)

	logHandler.InfoLogger.Printf("[%v] Listening on %v://%v:%v/", na, protocol, hostMachine, port)
	logHandler.ErrorLogger.Fatal(http.ListenAndServe(":"+port, router))
}

func newFunction(msg string) {
	logHandler.TranslationLogger.Println("Translating: ", stringHelpers.DChevrons(translate.Get(msg, "")))

	// Get a list of the locales
	localList := settings.GetTranslation_PermittedLocales()
	for _, locale := range localList {
		logHandler.TranslationLogger.Println(locale.Name + " " + stringHelpers.SBracket(translate.Get(msg, locale.Key)))
	}
}

func setupSystemUser() {
	logHandler.InfoBanner("System", "Users", "Setup")
	// Create the system user

	sysUCode := "sys"
	sysUName := "service"

	logHandler.InfoLogger.Printf("System User [%v] [%v] Available", sysUName, sysUCode)
}
