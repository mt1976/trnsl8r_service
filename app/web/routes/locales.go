package routes

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"

	"github.com/mt1976/frantic-core/commonConfig"
	"github.com/mt1976/frantic-core/htmlHelpers"
	"github.com/mt1976/frantic-core/logHandler"
	"github.com/mt1976/frantic-core/stringHelpers"
	"github.com/mt1976/frantic-core/timing"
	trnsl8r "github.com/mt1976/trnsl8r_connect"
)

type LocaleResponse struct {
	Locales []struct {
		Locale string `json:"locale"`
		Name   string `json:"name"`
	} `json:"locales"`
	Message string `json:"message"`
}

func Locales(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	c := commonConfig.Get()

	localeResp := LocaleResponse{}

	//originOfRequest := ps.ByName("origin")

	watch := timing.Start("Trnsl8r", "Locales", "All Locales")

	// realOrigin := idHelpers.GetUUIDv2Payload(originOfRequest)

	// if originOfRequest == "" || realOrigin == "" {
	// 	err := fmt.Errorf("no origin of request, a valid origin is required %v", originOfRequest)
	// 	logHandler.ErrorLogger.Println(err.Error())
	// 	watch.Stop(0)
	// 	oops(w, r, nil, "error", err.Error())
	// 	return
	// }

	// if !slices.Contains(c.GetTranslation_PermittedOrigins(), realOrigin) {
	// 	err := fmt.Errorf("invalid origin of request [%v]", realOrigin)
	// 	logHandler.ErrorLogger.Println(err.Error())
	// 	watch.Stop(0)
	// 	oops(w, r, nil, "error", err.Error())
	// 	return
	// }

	locs := c.GetTranslation_PermittedLocales()
	for _, loc := range locs {
		safeLocName, _ := htmlHelpers.ToPathSafe(loc.Name)
		safeLocKey, _ := htmlHelpers.ToPathSafe(loc.Key)
		logHandler.TranslationLogger.Println("Permitted Locale: ", loc)
		localeResp.Locales = append(localeResp.Locales, struct {
			Locale string `json:"locale"`
			Name   string `json:"name"`
		}{Locale: safeLocKey, Name: safeLocName})
	}

	localeResp.Message, _ = htmlHelpers.ToPathSafe(fmt.Sprintf("Locales List, %v Locales", len(localeResp.Locales)))

	// Respond with the translated item and a success status
	w.Header().Set("Content-Type", "application/json")
	w.Header().Add("Application", appName)

	// marshal the response
	resp, err := json.Marshal(localeResp)
	if err != nil {
		logHandler.ErrorLogger.Println("Error marshalling response: ", err)
		watch.Stop(0)
		oops(w, r, nil, "error", err.Error())
		return
	}

	//	fmt.Fprintf(w, "{\"message\":\"%v\"}", translatedItem)
	//htmlResp := fmt.Sprintf("{\"message\":\"%v\"}", translatedItem)
	//w.Write([]byte(htmlResp))

	w.Write(resp)
	w.WriteHeader(http.StatusOK)
	logHandler.TranslationLogger.Println("Response to translate message [", resp, "] Status=", http.StatusOK)
	//logHandler.InfoLogger.Printf("Translated message [%v] to [%v]", itemToTranslate, translatedItem)
	watch.Stop(1)
}

func LocaleTest(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	// Build a URI query string
	baseReq := trnsl8r.NewRequest().WithProtocol(trnsServerProtocol).WithHost(trnsServerHost).WithPort(trnsServerPort).WithLogger(logHandler.ServiceLogger).FromOrigin("trnsl8r_connect")

	//baseReq.Spew()

	logHandler.TranslationLogger.Println("Request to translate message ", stringHelpers.DCurlies(baseReq.String()))

	resp, err := baseReq.GetLocales()
	if err != nil {
		logHandler.ErrorLogger.Println("Error getting locales: ", err)
		oops(w, r, nil, "error", err.Error())
		return
	}

	logHandler.TranslationLogger.Printf("Response to translate message %+v", resp)

	x, err := json.Marshal(resp)
	if err != nil {
		logHandler.ErrorLogger.Println("Error marshalling response: ", err)
		oops(w, r, nil, "error", err.Error())
		return
	}

	w.Write(x)
}
