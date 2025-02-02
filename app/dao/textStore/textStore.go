package textStore

import (
	"context"
	"fmt"
	"strings"

	"github.com/mt1976/frantic-plum/common"
	"github.com/mt1976/frantic-plum/commonErrors"
	"github.com/mt1976/frantic-plum/dao/audit"
	"github.com/mt1976/frantic-plum/dao/database"
	"github.com/mt1976/frantic-plum/io"
	"github.com/mt1976/frantic-plum/logger"
	"github.com/mt1976/frantic-plum/paths"
	"github.com/mt1976/frantic-plum/timing"
	stopwatch "github.com/mt1976/frantic-plum/timing"
)

// New creates a New dest with the given name and saves it to the database
// It returns the created dest and an error if any occurreu.
func New(signature, message string) (TextStore, error) {

	//logger.InfoLogger.Printf("ACT: NEW New %v %v %v", tableName, name, destination)
	settings := common.Get()
	// Create a new d
	t := TextStore{}
	t.Signature = signature
	t.Message = message
	t.Original = message
	t.Source = settings.ApplicationName()
	t.Locale = settings.ApplicationLocale()

	t.ConsumedBy = addConsumer(t.ConsumedBy, settings.ApplicationName())

	if t.Localised == nil {
		t.Localised = make(map[string]string)
	}
	t.Localised["en_GB"] = "reserved"
	t.Localised["eu_ES"] = "reserved"

	// Add basic attributes

	// Record the create action in the audit data
	_ = t.Audit.Action(nil, audit.CREATE.WithMessage(fmt.Sprintf("New text [%v]", message)))

	// Log the dest instance before the creation
	xtext, err := t.prepare()
	if err == commonErrors.ErrorDuplicate {
		// This is OK, do nothing as this is a duplicate record
		// we ignore duplicate destinations.
		logger.WarningLogger.Printf("[%v] DUPLICATE %v already in use", strings.ToUpper(tableName), message)
		return xtext, nil
	}

	//u.Dump(!,"Post-Prepare-DupCheck")

	if err != nil {
		logger.ErrorLogger.Printf("Error=[%s]", err.Error())
		return TextStore{}, err
	}

	// Save the dest instance to the database
	if t.Signature == "" {
		logger.WarningLogger.Printf("[%v] ID is required, skipping", strings.ToUpper(tableName))
		return TextStore{}, nil
	}

	err = database.Create(&t)
	if err != nil {
		// Log and panic if there is an error creating the dest instance
		logger.ErrorLogger.Printf("[%v] Create [%v] %s", strings.ToUpper(tableName), t.Original, err.Error())
		panic(err)
	}

	//u.Dump(!,fmt.Sprintf("PostNew_dest_%d", u.ID))
	msg := fmt.Sprintf("New text translation available Id=[%v] Message=[%v]", signature, message)
	logger.TranslationLogger.Printf(msg)
	// Return the created dest and nil error
	logger.AuditLogger.Printf("[%v] [%v] ID=[%v] Notes[%v]", strings.ToUpper(tableName), audit.CREATE.Code(), signature, msg)

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

// update updates the current dest instance in the database
// It logs the update operation and records the audit data.
//
// Parameters:
// - dest: A pointer to the dest instance to be updateu.
//
// Returns:
// - error: An error if any occurred during the update operation.
func (u *TextStore) Update(ctx context.Context, note string) error {
	//logger.InfoLogger.Printf("ACT: UPD Update")

	err := u.validate()
	if err != nil {
		return err
	}

	// Run record level validation/business processing
	calculationError := u.calculate()
	if calculationError != nil {
		logger.ErrorLogger.Printf("[%v] Calculating %e", strings.ToUpper(tableName), calculationError)
		return calculationError
	}

	// Run record level validation/business processing
	_, validationError := u.prepare()
	if validationError != nil {
		logger.ErrorLogger.Printf("[%v] Validating %v", strings.ToUpper(tableName), validationError.Error())
		return validationError
	}

	// Record the update action in the audit data
	_ = u.Audit.Action(ctx, audit.UPDATE.WithMessage(note))

	// Log the dest instance before the update
	//u.Spew()

	//u.Dump(!,fmt.Sprintf("PreUpdate_dest_%d", u.ID))

	// Update the dest instance in the database
	err = database.Update(u)
	if err != nil {
		// Log and panic if there is an error updating the dest instance
		logger.ErrorLogger.Printf("[%v] Updating %e", strings.ToUpper(tableName), err)
		panic(err)
	}

	// Log the completion of the update operation
	//logger.InfoLogger.Printf("Update %v: [%v][%v] ", tableName, u.ID, u.Name)

	//	logger.InfoLogger.Printf("Update [%v] ID=[%v] Message=[%v]", strings.ToUpper(tableName), u.ID, u.Message)
	logger.AuditLogger.Printf("[%v] [%v] ID=[%v] Notes[%v]", audit.UPDATE, strings.ToUpper(tableName), u.Signature, note)

	// Return nil if the update operation is successful
	return nil
}

// Get retrieves a dest object from the database based on the given Iu.
// It returns the retrieved dest object and an error if any occurreu.
//
// Parameters:
// - id: The unique identifier of the dest object to retrieve.
//
// Returns:
// - dest: The retrieved dest object.
// - error: An error if any occurred during the retrieval operation.
func get(signature string) (TextStore, error) {

	get := stopwatch.Start(tableName, "get", signature)
	// Log the start of the retrieval operation
	//logger.InfoLogger.Printf("GET: [%v] Id=[%v]", strings.ToUpper(tableName), id)

	// Log the ID of the dest object being retrieved
	//logger.InfoLogger.Printf("GET: %v Object: %v", tableName, fmt.Sprintf("%+v", id))

	// Initialize an empty d object
	u := TextStore{}

	// Retrieve the dest object from the database based on the given ID
	err := database.Retrieve(Field_Signature, signature, &u)
	if err != nil {
		// Log and panic if there is an error reading the dest object
		//logger.InfoLogger.Printf("Reading %v: [%v] %v ", tableName, id, err.Error())
		return TextStore{}, fmt.Errorf("[%v] Error Reading Id=[%v] %v ", strings.ToUpper(tableName), signature, err.Error())
		//	panic(err)
	}

	// Log the retrieved dest object
	//u.Spew()

	// Log the completion of the retrieval operation
	//logger.InfoLogger.Printf("GET: [%v] Id=[%v] RealName=[%v] ", strings.ToUpper(tableName), u.ID, u.RealName)

	err = u.postGet()
	if err != nil {
		return TextStore{}, err
	}
	get.Stop(1)
	//u.Dump(!,"PostGet")
	// Return the retrieved dest object and nil error
	return u, nil
}

// GetAll retrieves all dest objects from the database
// It returns a slice of dest objects and an error if any occurreu.
//
// Parameters:
//
//	None
//
// Returns:
//
//	[]dest: A slice of dest objects.
//	error: An error if any occurred during the retrieval operation.
func GetAll() ([]TextStore, error) {

	uList := []TextStore{}

	//logger.InfoLogger.Printf("GTA: [%v] Get All", strings.ToUpper(tableName))

	gall := stopwatch.Start(tableName, "Get All", "Get")
	errG := database.GetAll(&uList)
	if errG != nil {
		logger.ErrorLogger.Printf("[%v] Reading Id=[%v] %v ", strings.ToUpper(tableName), "ALL", errG.Error())
		panic(errG)
	}
	gall.Stop(len(uList))

	//logger.InfoLogger.Printf("GTA: [%v] Count=[%v]", strings.ToUpper(tableName), len(dList))

	dList, errPost := postGet(&uList)
	if errPost != nil {
		return nil, errPost
	}

	return dList, nil
}

// delete deletes the current dest instance from the database
// It logs the deletion operation and records the audit data.
//
// Parameters:
// - dest: A pointer to the dest instance to be deleteu.
//
// Returns:
// - error: An error if any occurred during the deletion operation.
func (u *TextStore) delete(ctx context.Context, note string) error {

	// Log the start of the deletion operation
	logger.InfoLogger.Printf("DEL: [%v] Id=[%v] Message=[%v]", strings.ToUpper(tableName), u.Signature, u.Message)

	// Record the delete action in the audit data
	_ = u.Audit.Action(ctx, audit.DELETE.WithMessage(note))

	// Log the dest instance before the deletion
	//u.Spew()
	u.Dump("DEL")

	// Delete the dest instance from the database
	err := database.Drop(u)
	if err != nil {
		// Log and panic if there is an error deleting the dest instance
		logger.ErrorLogger.Printf("[%v] Deleting %e ", strings.ToUpper(tableName), err)
		panic(err)
	}

	// Log the completion of the deletion operation
	logger.AuditLogger.Printf("DEL: [%v] ID=[%04v] RealName=[%v] ", strings.ToUpper(tableName), u.Signature, u.Message)

	// Return nil if the deletion operation is successful
	return nil
}

// deleteBySignature deletes a dest instance from the database based on the given Iu.
// It logs the deletion operation and records the audit data.
//
// Parameters:
// - id: The unique identifier of the dest instance to delete.
//
// Returns:
// - error: An error if any occurred during the deletion operation.
func deleteBySignature(ctx context.Context, signature string, note string) error {
	// Log the start of the deletion operation
	logger.InfoLogger.Printf("DLI: [%v] Id=[%v]", tableName, signature)

	// Log the dest instance before the deletion
	dest, err := get(signature)
	if err != nil {
		// Log and panic if there is an error reading the dest instance
		logger.ErrorLogger.Printf("[%v] Reading Id=[%v] %v", strings.ToUpper(tableName), signature, err.Error())
		panic(err)
	}

	// Record the delete action in the audit data
	_ = dest.Audit.Action(ctx, audit.DELETE.WithMessage(note))
	dest.Dump("DEL")

	// Delete the dest instance from the database
	err = database.Drop(dest)
	if err != nil {
		// Log and panic if there is an error deleting the dest instance
		logger.ErrorLogger.Printf("[%v] Deleting  %e ", strings.ToUpper(tableName), err)
		panic(err)
	}

	// Log the completion of the deletion operation
	logger.AuditLogger.Printf("DLI: [%v] Id=[%04v] ", strings.ToUpper(tableName), signature)

	// Return nil if the deletion operation is successful
	return nil
}

// Spew prints the dest instance details along with the audit data.
// It logs the dest instance details and the number of updates.
// If there are updates, it logs each update action along with the timestamp,
// the text who made the update, and the date of the update.
//
// Parameters:
// - dest: A pointer to the dest instance to be printeu.
//
// Returns:
// - None
func (u *TextStore) Spew() {
	logger.InfoLogger.Printf(" [%v] ID=[%v] Message=[%v] Original=[%v]", strings.ToUpper(tableName), u.Signature, u.Message, u.Original)
}

func (u *TextStore) validate() error {
	//logger.InfoLogger.Printf("ACT: VAL Validate")

	//make sure the DIsplay ID is populated

	return nil
}

func postGet(textList *[]TextStore) ([]TextStore, error) {
	//	logger.InfoLogger.Printf("ACT: PGT PostGet List")

	newList := []TextStore{}

	for _, text := range *textList {
		err := text.postGet()
		if err != nil {
			return nil, err
		}
		newList = append(newList, text)
		//	u.Spew()
	}

	return newList, nil
}

func (u *TextStore) postGet() error {
	//	logger.InfoLogger.Printf("ACT: PGT PostGet")

	//u.Dump(!,fmt.Sprintf("PostGet_dest_%d", u.ID))

	return nil
}

func (u *TextStore) Dump(name string) {
	//output := fmt.Sprintf("%v+", yy)
	io.Dump(tableName, paths.Dumps(), name, u.Signature, u)
}

func DumpAll(ctx context.Context) {
	dList, _ := GetAll()
	if len(dList) == 0 {
		logger.EventLogger.Printf("Backup [%v] no data found", strings.ToUpper(tableName))
		return
	}

	for _, yy := range dList {
		yy.Dump("EXPORT")
	}
}

func GetBySignature(signature string) (TextStore, error) {

	get := timing.Start(tableName, "Get", signature)
	// Log the start of the retrieval operation
	//logger.InfoLogger.Printf("GET: [%v] Id=[%v]", strings.ToUpper(tableName), id)

	// Log the ID of the status object being retrieved
	//logger.InfoLogger.Printf("GET: %v Object: %v", tableName, fmt.Sprintf("%+v", id))

	// Initialize an empty txt object
	txt := TextStore{}

	// Retrieve the status object from the database based on the given ID
	err := database.Retrieve(Field_Signature, signature, &txt)
	if err != nil {
		// Log and panic if there is an error reading the status object
		msg := fmt.Sprintf("[%v] Reading Id=[%v] %v", strings.ToUpper(tableName), signature, err.Error())
		logger.WarningLogger.Println(msg)
		return TextStore{}, fmt.Errorf(msg, "")
	}

	// Log the retrieved status object
	//status.Spew()

	// Log the completion of the retrieval operation
	//logger.InfoLogger.Printf("GET: [%v] Id=[%v] Name=[%v] ", strings.ToUpper(tableName), status.ID, status.Name)

	err = txt.postGet()
	if err != nil {
		return TextStore{}, err
	}

	get.Stop(1)
	// Return the retrieved status object and nil error
	return txt, nil
}

func Get(signature string) (TextStore, error) {

	watch := stopwatch.Start(tableName, "Get", signature)
	// Log the start of the retrieval operation
	//logger.InfoLogger.Printf("GET: [%v] Id=[%v]", strings.ToUpper(tableName), id)

	// Log the ID of the dest object being retrieved
	//logger.InfoLogger.Printf("GET: %v Object: %v", tableName, fmt.Sprintf("%+v", id))

	// Initialize an empty d object
	u := TextStore{}

	// Retrieve the dest object from the database based on the given ID
	err := database.Retrieve(Field_Signature, signature, &u)
	if err != nil {
		// Log and panic if there is an error reading the dest object
		//logger.InfoLogger.Printf("Reading %v: [%v] %v ", tableName, id, err.Error())
		return TextStore{}, fmt.Errorf("[%v] Reading Id=[%v] %v ", strings.ToUpper(tableName), signature, err.Error())
		//	panic(err)
	}

	// Log the retrieved dest object
	//u.Spew()

	// Log the completion of the retrieval operation
	//logger.InfoLogger.Printf("GET: [%v] Id=[%v] RealName=[%v] ", strings.ToUpper(tableName), u.ID, u.RealName)

	err = u.PostGet()
	if err != nil {
		return TextStore{}, err
	}
	watch.Stop(1)
	//u.Dump(!,"PostGet")
	// Return the retrieved dest object and nil error
	return u, nil
}

func Drop() error {
	return database.Drop(tableName)
}
