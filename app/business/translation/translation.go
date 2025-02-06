package translation

import (
	"strings"

	"github.com/mt1976/frantic-plum/id"
	"github.com/mt1976/frantic-plum/logger"
	"github.com/mt1976/trnsl8r_service/app/dao/textStore"
)

func Get(in, localeFilter string) string {
	// Validate the input data
	logger.InfoLogger.Printf("Translating [%v][%v] locale=[%v]", in, strings.ToUpper(in), localeFilter)
	id := id.Encode(strings.ToUpper(in))

	text, err := textStore.Get(id, localeFilter)
	if err != nil {
		logger.TranslationLogger.Printf("New text translation available Id=[%v], for [%v]", id, in)
		text, err := textStore.New(id, in)
		if err != nil {
			logger.ErrorLogger.Printf("Error creating translation! In=[%v] Working=[%v] %v", in, id, err.Error())
			logger.TranslationLogger.Panicf("Error creating translation! In=[%v] Working=[%v] %v", in, id, err.Error())
			return ""
		}
		logger.TranslationLogger.Printf("Translated [%v] to [%v]", in, text.Message)
		return text.Message
	}

	if localeFilter != "" {
		logger.TranslationLogger.Printf("Locale filter [%v] for [%v]", localeFilter, in)
		localisedText, ok := text.Localised[localeFilter]
		if !ok || localisedText == "" {
			logger.TranslationLogger.Printf("Locale [%v] not found for [%v]", localeFilter, in)
			logger.TranslationLogger.Printf("Translated [%v] to [%v]", in, text.Message)
			return text.Message
		}
		// If the locale is found, return it, otherwise proceed to the default
		logger.TranslationLogger.Printf("Translated [%v] to [%v]", in, localisedText)
		return localisedText

	}

	logger.TranslationLogger.Printf("Translated [%v] to [%v]", in, text.Message)
	return text.Message
}
