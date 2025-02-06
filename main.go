package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/julienschmidt/httprouter"
	"github.com/mt1976/frantic-plum/common"
	"github.com/mt1976/frantic-plum/logger"
	"github.com/mt1976/frantic-plum/stringHelpers"
	"github.com/mt1976/frantic-plum/timing"
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

	err := startup()
	if err != nil {
		logger.ErrorLogger.Fatal(err.Error())
	}

	err = error(nil)

	logger.InfoLogger.Printf("[%v] Starting...", settings.ApplicationName())
	logger.InfoLogger.Printf("[%v] Connecting...", settings.ApplicationName())
	err = dao.Initialise(settings)
	if err != nil {
		logger.ErrorLogger.Fatal(err.Error())
	}
	logger.InfoLogger.Printf("[%v] Connected", settings.ApplicationName())
	logger.ServiceLogger.Printf("[%v] Backup Starting...", settings.ApplicationName())

	err = jobs.DatabaseBackup.Run()
	if err != nil {
		logger.ErrorLogger.Fatal(err.Error())
	}

	err = jobs.DatabasePrune.Run()
	if err != nil {
		logger.PanicLogger.Fatal(err.Error())
	}

	logger.ServiceLogger.Printf("[%v] Backup Done", settings.ApplicationName())

	logger.InfoLogger.Printf("[%v] Starting...", settings.ApplicationName())
	setupSystemUser()

	na := strings.ToUpper(settings.ApplicationName())

	timer := timing.Start(na, "Initialise", "Service")

	logger.Banner(na, "Initialise", "Start...")

	// Preload the text store
	logger.Banner(na, "Texts", "Importing")
	err = textStore.ImportCSV()
	if err != nil {
		logger.ErrorLogger.Fatal(err.Error())
	}

	logger.Banner(na, "Texts", "Imported")

	logger.Banner(na, "Texts", "Upgrading")
	err = jobs.LocaleUpdate.Run()
	if err != nil {
		logger.ErrorLogger.Fatal(err.Error())
	}

	logger.Banner(na, "Texts", "Upgraded")

	logger.Banner(na, "Initialise", "Done")

	logger.Banner(na, "Routes", "Setup")

	router := httprouter.New()
	router = routes.Setup(router)

	// ANNOUNCE ROUTES ABOVE
	router.NotFound = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, translation.Get("404 page not found", ""), http.StatusNotFound)
	})

	//logger.InfoLogger.Println("APP: Routes Setup")
	logger.Banner(na, "Routes", "Done")

	// Start the job processor
	jobs.Start()

	// test get of "🙁 ERROR [%v]"

	msg := "🙁 ERROR [%v]"
	newFunction(msg)
	newFunction(settings.ApplicationDescription())
	newFunction(settings.ApplicationName())

	port := settings.ApplicationPortString()
	hostMachine := "localhost"
	protocol := settings.ServerProtocol()

	logger.InfoLogger.Printf("[%v] Starting Server Port=[%v]", na, port)

	timer.Stop(1)

	logger.InfoLogger.Printf("[%v] Listening on %v://%v:%v/", na, protocol, hostMachine, port)
	logger.ErrorLogger.Fatal(http.ListenAndServe(":"+port, router))
}

func newFunction(msg string) {
	logger.InfoLogger.Println(stringHelpers.DChevrons(translation.Get(msg, "")))

	// Get a list of the locales
	localList := settings.GetLocales()
	for _, locale := range localList {
		logger.InfoLogger.Println(locale.Name + " " + stringHelpers.SBracket(translation.Get(msg, locale.Key)))
	}
}

func setupSystemUser() {
	logger.Banner("System", "Users", "Setup")
	// Create the system user

	sysUCode := "sys"
	sysUName := "service"

	logger.InfoLogger.Printf("System User [%v] [%v] Available", sysUName, sysUCode)
}

func startup() error {
	list, err := os.ReadDir("defaults")
	if err != nil {
		log.Fatal(err)
	}

	for _, entry := range list {
		if entry.IsDir() {
			continue
		}
		if strings.HasSuffix(entry.Name(), ".toml") {
			continue
		}
		if entry.Name() == ".DS_Store" {
			continue
		}
		if entry.Name() == ".keep" {
			continue
		}

		//logger.InfoLogger.Printf("Copying %v", entry.Name())
		from := "./defaults" + string(os.PathSeparator) + entry.Name()
		to := "./data" + string(os.PathSeparator) + "defaults" + string(os.PathSeparator) + entry.Name()
		logger.EventLogger.Printf("Copying [%v] to [%v]", from, to)
		err = startupCopyFile(from, to)
		if err != nil {
			logger.ErrorLogger.Println(err.Error())
		}
		//err = CopyFile("defaults/defaults.toml", "data/defaults/defaults.toml")

	}
	return err
}

// startupCopyFile copies a file from src to dst. If src and dst files exist, and are
// the same, then return success. Otherise, attempt to create a hard link
// between the two files. If that fail, copy the file contents from src to dst.
func startupCopyFile(src, dst string) (err error) {
	sfi, err := os.Stat(src)
	if err != nil {
		return
	}
	if !sfi.Mode().IsRegular() {
		// cannot copy non-regular files (e.g., directories,
		// symlinks, devices, etc.)
		return fmt.Errorf("CopyFile: non-regular source file %s (%q)", sfi.Name(), sfi.Mode().String())
	}
	dfi, err := os.Stat(dst)
	if err != nil {
		if !os.IsNotExist(err) {
			return
		}
	} else {
		if !(dfi.Mode().IsRegular()) {
			return fmt.Errorf("CopyFile: non-regular destination file %s (%q)", dfi.Name(), dfi.Mode().String())
		}
		if os.SameFile(sfi, dfi) {
			return
		}
	}
	if err = os.Link(src, dst); err == nil {
		return
	}
	err = startupCopyFileContents(src, dst)
	return
}

// startupCopyFileContents copies the contents of the file named src to the file named
// by dst. The file will be created if it does not already exist. If the
// destination file exists, all it's contents will be replaced by the contents
// of the source file.
func startupCopyFileContents(src, dst string) (err error) {
	logger.EventLogger.Printf("Copying [%v] to [%v]", src, dst)
	in, err := os.Open(src)
	if err != nil {
		return
	}
	defer in.Close()
	out, err := os.Create(dst)
	if err != nil {
		return
	}
	defer func() {
		cerr := out.Close()
		if err == nil {
			err = cerr
		}
	}()
	if _, err = io.Copy(out, in); err != nil {
		return
	}
	err = out.Sync()
	return
}
