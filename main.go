package main

import (
	"net/http"
	"strings"

	"github.com/julienschmidt/httprouter"
	"github.com/mt1976/trnsl8r_service/app/business/translation"
	"github.com/mt1976/trnsl8r_service/app/dao"
	"github.com/mt1976/trnsl8r_service/app/dao/textStore"
	"github.com/mt1976/trnsl8r_service/app/jobs"
	"github.com/mt1976/trnsl8r_service/app/support"
	"github.com/mt1976/trnsl8r_service/app/support/config"
	"github.com/mt1976/trnsl8r_service/app/support/logger"
	"github.com/mt1976/trnsl8r_service/app/support/timing"
	"github.com/mt1976/trnsl8r_service/app/web/routes"
)

var cfg *config.Configuration

func init() {
	cfg = config.Get()
}

func main() {

	err := error(nil)
	logger.InfoLogger.Printf("[%v] Starting...", cfg.ApplicationName())
	logger.InfoLogger.Printf("[%v] Connecting...", cfg.ApplicationName())
	err = dao.Initialise(cfg)
	if err != nil {
		logger.ErrorLogger.Fatal(err.Error())
	}
	logger.InfoLogger.Printf("[%v] Connected", cfg.ApplicationName())
	logger.ServiceLogger.Printf("[%v] Backup Starting...", cfg.ApplicationName())

	err = jobs.DatabaseBackup.Run()
	if err != nil {
		logger.ErrorLogger.Fatal(err.Error())
	}

	err = jobs.DatabasePrune.Run()
	if err != nil {
		logger.PanicLogger.Fatal(err.Error())
	}

	logger.ServiceLogger.Printf("[%v] Backup Done", cfg.ApplicationName())

	logger.InfoLogger.Printf("[%v] Starting...", cfg.ApplicationName())
	setupSystemUser()

	na := strings.ToUpper(cfg.ApplicationName())

	timer := timing.Start(na, "Initialise", "Service")

	support.Banner(na, "Initialise", "Start...")

	// Preload the text store
	support.Banner(na, "Texts", "Importing")
	err = textStore.ImportCSV()
	if err != nil {
		logger.ErrorLogger.Fatal(err.Error())
	}

	support.Banner(na, "Texts", "Imported")

	support.Banner(na, "Initialise", "Done")

	support.Banner(na, "Routes", "Setup")

	router := httprouter.New()
	router = routes.Setup(router)

	// ANNOUNCE ROUTES ABOVE
	router.NotFound = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, translation.Get("404 page not found"), http.StatusNotFound)
	})

	//logger.InfoLogger.Println("APP: Routes Setup")
	support.Banner(na, "Routes", "Done")

	// Start the job processor
	jobs.Start()

	port := cfg.ApplicationPortString()
	hostMachine := "localhost"
	protocol := cfg.ServerProtocol()

	logger.InfoLogger.Printf("[%v] Starting Server Port=[%v]", na, port)

	timer.Stop(1)

	logger.InfoLogger.Printf("[%v] Listening on %v://%v:%v/", na, protocol, hostMachine, port)
	logger.ErrorLogger.Fatal(http.ListenAndServe(":"+port, router))
}

func setupSystemUser() {
	support.Banner("System", "Users", "Setup")
	// Create the system user

	sysUCode := "sys"
	sysUName := "service"

	logger.InfoLogger.Printf("System User [%v] [%v] Available", sysUName, sysUCode)
}
