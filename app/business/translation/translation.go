package translation

import (
	"github.com/mt1976/trnsl8r_service/app/dao/textStore"
	"github.com/mt1976/trnsl8r_service/app/support/id"
	"github.com/mt1976/trnsl8r_service/app/support/logger"
)

func Get(in string) string {
	// Validate the input data

	id := id.Encode(in)

	text, err := textStore.Get(id)
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
	logger.TranslationLogger.Printf("Translated [%v] to [%v]", in, text.Message)
	return text.Message
}
