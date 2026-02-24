package translate

import (
	"context"

	"github.com/mt1976/frantic-core/logHandler"
	"github.com/mt1976/trnsl8r_service/app/business/text"
	textDataAccess "github.com/mt1976/trnsl8r_service/app/business/text"
	"github.com/mt1976/trnsl8r_service/app/dao/textStore"
)

func Get(ctx context.Context, in, localeFilter string) string {
	// Validate the input data
	// logHandler.Translation.Printf("Translating [%v] locale=[%v]", in, localeFilter)
	signature := text.BuildSignature(in)

	inRec := textStore.New()
	inRec.Signature = signature
	inRec.Message = in

	text, err := textDataAccess.GetLocalised(signature, localeFilter)
	if err != nil {
		logHandler.Translation.Printf("New text translation available Id=[%v], for [%v]", signature, in)
		text, err := textStore.Create(ctx, inRec)
		if err != nil {
			logHandler.Error.Printf("Error creating translation! In=[%v] Working=[%v] %v", in, signature, err.Error())
			logHandler.Translation.Panicf("Error creating translation! In=[%v] Working=[%v] %v", in, signature, err.Error())
			return ""
		}
		logHandler.Translation.Printf("Translated [%v] to [%v]", in, text.Message)
		return text.Message
	}

	if localeFilter != "" && localeFilter != "en" && localeFilter != "en_GB" && localeFilter != "en_US" {
		//	logHandler.Translation.Printf("Locale filter [%v] for [%v]", localeFilter, in)
		localisedText, ok := text.Localised[localeFilter]
		if !ok || localisedText == "" {
			logHandler.Translation.Printf("Locale [%v] not found for [%v]", localeFilter, in)
			// logHandler.Translation.Printf("Translated [%v] to [%v]", in, text.Message)
			return text.Message
		}
		// If the locale is found, return it, otherwise proceed to the default
		// logHandler.Translation.Printf("Translated [%v] to [%v]", in, localisedText)
		return localisedText

	}

	// logHandler.Translation.Printf("Translated [%v] to [%v]", in, text.Message)
	return text.Message
}
