package routes

import (
	"fmt"
	"net/http"
	"slices"
	"strings"

	"github.com/julienschmidt/httprouter"
	"github.com/mt1976/trnsl8r_service/app/business/domains"
	"github.com/mt1976/trnsl8r_service/app/business/translate"

	"github.com/mt1976/frantic-core/commonConfig"
	"github.com/mt1976/frantic-core/htmlHelpers"
	"github.com/mt1976/frantic-core/idHelpers"
	"github.com/mt1976/frantic-core/logHandler"
	"github.com/mt1976/frantic-core/stringHelpers"
	"github.com/mt1976/frantic-core/timing"
	trnsl8r "github.com/mt1976/trnsl8r_connect"
	"github.com/mt1976/trnsl8r_service/app/dao/textstore"
)

func Trnsl8r(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	c := commonConfig.Get()

	itemToTranslate := ps.ByName("message")
	originOfRequest := ps.ByName("origin")
	localesList := getLocalesList(c)

	filterLocale := r.URL.Query().Get(trnsl8r.LOCALE.Key())
	//fmt.Print("filterLocale1: ", filterLocale, "\n")

	if filterLocale != "" {
		filterLocale = getLocale(filterLocale, w, r, localesList)
	}

	//fmt.Print("filterLocale5: ", filterLocale, "\n")

	//logHandler.TranslationLogger.Println("Filtering by locale [", filterLocale, "]")

	watch := timing.Start("Trnsl8r", "Translate", itemToTranslate)

	logHandler.TranslationLogger.Println("Request to translate message [", itemToTranslate, "], origin [", originOfRequest, "], locale [", filterLocale, "]")

	if itemToTranslate == "" {
		err := fmt.Errorf("no message to translate")
		logHandler.ErrorLogger.Println(err.Error())
		watch.Stop(0)
		oops(w, r, nil, "error", err.Error())
		return
	}

	// Needs to be decoded from the URL
	itemToTranslate, err := htmlHelpers.FromPathSafe(itemToTranslate)
	if err != nil {
		logHandler.ErrorLogger.Println(err.Error())
		oops(w, r, nil, "error", err.Error())
		return
	}

	realOrigin := idHelpers.GetUUIDv2Payload(originOfRequest)

	if originOfRequest == "" || realOrigin == "" {
		err := fmt.Errorf("no origin of request, a valid origin is required %v", originOfRequest)
		logHandler.ErrorLogger.Println(err.Error())
		watch.Stop(0)
		oops(w, r, nil, "error", err.Error())
		return
	}

	if !slices.Contains(c.GetTranslation_PermittedOrigins(), realOrigin) {
		err := fmt.Errorf("invalid origin of request [%v]", realOrigin)
		logHandler.ErrorLogger.Println(err.Error())
		watch.Stop(0)
		oops(w, r, nil, "error", err.Error())
		return
	}

	translatedItem := translate.Get(itemToTranslate, filterLocale)

	if translatedItem == "" {
		err := fmt.Errorf("no translation available")
		logHandler.ErrorLogger.Println(err.Error())
		watch.Stop(0)
		oops(w, r, nil, "error", err.Error())
		return
	}

	// Respond with the translated item and a success status
	w.Header().Set("Content-Type", "application/json")
	w.Header().Add("Application", appName)

	//	fmt.Fprintf(w, "{\"message\":\"%v\"}", translatedItem)
	htmlResp := fmt.Sprintf("{\"message\":\"%v\"}", translatedItem)
	w.Write([]byte(htmlResp))
	w.WriteHeader(http.StatusOK)
	logHandler.TranslationLogger.Println("Response to translate message [", htmlResp, "] Status=", http.StatusOK)
	//logHandler.InfoLogger.Printf("Translated message [%v] to [%v]", itemToTranslate, translatedItem)
	watch.Stop(1)
}

func getLocale(filterLocale string, w http.ResponseWriter, r *http.Request, localesList []string) string {
	//fmt.Print("filterLocale2: ", filterLocale, "\n")

	filterLocale = strings.Trim(filterLocale, " ")
	filterLocale, err := htmlHelpers.FromPathSafe(filterLocale)
	if err != nil {
		logHandler.ErrorLogger.Println(err.Error())
		oops(w, r, nil, "error", err.Error())
		return ""
	}
	//fmt.Print("filterLocale3: ", filterLocale, "\n")

	// Check that this is a valid locale
	if !slices.Contains(localesList, filterLocale) {
		err := fmt.Errorf("invalid locale [%v]", filterLocale)
		logHandler.ErrorLogger.Println(err.Error())
		oops(w, r, nil, "error", err.Error())
		return ""
	}
	//fmt.Print("filterLocale4: ", filterLocale, "\n")
	return filterLocale
}

func Trnsl8r_Test(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	// Build a URI query string
	baseReq := trnsl8r.NewRequest().WithProtocol(trnsServerProtocol).WithHost(trnsServerHost).WithPort(trnsServerPort).WithLogger(logHandler.ServiceLogger).FromOrigin("trnsl8r_connect")

	//baseReq.Spew()

	logHandler.TranslationLogger.Println("Request to translate message ", stringHelpers.DCurlies(baseReq.String()))

	all, err := textstore.GetAll()
	if err != nil {
		logHandler.ErrorLogger.Println(err.Error())
	}

	locales := getLocalesList(commonConfig.Get())

	for _, item := range all {
		//	logHandler.TranslationLogger.Println("Original: ", stringHelpers.DCurlies(item.Original))

		translation, err := baseReq.Get(item.Original)
		if err != nil {
			logHandler.ErrorLogger.Println(err.Error())
		}
		logHandler.EventLogger.Println("Original:", stringHelpers.DCurlies(item.Original), " Translation:", stringHelpers.DCurlies(translation.String()), "Information:", stringHelpers.DCurlies(translation.Information), "Locale:", stringHelpers.DCurlies(""))

		for _, locale := range locales {
			useReq, err := baseReq.WithLocale(locale)
			if err != nil {
				logHandler.ErrorLogger.Println(err.Error())
			}
			//useReq.Spew()
			translation, err := useReq.Get(item.Original)
			if err != nil {
				logHandler.ErrorLogger.Println(err.Error())
			}
			logHandler.EventLogger.Println("Original:", stringHelpers.DCurlies(item.Original), " Translation:", stringHelpers.DCurlies(translation.String()), "Information:", stringHelpers.DCurlies(translation.Information), "Locale:", stringHelpers.DCurlies(locale))
		}
	}
}

func Trnsl8r_Export(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	//logger.EventLogger.Printf("[TEST] [View]")

	trace(r)

	err := textstore.ExportCSV()
	if err != nil {
		logHandler.ErrorLogger.Print(err.Error())
		oops(w, r, nil, "error", err.Error())
	}
	successMessage(w, r, nil, "success - translations exported")
	logHandler.EventLogger.Printf("[TEST] [Export] [Translations] [Success]")
}

func Trnsl8r_Refresh(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	logHandler.InfoBanner(domains.TEXT.String(), "Texts", "Importing")
	err := textstore.ImportCSV()
	if err != nil {
		logHandler.ErrorLogger.Fatal(err.Error())
		oops(w, r, nil, "error", err.Error())
	}

	logHandler.InfoBanner(domains.TEXT.String(), "Texts", "Imported")
	successMessage(w, r, nil, "success - translations imported")
}

func Trnsl8r_Rebuild(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	// Need to drop the existing text store
	err := textstore.Drop()
	if err != nil {
		logHandler.ErrorLogger.Fatal(err.Error())
		oops(w, r, nil, "error", err.Error())
	}

	logHandler.InfoBanner(domains.TEXT.String(), "Texts", "Importing")
	err = textstore.ImportCSV()
	if err != nil {
		logHandler.ErrorLogger.Fatal(err.Error())
		oops(w, r, nil, "error", err.Error())
	}

	logHandler.InfoBanner(domains.TEXT.String(), "Texts", "Imported")
	successMessage(w, r, nil, "success - translations imported")
}

func getLocalesList(c *commonConfig.Settings) []string {
	locales := []string{}
	all := c.GetTranslation_PermittedLocales()
	for _, item := range all {
		locales = append(locales, item.Key)
	}
	return locales
}
