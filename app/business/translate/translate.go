package translate

import (
	"strings"

	"github.com/mt1976/frantic-core/idHelpers"
	"github.com/mt1976/frantic-core/logHandler"
	textDataAccess "github.com/mt1976/trnsl8r_service/app/dao/textstore"
)

func Get(in, localeFilter string) string {
	// Validate the input data
	//logHandler.TranslationLogger.Printf("Translating [%v] locale=[%v]", in, localeFilter)
	id := idHelpers.Encode(strings.ToUpper(in))

	text, err := textDataAccess.GetLocalised(id, localeFilter)
	if err != nil {
		logHandler.TranslationLogger.Printf("New text translation available Id=[%v], for [%v]", id, in)
		text, err := textDataAccess.New(id, in)
		if err != nil {
			logHandler.ErrorLogger.Printf("Error creating translation! In=[%v] Working=[%v] %v", in, id, err.Error())
			logHandler.TranslationLogger.Panicf("Error creating translation! In=[%v] Working=[%v] %v", in, id, err.Error())
			return ""
		}
		logHandler.TranslationLogger.Printf("Translated [%v] to [%v]", in, text.Message)
		return text.Message
	}

	if localeFilter != "" {
		//	logHandler.TranslationLogger.Printf("Locale filter [%v] for [%v]", localeFilter, in)
		localisedText, ok := text.Localised[localeFilter]
		if !ok || localisedText == "" {
			logHandler.TranslationLogger.Printf("Locale [%v] not found for [%v]", localeFilter, in)
			//logHandler.TranslationLogger.Printf("Translated [%v] to [%v]", in, text.Message)
			return text.Message
		}
		// If the locale is found, return it, otherwise proceed to the default
		//logHandler.TranslationLogger.Printf("Translated [%v] to [%v]", in, localisedText)
		return localisedText

	}

	//logHandler.TranslationLogger.Printf("Translated [%v] to [%v]", in, text.Message)
	return text.Message
}
