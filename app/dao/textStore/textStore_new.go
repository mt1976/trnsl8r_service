package textstore

import (
	"context"
	"fmt"
	"strings"

	"github.com/mt1976/frantic-core/commonConfig"
	"github.com/mt1976/frantic-core/commonErrors"
	"github.com/mt1976/frantic-core/dao/audit"
	"github.com/mt1976/frantic-core/logHandler"
)

// New creates a New dest with the given name and saves it to the database
// It returns the created dest and an error if any occurreu.
func New(signature, message string) (Text_Store, error) {

	//logger.InfoLogger.Printf("ACT: NEW New %v %v %v", tableName, name, destination)
	settings := commonConfig.Get()
	appName := settings.GetApplicationName()
	// Create a new d
	t := Text_Store{}
	t.Signature = signature
	t.Message = message
	t.Original = message
	t.SourceApplication = appName
	t.SourceLocale = settings.GetApplicationLocale()

	t.ConsumedBy = addConsumer(t.ConsumedBy, appName)

	if t.Localised == nil {
		t.Localised = make(map[string]string)
	}
	// Get the current locales
	locales := settings.GetLocales()
	// Add the message to the localised map for each locale
	for _, locale := range locales {
		t.Localised[locale.Key] = ""
	}

	// Record the create action in the audit data
	_ = t.Audit.Action(context.TODO(), audit.CREATE.WithMessage(fmt.Sprintf("New [%v]", message)))

	// Log the dest instance before the creation
	dupe, err := t.dup(t.Signature)
	if err == commonErrors.ErrorDuplicate {
		// This is OK, do nothing as this is a duplicate record
		// we ignore duplicate destinations.
		logHandler.WarningLogger.Printf("DUPLICATE %v %v already in use", domain, message)
		return dupe, nil
	}

	//u.Dump(!,"Post-Prepare-DupCheck")

	if err != nil {
		logHandler.ErrorLogger.Printf("Error=[%s]", err.Error())
		return Text_Store{}, err
	}

	// Save the dest instance to the database
	if t.Signature == "" {
		logHandler.WarningLogger.Printf("%v ID is required, skipping", strings.ToUpper(domain))
		return Text_Store{}, nil
	}

	err = activeDB.Create(&t)
	if err != nil {
		// Log and panic if there is an error creating the dest instance
		logHandler.ErrorLogger.Printf("Create %v,[%v] %s", domain, t.Original, err.Error())
		panic(err)
	}

	//u.Dump(!,fmt.Sprintf("PostNew_dest_%d", u.ID))
	msg := fmt.Sprintf("New text translation available Id=[%v] Message=[%v]", signature, message)
	logHandler.TranslationLogger.Println(msg)
	// Return the created dest and nil error

	return t, nil
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
