package routes

import (
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/mt1976/trnsl8r_service/app/business/translation"

	trnsl8r "github.com/mt1976/trnsl8r_connect"
	"github.com/mt1976/trnsl8r_service/app/dao/textStore"
	"github.com/mt1976/trnsl8r_service/app/support/logger"
	"github.com/mt1976/trnsl8r_service/app/support/timing"
)

func Trnsl8r(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	itemToTranslate := ps.ByName("message")

	watch := timing.Start("Trnsl8r", "Translate", itemToTranslate)

	logger.TranslationLogger.Println("Request to translate message {{", itemToTranslate, "}}")

	if itemToTranslate == "" {
		err := fmt.Errorf("No message to translate")
		logger.ErrorLogger.Println(err.Error())
		watch.Stop(0)
		oops(w, r, nil, "error", err.Error())
	}

	translatedItem := translation.Get(itemToTranslate)

	if translatedItem == "" {
		err := fmt.Errorf("No translation available")
		logger.ErrorLogger.Println(err.Error())
		watch.Stop(0)
		oops(w, r, nil, "error", err.Error())
	}

	// Respond with the translated item and a success status
	w.Header().Set("Content-Type", "application/json")
	w.Header().Add("Application", cfg.ApplicationName())
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "{\"message\":\"%v\"}", translatedItem)
	watch.Stop(1)
}

func Trnsl8r_Test(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	// Build a URI query string
	tl8 := trnsl8r.NewRequest().WithProtocol(cfg.TranslationProtocol()).WithHost(cfg.TranslationHost()).WithPort(cfg.TranslationPort()).WithLogger(logger.TranslationLogger).WithOriginOf(cfg.ApplicationName())

	tl8.Spew()

	logger.TranslationLogger.Println("Request to translate message {{", tl8.String(), "}}")

	all, err := textStore.GetAll()
	if err != nil {
		logger.ErrorLogger.Println(err.Error())
	}

	for _, item := range all {
		logger.TranslationLogger.Println("Original: {{", item.Original, "}}")
		translation, err := tl8.Get(item.Original)
		if err != nil {
			logger.ErrorLogger.Println(err.Error())
		}
		logger.InfoLogger.Println("Original: {{", item.Original, "}} Translation: {{", translation.String(), "}}", "Information: {{", translation.Information, "}}")
	}
}
