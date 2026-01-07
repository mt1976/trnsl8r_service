package textStore

// Data Access Object Text
// Version: 0.3.0
// Updated on: 2025-12-31

//TODO: RENAME "Text" TO THE NAME OF THE DOMAIN ENTITY
//TODO: Update the New function to implement the creation of a new domain entity
//TODO: Create any new functions required to support the domain entity

import (
	"context"
	"fmt"

	"github.com/mt1976/frantic-core/commonConfig"
	"github.com/mt1976/frantic-core/commonErrors"
	"github.com/mt1976/frantic-core/dao"

	"github.com/mt1976/frantic-core/dao/audit"
	"github.com/mt1976/frantic-core/idHelpers"
	"github.com/mt1976/frantic-core/logHandler"
	"github.com/mt1976/frantic-core/timing"
)

// New creates a new Text instance
func New() TextStore {
	return TextStore{}
}

// Create creates a new Text instance in the database
// It takes name as a parameter and returns the created Text instance or an error if any occurs
// It also checks if the DAO is ready for operations
// It records the creation action in the audit data and saves the instance to the database
// Parameters: (all but ctx are used to create a new Text instance and should be replaced as needed)
//   - ctx context.Context: The context for managing request-scoped values, cancellation signals, and deadlines.
//   - userName string: The user name for the new Text instance.
//   - uid string: The UID for the new Text instance.
//   - realName string: The real name for the new Text instance.
//   - email string: The email for the new Text instance.
//   - gid string: The GID for the new Text instance.
//
// Returns:
//   - TextStore: The created Text instance.
//   - error: An error object if any issues occur during the creation process; otherwise, nil.

func Create(ctx context.Context, signature, message string) (TextStore, error) {

	dao.CheckDAOReadyState(Domain, audit.CREATE, initialised) // Check the DAO has been initialised, Mandatory.

	//logHandler.InfoLogger.Printf("New %v (%v=%v)", domain, Fields.ID, field1)
	clock := timing.Start(Domain, "Create", fmt.Sprintf("%v", signature))

	settings := commonConfig.Get()
	appName := settings.GetApplication_Name()
	// Create a new d
	record := TextStore{}
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
	record.Raw = signature
	record.Signature = signature
	record.Original = message
	record.Message = message

	// Record the create action in the audit data
	auditErr := record.Audit.Action(ctx, audit.CREATE.WithMessage(fmt.Sprintf("New %v created %v", Domain, signature)))
	if auditErr != nil {
		// Log and panic if there is an error creating the status instance
		logHandler.ErrorLogger.Panic(commonErrors.ErrDAOUpdateAuditWrapper(Domain, record.ID, auditErr))
	}

	// Save the status instance to the database
	writeErr := activeDB.Create(&record)
	if writeErr != nil {
		// Log and panic if there is an error creating the status instance
		logHandler.ErrorLogger.Panic(commonErrors.ErrDAOCreateWrapper(Domain, record.ID, writeErr))
		//	panic(writeErr)
	}

	//logHandler.AuditLogger.Printf("[%v] [%v] ID=[%v] Notes[%v]", audit.CREATE, domain, record.ID, fmt.Sprintf("New %v: %v", domain, field1))
	clock.Stop(1)
	return record, nil
}
