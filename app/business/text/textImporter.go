package text

import (
	"context"
	"strings"

	"github.com/mt1976/frantic-core/commonConfig"
	"github.com/mt1976/frantic-core/idHelpers"
	"github.com/mt1976/frantic-core/logHandler"
	"github.com/mt1976/trnsl8r_service/app/dao/textStore"
)

var (
	settings *commonConfig.Settings
	appName  string
)

func init() {
	settings = commonConfig.Get()
	appName = settings.GetApplication_Name()
}

func Importer(ctx context.Context, in *textStore.TextStore) (err error) {
	logHandler.Trace.Printf("Importing text: %+v", in)
	signature := BuildSignature(in.Original)
	in.Signature = signature
	in.Raw = signature
	in.Key = signature
	in.Notes = strings.ToUpper(idHelpers.SanitizeID(in.Original))
	// godump.Dump(in)
	if in.SourceApplication == "" {
		in.SourceApplication = appName
	}
	if in.SourceLocale == "" {
		in.SourceLocale = settings.GetApplication_Locale()
	}
	if in.ConsumedBy == nil {
		in.ConsumedBy = addConsumer(in.ConsumedBy, appName)
	}

	if in.Localised == nil {
		in.Localised = make(map[string]string)
	}
	// Get the current locales
	locales := settings.GetTranslation_PermittedLocales()
	// Add the message to the localised map for each locale
	for _, locale := range locales {
		if _, ok := in.Localised[locale.Key]; !ok {
			in.Localised[locale.Key] = ""
		}
	}

	//	godump.Dump(in)

	return
}
