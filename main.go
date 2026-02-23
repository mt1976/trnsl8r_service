package main

import (
	"context"
	"net/http"
	"strings"

	"github.com/julienschmidt/httprouter"
	"github.com/mt1976/frantic-core/commonConfig"
	dockerHelpers "github.com/mt1976/frantic-core/dockerHelpers"
	"github.com/mt1976/frantic-core/logHandler"
	"github.com/mt1976/frantic-core/stringHelpers"
	"github.com/mt1976/frantic-core/timing"
	"github.com/mt1976/trnsl8r_service/app/business/translate"
	"github.com/mt1976/trnsl8r_service/app/dao"
	"github.com/mt1976/trnsl8r_service/app/dao/textStore"
	"github.com/mt1976/trnsl8r_service/app/jobs"
	"github.com/mt1976/trnsl8r_service/app/web/routes"
)

var settings *commonConfig.Settings

func init() {
	settings = commonConfig.Get()
}

func main() {
	ctx := context.Background()

	// Deploy the default payload
	err := dockerHelpers.DeployDefaultsPayload()
	if err != nil {
		logHandler.Error.Fatal(err.Error())
	}

	err = error(nil)
	appName := settings.GetApplication_Name()
	logHandler.Info.Printf("[%v] Starting...", appName)

	logHandler.Info.Printf("[%v] Connecting...", appName)
	err = dao.Initialise(settings)
	if err != nil {
		logHandler.Error.Fatal(err.Error())
	}

	// //textstore.Initialise(context.TODO())
	// xx0, err := textstore.Count()
	// if err != nil {
	// 	logHandler.Error.Fatal(err.Error())
	// }
	// logHandler.Info.Printf("[%v] Texts [%v]", appName, xx0)

	logHandler.Info.Printf("[%v] Connected", appName)
	logHandler.Service.Printf("[%v] Backup Starting...", appName)

	// Add the functions DBs to the job before running
	jobs.DatabaseBackup.AddDatabaseAccessFunctions(textStore.GetDatabaseConnections())

	err = jobs.DatabaseBackup.Run()
	if err != nil {
		logHandler.Error.Fatal(err.Error())
	}

	// textstore.Initialise(context.TODO())

	// xx, err := textstore.Count()
	// if err != nil {
	// 	logHandler.Error.Fatal(err.Error())
	// }
	// logHandler.Info.Printf("[%v] Texts [%v]", appName, xx)

	err = jobs.DatabasePrune.Run()
	if err != nil {
		logHandler.Panic.Fatal(err.Error())
	}

	logHandler.Service.Printf("[%v] Backup Done", appName)

	logHandler.Info.Printf("[%v] Starting...", appName)
	setupSystemUser()

	na := strings.ToUpper(appName)

	timer := timing.Start("", "Initialise", appName)

	// textstore.Initialise(context.TODO())

	// xx2, err := textstore.Count()
	// if err != nil {
	// 	logHandler.Error.Fatal(err.Error())
	// }
	// logHandler.Info.Printf("[%v] Texts [%v]", na, xx2)

	logHandler.Banner(na, "Initialise", "Start...")

	// Preload the text store
	logHandler.Banner(na, "Texts", "Importing")
	err = textStore.ImportDefaults(ctx)
	if err != nil {
		logHandler.Error.Fatal(err.Error())
	}

	logHandler.Banner(na, "Texts", "Imported")

	logHandler.Banner(na, "Texts", "Upgrading")
	err = jobs.LocaleUpdate.Run()
	if err != nil {
		logHandler.Error.Fatal(err.Error())
	}

	logHandler.Banner(na, "Texts", "Upgraded")

	logHandler.Banner(na, "Initialise", "Done")

	logHandler.Banner(na, "Routes", "Setup")

	router := httprouter.New()
	router = routes.Setup(router)

	// ANNOUNCE ROUTES ABOVE
	router.NotFound = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, translate.Get(ctx, "404 page not found", ""), http.StatusNotFound)
	})

	// logger.Info.Println("APP: Routes Setup")
	logHandler.Banner(na, "Routes", "Done")

	// Start the job processor
	jobs.Start()

	// test get of "🙁 ERROR [%v]"

	msg := "🙁 ERROR [%v]"
	tl8(ctx, msg)
	tl8(ctx, settings.GetApplication_Description())
	tl8(ctx, appName)

	port := settings.GetServer_PortString()
	hostMachine := "localhost"
	protocol := settings.GetServer_Protocol()

	logHandler.Info.Printf("[%v] Starting Server Port=[%v]", na, port)

	timer.Stop(1)

	logHandler.Info.Printf("[%v] Listening on %v://%v:%v/", na, protocol, hostMachine, port)
	logHandler.Error.Fatal(http.ListenAndServe(":"+port, router))
}

func tl8(ctx context.Context, msg string) {
	logHandler.Translation.Println("Translating: ", stringHelpers.DChevrons(translate.Get(ctx, msg, "")))

	// Get a list of the locales
	localList := settings.GetTranslation_PermittedLocales()
	for _, locale := range localList {
		logHandler.Translation.Println(locale.Name + " " + stringHelpers.SBracket(translate.Get(ctx, msg, locale.Key)))
	}
}

func setupSystemUser() {
	logHandler.Banner("System", "Users", "Setup")
	// Create the system user

	sysUCode := "sys"
	sysUName := "service"

	logHandler.Info.Printf("System User [%v] [%v] Available", sysUName, sysUCode)
}
