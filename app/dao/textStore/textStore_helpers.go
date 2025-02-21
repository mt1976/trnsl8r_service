package textstore

// Data Access Object Template - templatestore
// Version: 0.2.0
// Updated on: 2021-09-10

import (
	"context"
	"fmt"
	"strings"

	"github.com/mt1976/frantic-core/commonErrors"
	"github.com/mt1976/frantic-core/dao"
	"github.com/mt1976/frantic-core/dao/actions"
	"github.com/mt1976/frantic-core/dao/audit"
	"github.com/mt1976/frantic-core/dao/database"
	"github.com/mt1976/frantic-core/logHandler"
	logger "github.com/mt1976/frantic-core/logHandler"
	"github.com/mt1976/frantic-core/timing"
)

func (record *Text_Store) prepare() (Text_Store, error) {
	//os.Exit(0)
	//logger.ErrorLogger.Printf("ACT: VAL Validate")
	dao.CheckDAOReadyState(domain, audit.PROCESS, initialised) // Check the DAO has been initialised, Mandatory.

	return *record, nil
}

func (record *Text_Store) calculate() error {

	dao.CheckDAOReadyState(domain, audit.PROCESS, initialised) // Check the DAO has been initialised, Mandatory.

	// Calculate the duration in days between the start and end dates
	return nil
}

func (record *Text_Store) isDuplicateOf(id string) (Text_Store, error) {

	dao.CheckDAOReadyState(domain, audit.PROCESS, initialised) // Check the DAO has been initialised, Mandatory.

	//TODO: Could be replaced with a simple read...

	// Get all status
	recordList, err := GetAll()
	if err != nil {
		logHandler.ErrorLogger.Printf("Getting all %v failed %v", domain, err.Error())
		return Text_Store{}, err
	}

	// range through status list, if status code is found and deletedby is empty then return error
	for _, checkRecord := range recordList {
		//s.Dump(!,strings.ToUpper(code) + "-uchk-" + s.Code)
		testValue := strings.ToUpper(id)
		checkValue := strings.ToUpper(checkRecord.Signature)
		//logger.InfoLogger.Printf("CHK: TestValue:[%v] CheckValue:[%v]", testValue, checkValue)
		//logger.InfoLogger.Printf("CHK: Code:[%v] s.Code:[%v] s.Audit.DeletedBy:[%v]", testCode, s.Code, s.Audit.DeletedBy)
		if checkValue == testValue && checkRecord.Audit.DeletedBy == "" {
			logHandler.WarningLogger.Printf("Duplicate %v, %v already in use", strings.ToUpper(domain), record.ID)
			return checkRecord, commonErrors.ErrorDuplicate
		}
	}

	return Text_Store{}, nil
}

func ClearDown(ctx context.Context) error {
	logHandler.InfoLogger.Printf("Clearing %v", domain)

	dao.CheckDAOReadyState(domain, audit.PROCESS, initialised) // Check the DAO has been initialised, Mandatory.

	clock := timing.Start(domain, actions.CLEAR.GetCode(), "INITIALISE")

	// Delete all active session recordList
	recordList, err := GetAll()
	if err != nil {
		logHandler.ErrorLogger.Print(commonErrors.WrapDAOInitialisationError(domain, err).Error())
		clock.Stop(0)
		return commonErrors.WrapDAOInitialisationError(domain, err)
	}

	noRecords := len(recordList)
	count := 0

	for thisRecord, record := range recordList {
		logHandler.InfoLogger.Printf("Deleting %v (%v/%v) %v", domain, thisRecord, noRecords, record.Signature)
		delErr := Delete(ctx, record.ID, fmt.Sprintf("Clearing %v %v @ initialisation ", domain, record.ID))
		if delErr != nil {
			logHandler.ErrorLogger.Print(commonErrors.WrapDAOInitialisationError(domain, delErr).Error())
			continue
		}
		count++
	}

	clock.Stop(count)

	return nil
}

func (u *Text_Store) dup(name string) (Text_Store, error) {

	//logger.InfoLogger.Printf("CHK: CheckUniqueCode %v", name)

	// Get all status
	statusList, err := GetAll()
	if err != nil {
		logger.ErrorLogger.Printf("Error Getting all status: %v", err.Error())
		return Text_Store{}, err
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
			return s, commonErrors.ErrorDuplicate
		}
	}

	//logger.InfoLogger.Printf("CHK: %v is unique", strings.ToUpper(name))

	// Return nil if the code is unique

	return Text_Store{}, nil
}

func GetBySignature(key any) (Text_Store, error) {
	return GetBy(FIELD_Signature, key)
}

func DeleteBySignature(ctx context.Context, key string, note string) error {
	return DeleteBy(ctx, FIELD_Signature, key, note)
}

func GetLocalised(signature, localeFilter string) (Text_Store, error) {

	watch := timing.Start(domain, actions.GET.GetCode(), signature)
	// Log the start of the retrieval operation
	//logger.InfoLogger.Printf("GET: [%v] Id=[%v]", strings.ToUpper(tableName), id)

	// Log the ID of the dest object being retrieved
	//logger.InfoLogger.Printf("GET: %v Object: %v", tableName, fmt.Sprintf("%+v", id))

	// Initialize an empty d object
	u := Text_Store{}

	// Retrieve the dest object from the database based on the given IDs
	u, err := GetBy(FIELD_Signature, signature)
	if err != nil {
		// Log and panic if there is an error reading the dest object
		//logger.InfoLogger.Printf("Reading %v: [%v] %v ", tableName, id, err.Error())
		return Text_Store{}, fmt.Errorf("reading %v Id=[%v] %v ", strings.ToUpper(domain), signature, err.Error())
		//	panic(err)
	}

	// Log the retrieved dest object
	//u.Spew()

	// Log the completion of the retrieval operation
	//logger.InfoLogger.Printf("GET: [%v] Id=[%v] RealName=[%v] ", strings.ToUpper(tableName), u.ID, u.RealName)

	err = u.PostGet()
	if err != nil {
		return Text_Store{}, err
	}
	watch.Stop(1)

	return u, nil
}

func GetDB() func() (*database.DB, error) {
	logHandler.InfoLogger.Println("GETDB")
	return func() (*database.DB, error) {
		logHandler.InfoLogger.Printf("GETDB2")
		return database.Connect(), nil
	}
}
