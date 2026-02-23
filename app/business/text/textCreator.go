package text

import (
	"context"

	"github.com/mt1976/frantic-core/commonConfig"
	"github.com/mt1976/frantic-core/idHelpers"
	"github.com/mt1976/trnsl8r_service/app/dao/textStore"
)

func Creator(ctx context.Context, in *textStore.TextStore) (id string, skipPostCreate bool, record *textStore.TextStore, err error) {
	// Create a new struct
	record = textStore.New()

	signature := in.Signature
	message := in.Message

	settings := commonConfig.Get()
	appName := settings.GetApplication_Name()
	// Create a new d
	record = textStore.New()
	record.Signature = signature
	record.Message = message
	record.Original = message
	record.SourceApplication = appName
	record.SourceLocale = settings.GetApplication_Locale()
	record.ConsumedBy = addConsumer(record.ConsumedBy, appName)

	if record.Localised == nil {
		record.Localised = make(map[string]string)
	}
	// Get the current locales
	locales := settings.GetTranslation_PermittedLocales()
	// Add the message to the localised map for each locale
	for _, locale := range locales {
		record.Localised[locale.Key] = ""
	}

	// Create a new struct

	record.Key = idHelpers.Encode(signature)
	id = signature
	record.Signature = signature
	record.Original = message
	record.Message = message

	skipPostCreate = false
	err = nil

	return
}

// addConsumer adds the given appName to the list of consumers if it is not already present.
// If the input list is nil, it initializes a new list with the appName.
// Parameters:
// - u: A slice of strings representing the list of consumers.
// - appName: A string representing the name of the application to be added to the list.
// Returns:
// - A slice of strings with the appName added if it was not already present.
func addConsumer(u []string, appName string) []string {
	if u == nil {
		u = []string{}
		u = append(u, appName)
		return u
	}

	inList := false

	for _, v := range u {
		if v == appName {
			// Already in the list
			inList = true
		}
	}

	if !inList {
		u = append(u, appName)
	}

	return u
}
