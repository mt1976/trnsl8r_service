package textStore

// Data Access Object Text
// Version: 0.3.0
// Updated on: 2025-12-31

import (
	"context"
	"fmt"
	"strings"

	"github.com/mt1976/frantic-core/commonErrors"
	ce "github.com/mt1976/frantic-core/commonErrors"
	"github.com/mt1976/frantic-core/dao/database"
	"github.com/mt1976/frantic-core/logHandler"
	"github.com/mt1976/frantic-core/timing"
)

//TODO: RENAME "Text" TO THE NAME OF THE DOMAIN ENTITY
//TODO: Implement the validate function to process the domain entity
//TODO: Implement the calculate function to process the domain entity
//TODO: Implement the isDuplicateOf function to process the domain entity
//TODO: Implement the postGetProcessing function to process the domain entity

// upgradeProcessing performs any necessary upgrade processing for the TextStore record.
// This processing is triggered directly after the record has been retrieved from the database.
func (record *TextStore) upgradeProcessing() error {
	//TODO: Add any upgrade processing here
	//This processing is triggered directly after the record has been retrieved from the database
	return nil
}

// defaultProcessing performs default calculations for the TextStore record.
// This processing is triggered directly before the record is saved to the database.
func (record *TextStore) defaultProcessing() error {
	//TODO: Add any default calculations here
	//This processing is triggered directly before the record is saved to the database
	return nil
}

// validationProcessing performs validation checks for the TextStore record.
// This processing is triggered directly before the record is saved to the database and after the default calculations.
func (record *TextStore) validationProcessing() error {
	//TODO: Add any record validation here
	//This processing is triggered directly before the record is saved to the database and after the default calculations
	return nil
}

// postGetProcessing performs any necessary processing for the TextStore record after it has been retrieved from the database.
// This processing is triggered directly after the record has been retrieved from the database and after the upgrade processing.
func (h *TextStore) postGetProcessing() error {
	//TODO: Add any post get processing here
	//This processing is triggered directly after the record has been retrieved from the database and after the upgrade processing
	return nil
}

// preDeleteProcessing performs any necessary processing for the TextStore record before it is deleted from the database.
// This processing is triggered directly before the record is deleted from the database.
func (record *TextStore) preDeleteProcessing() error {
	//TODO: Add any pre delete processing here
	//This processing is triggered directly before the record is deleted from the database
	return nil
}

// TextClone creates a clone of the given TextStore record.
// This function is used to duplicate a TextStore record.
func TextClone(ctx context.Context, source TextStore) (TextStore, error) {
	//TODO: Add any clone processing here
	panic("Not Implemented")
}

// assertTextStore asserts that the given result is of type *TextStore.
// It returns the asserted TextStore and any error encountered during the assertion.
func assertTextStore(result any, field database.Field, value any) (*TextStore, error) {
	x, ok := result.(*TextStore)
	if !ok {
		return nil, ce.ErrDAOAssertWrapper(Domain, field.String(), value,
			ce.ErrInvalidTypeWrapper(field.String(), fmt.Sprintf("%T", result), "*TextStore"))
	}
	return x, nil
}

// // prepare performs preparation steps for the TextStore record before it is saved to the database.
// func (u *TextStore) prepare() (TextStore, error) {
// 	//os.Exit(0)
// 	// logHandler.ErrorLogger.Printf("ACT: VAL Validate")
// 	user, err := u.dup()
// 	if err == commonErrors.ErrorDuplicate {
// 		return *u, nil
// 	}
// 	if err != nil {
// 		return user, err
// 	}
// 	return *u, nil
// }

// // calculate performs calculations for the TextStore record before it is saved to the database.
// func (u *TextStore) calculate() error {
// 	// Calculate the duration in days between the start and end dates
// 	return nil
// }

// // dup checks for duplicate TextStore records based on UserCode and UserName.
// func (u *TextStore) dup() (TextStore, error) {

// 	// logHandler.InfoLogger.Printf("CHK: CheckUniqueCode %v", name)

// 	// Get all status
// 	userList, err := GetAll()
// 	if err != nil {
// 		logHandler.ErrorLogger.Printf("Error Getting all Users: %v", err.Error())
// 		return TextStore{}, err
// 	}

// 	// range through status list, if status code is found and deletedby is empty then return error
// 	for _, s := range userList {
// 		//s.Dump(strings.ToUpper(code) + "-uchk-" + s.Code)
// 		testValue := strings.ToUpper(u.UserCode)
// 		checkValue := strings.ToUpper(s.UserCode)
// 		// logHandler.InfoLogger.Printf("CHK: TestValue:[%v] CheckValue:[%v]", testValue, checkValue)
// 		// logHandler.InfoLogger.Printf("CHK: Code:[%v] s.Code:[%v] s.Audit.DeletedBy:[%v]", testCode, s.Code, s.Audit.DeletedBy)
// 		if checkValue == testValue && s.Audit.DeletedBy == "" {
// 			logHandler.WarningLogger.Printf("[%v] DUPLICATE UID [%v] already in use for [%v]", TableName, testValue, s.UserName)
// 			return s, commonErrors.ErrorDuplicate
// 		}
// 		testValue = strings.ToUpper(u.UserName)
// 		checkValue = strings.ToUpper(s.UserName)
// 		// logHandler.InfoLogger.Printf("CHK: TestValue:[%v] CheckValue:[%v]", testValue, checkValue)
// 		// logHandler.InfoLogger.Printf("CHK: Code:[%v] s.Code:[%v] s.Audit.DeletedBy:[%v]", testCode, s.Code, s.Audit.DeletedBy)
// 		if checkValue == testValue && s.Audit.DeletedBy == "" {
// 			logHandler.WarningLogger.Printf("[%v] DUPLICATE User Name [%v] already in use for [%v]", TableName, testValue, s.UserName)
// 			return s, commonErrors.ErrorDuplicate
// 		}
// 	}

// 	// Return nil if the code is unique

// 	return *u, nil
// }

// End of Model Helpers

// Insert additional functions below this line

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

func (u *TextStore) dup(name string) (TextStore, error) {

	//logger.InfoLogger.Printf("CHK: CheckUniqueCode %v", name)

	// Get all status
	statusList, err := GetAll()
	if err != nil {
		logHandler.ErrorLogger.Printf("Error Getting all status: %v", err.Error())
		return TextStore{}, err
	}

	// range through status list, if status code is found and deletedby is empty then return error
	for _, s := range statusList {
		//s.Dump(!,strings.ToUpper(code) + "-uchk-" + s.Code)
		testValue := strings.ToUpper(name)
		checkValue := strings.ToUpper(s.Signature)
		//logger.InfoLogger.Printf("CHK: TestValue:[%v] CheckValue:[%v]", testValue, checkValue)
		//logger.InfoLogger.Printf("CHK: Code:[%v] s.Code:[%v] s.Audit.DeletedBy:[%v]", testCode, s.Code, s.Audit.DeletedBy)
		if checkValue == testValue && s.Audit.DeletedBy == "" {
			//logger.InfoLogger.Printf("[%v] DUPLICATE %v already in use", strings.ToUpper(tableName), name)
			return s, commonErrors.ErrDuplicate
		}
	}

	//logger.InfoLogger.Printf("CHK: %v is unique", strings.ToUpper(name))

	// Return nil if the code is unique

	return TextStore{}, nil
}

func GetBySignature(key any) (TextStore, error) {
	return GetBy(Fields.Signature, key)
}

func DeleteBySignature(ctx context.Context, key string, note string) error {
	return DeleteBy(ctx, Fields.Signature, key, note)
}

func GetLocalised(signature, localeFilter string) (TextStore, error) {

	watch := timing.Start(Domain, "Get", signature)
	// Log the start of the retrieval operation
	//logger.InfoLogger.Printf("GET: [%v] Id=[%v]", strings.ToUpper(tableName), id)

	// Log the ID of the dest object being retrieved
	//logger.InfoLogger.Printf("GET: %v Object: %v", tableName, fmt.Sprintf("%+v", id))

	// Initialize an empty d object
	u := TextStore{}

	// Retrieve the dest object from the database based on the given IDs
	u, err := GetBy(Fields.Signature, signature)
	if err != nil {
		// Log and panic if there is an error reading the dest object
		//logger.InfoLogger.Printf("Reading %v: [%v] %v ", tableName, id, err.Error())
		return TextStore{}, fmt.Errorf("reading %v Id=[%v] %v ", Domain, signature, err.Error())
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
	watch.Stop(1)

	return u, nil
}

func FetchDatabaseInstances() func() ([]*database.DB, error) {
	return func() ([]*database.DB, error) {
		return []*database.DB{activeDB}, nil
	}
}
