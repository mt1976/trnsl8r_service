package textStore

import (
	"strings"

	"github.com/mt1976/frantic-plum/errors"
	"github.com/mt1976/frantic-plum/logger"
)

func (u *TextStore) prepare() (TextStore, error) {
	//os.Exit(0)
	//logger.ErrorLogger.Printf("ACT: VAL Validate")
	// text, err := u.dup(u.ID)
	// if err != nil {
	// 	return text, err
	// }
	return *u, nil
}

func (u *TextStore) calculate() error {
	// Calculate the duration in days between the start and end dates
	return nil
}

func (u *TextStore) dup(name string) (TextStore, error) {

	//logger.InfoLogger.Printf("CHK: CheckUniqueCode %v", name)

	// Get all status
	statusList, err := GetAll()
	if err != nil {
		logger.ErrorLogger.Printf("Error Getting all status: %v", err.Error())
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
			return s, errors.ErrorDuplicate
		}
	}

	//logger.InfoLogger.Printf("CHK: %v is unique", strings.ToUpper(name))

	// Return nil if the code is unique

	return TextStore{}, nil
}

func (u *TextStore) PostGet() error {
	//	logger.InfoLogger.Printf("ACT: PGT PostGet")

	//u.Dump(!,fmt.Sprintf("PostGet_dest_%d", u.ID))

	return nil
}
