package support

import (
	"errors"
)

var (
	ErrorEndDateBeforeStartDate = errors.New("end date is before start date")
	ErrorEmptyName              = errors.New("name is empty")
	ErrorNameTooLong            = errors.New("name is too long, max 50 characters")
	ErrorDuplicate              = errors.New("duplicate")
	ErrorNegativeValue          = errors.New("negative value")
	ErrorNotFound               = errors.New("not found %v %v")
	ErrorPasswordMismatch       = errors.New("password mismatch")
)

func HandleGoValidatorError(err error) error {
	return nil
	// if err != nil {

	// 	if _, ok := err.(*validator.InvalidValidationError); ok {
	// 		logger.InfoLogger.Println(err)
	// 		return err
	// 	}

	// 	for _, err := range err.(validator.ValidationErrors) {

	// 		op := fmt.Sprintf("VALIDATION: Field[%s] Tag[%s] Kind[%s] Param[%s] Value[%s]", err.Field(), err.Tag(), err.Kind(), err.Param(), err.Value())
	// 		logger.InfoLogger.Println(op)

	// 	}

	// 	return err
	// }
	// return nil
}
