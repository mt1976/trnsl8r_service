package routes

import (
	"fmt"
	"net/http"
	"slices"

	"github.com/julienschmidt/httprouter"
	"github.com/mt1976/trnsl8r_service/app/business/domains"
	"github.com/mt1976/trnsl8r_service/app/business/translation"

	"github.com/mt1976/frantic-plum/common"
	h "github.com/mt1976/frantic-plum/html"
	"github.com/mt1976/frantic-plum/id"
	"github.com/mt1976/frantic-plum/logger"
	"github.com/mt1976/frantic-plum/timing"
	trnsl8r "github.com/mt1976/trnsl8r_connect"
	"github.com/mt1976/trnsl8r_service/app/dao/textStore"
)

func Trnsl8r(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	c := common.Get()

	itemToTranslate := ps.ByName("message")
	originOfRequest := ps.ByName("origin")
	localesList := getLocalesList(c)

	filterLocale := r.URL.Query().Get("locale")
	if filterLocale != "" {
		filterLocale, err := h.FromPathSafe(filterLocale)
		if err != nil {
			logger.ErrorLogger.Println(err.Error())
			oops(w, r, nil, "error", err.Error())
			return
		}
		logger.TranslationLogger.Println("Filtering by locale [", filterLocale, "]")
		// Check that this is a valid locale
		if !slices.Contains(localesList, filterLocale) {
			err := fmt.Errorf("invalid locale [%v]", filterLocale)
			logger.ErrorLogger.Println(err.Error())
			oops(w, r, nil, "error", err.Error())
			return
		}
	}

	watch := timing.Start("Trnsl8r", "Translate", itemToTranslate)

	logger.TranslationLogger.Println("Request to translate message [", itemToTranslate, "]")

	if itemToTranslate == "" {
		err := fmt.Errorf("no message to translate")
		logger.ErrorLogger.Println(err.Error())
		watch.Stop(0)
		oops(w, r, nil, "error", err.Error())
		return
	}

	// Needs to be decoded from the URL
	itemToTranslate, err := h.FromPathSafe(itemToTranslate)
	if err != nil {
		logger.ErrorLogger.Println(err.Error())
		oops(w, r, nil, "error", err.Error())
		return
	}

	realOrigin := id.GetUUIDv2Payload(originOfRequest)

	if originOfRequest == "" || realOrigin == "" {
		err := fmt.Errorf("no origin of request, a valid origin is required %v", originOfRequest)
		logger.ErrorLogger.Println(err.Error())
		watch.Stop(0)
		oops(w, r, nil, "error", err.Error())
		return
	}

	if !slices.Contains(c.GetValidOrigins(), realOrigin) {
		err := fmt.Errorf("invalid origin of request [%v]", realOrigin)
		logger.ErrorLogger.Println(err.Error())
		watch.Stop(0)
		oops(w, r, nil, "error", err.Error())
		return
	}

	translatedItem := translation.Get(itemToTranslate, filterLocale)

	if translatedItem == "" {
		err := fmt.Errorf("no translation available")
		logger.ErrorLogger.Println(err.Error())
		watch.Stop(0)
		oops(w, r, nil, "error", err.Error())
		return
	}

	// Respond with the translated item and a success status
	w.Header().Set("Content-Type", "application/json")
	w.Header().Add("Application", settings.ApplicationName())

	fmt.Fprintf(w, "{\"message\":\"%v\"}", translatedItem)
	w.WriteHeader(http.StatusOK)
	logger.InfoLogger.Println(fmt.Sprintf("Translated message [%v] to [%v]", itemToTranslate, translatedItem))
	watch.Stop(1)
}

func Trnsl8r_Test(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	// Build a URI query string
	tl8 := trnsl8r.NewRequest().WithProtocol(settings.TranslationProtocol()).WithHost(settings.TranslationHost()).WithPort(settings.TranslationPort()).WithLogger(logger.InfoLogger).WithOriginOf("trnsl8r_connect")

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

func Trnsl8r_Export(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	//logger.EventLogger.Printf("[TEST] [View]")

	trace(r)

	err := textStore.ExportCSV()
	if err != nil {
		logger.ErrorLogger.Print(err.Error())
		oops(w, r, nil, "error", err.Error())
	}
	successMessage(w, r, nil, "success - translations exported")
	logger.EventLogger.Printf("[TEST] [Export] [Translations] [Success]")
}

func Trnsl8r_Refresh(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	logger.Banner(domains.TEXT.String(), "Texts", "Importing")
	err := textStore.ImportCSV()
	if err != nil {
		logger.ErrorLogger.Fatal(err.Error())
		oops(w, r, nil, "error", err.Error())
	}

	logger.Banner(domains.TEXT.String(), "Texts", "Imported")
	successMessage(w, r, nil, "success - translations imported")
}

func Trnsl8r_Rebuild(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	// Need to drop the existing text store
	err := textStore.Drop()
	if err != nil {
		logger.ErrorLogger.Fatal(err.Error())
		oops(w, r, nil, "error", err.Error())
	}

	logger.Banner(domains.TEXT.String(), "Texts", "Importing")
	err = textStore.ImportCSV()
	if err != nil {
		logger.ErrorLogger.Fatal(err.Error())
		oops(w, r, nil, "error", err.Error())
	}

	logger.Banner(domains.TEXT.String(), "Texts", "Imported")
	successMessage(w, r, nil, "success - translations imported")
}

func getLocalesList(c *common.Settings) []string {
	locales := []string{}
	all := c.GetLocales()
	for _, item := range all {
		locales = append(locales, item.Key)
	}
	return locales
}
