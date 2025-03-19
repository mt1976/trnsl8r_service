package commonErrors

import (
	"errors"
	"fmt"

	"github.com/mt1976/frantic-core/logHandler"
)

var (
	ErrorEndDateBeforeStartDate = errors.New("end date is before start date")
	ErrorEmptyName              = errors.New("name is empty")
	ErrorNameTooLong            = errors.New("name is too long, max 50 characters") // Deprecated: use StringTooLongError
	ErrorDuplicate              = errors.New("duplicate")
	ErrorNegativeValue          = errors.New("negative value")
	//ErrorNotFound               = errors.New("not found %w %w") // Deprecated: use NotFoundError
	ErrorPasswordMismatch       = errors.New("password mismatch")
	ErrorUserNotFound           = errors.New("user not found")
	ErrorUserNotActive          = errors.New("user not active")
	ErrNoTranslation            = errors.New("no translation available")
	ErrNoMessageToTranslate     = errors.New("no message to translate")
	ErrProtocolIsRequired       = errors.New("protocol is required")
	ErrInvalidProtocol          = errors.New("invalid protocol")
	ErrHostIsRequired           = errors.New("host is required")
	ErrInvalidHost              = errors.New("invalid host")
	ErrPortIsRequired           = errors.New("port is required")
	ErrInvalidPort              = errors.New("invalid port")
	ErrUsernameIsRequired       = errors.New("username is required")
	ErrInvalidUsername          = errors.New("invalid username")
	ErrPasswordIsRequired       = errors.New("password is required")
	ErrInvalidPassword          = errors.New("invalid password")
	ErrOriginIsRequired         = errors.New("no origin defined, and origin identifier is required")
	ErrInvalidOrigin            = errors.New("invalid origin")
	ErrContextCannotGetUserCode = errors.New("cannot get user from context")
)

func WrapStringTooLongErr(err error, ln int) error {
	return fmt.Errorf("string too long, max %d characters error (%w)", ln, err)
}

func WrapNotFoundError(table string, err error) error {
	return fmt.Errorf("%v not found (%w)", table, err)
}
func WrapReadError(err error) error {
	return fmt.Errorf("read error (%w)", err)
}
func WrapWriteError(err error) error {
	return fmt.Errorf("write error (%w)", err)
}
func WrapEmptyError(err error) error {
	return fmt.Errorf("empty error (%w)", err)
}
func WrapClearError(err error) error {
	return fmt.Errorf("clear error (%w)", err)
}
func WrapUpdateError(err error) error {
	return fmt.Errorf("update error (%w)", err)
}
func WrapCreateError(err error) error {
	return fmt.Errorf("create error (%w)", err)
}
func WrapDeleteError(err error) error {
	return fmt.Errorf("delete error (%w)", err)
}
func WrapDropError(err error) error {
	return fmt.Errorf("drop error (%w)", err)
}
func WrapValidationError(err error) error {
	return fmt.Errorf("validate error (%w)", err)
}
func WrapDisconnectError(err error) error {
	return fmt.Errorf("disconnect error (%w)", err)
}
func WrapConnectError(err error) error {
	return fmt.Errorf("connect error (%w)", err)
}
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
func WrapEmailError(err error) error {
	return fmt.Errorf("send email error (%w)", err)
}
func WrapIDGenerationError(err error) error {
	return fmt.Errorf("ID generation error (%w)", err)
}

func WrapOSError(err error) error {
	return fmt.Errorf("OS error (%w)", err)
}

func WrapErrorForMocking(err error) error {
	return fmt.Errorf("mocking error (%w)", err)
}

func WrapNotificationError(err error) error {
	return fmt.Errorf("notification error (%w)", err)
}

func WrapFunctionalError(err error, f string) error {
	return fmt.Errorf("functional error - %v (%w)", f, err)
}

func WrapError(err error) error {
	logHandler.WarningLogger.Println("It is not advised to wrap errors without a specific error message")
	return fmt.Errorf("error (%w)", err)
}

func WrapInvalidFilterError(err error, f string) error {
	return fmt.Errorf("invalid filter [%v] (%w)", f, err)
}

func WrapInvalidHttpReturnStatusError(s string) error {
	return fmt.Errorf("inavalid/unsupported http return status [%v]", s)
}

func WrapInvalidHttpReturnStatusWithMessageError(status, message string) error {
	return fmt.Errorf("inavalid/unsupported http return status [%v] (%v)", status, message)
}

func WrapInvalidFieldError(f string) error {
	return fmt.Errorf("invalid field %v", f)
}

func WrapInvalidTypeError(f, d, s string) error {
	return fmt.Errorf("invalid type for field %v (%v != %v)", f, d, s)
}

func WrapRecordNotFoundError(table, field, id string) error {
	return fmt.Errorf("%v not found where (%v=%v)", table, field, id)
}

func WrapDAOUpdateAuditError(table string, id any, auditErr error) error {
	return fmt.Errorf("updating %v audit failed (ID=%v) %e", table, id, auditErr)
}

func WrapDAOCreateError(table string, id any, createErr error) error {
	return fmt.Errorf("creating %v failed (ID=%v) %e", table, id, createErr)
}

func WrapDAOInitialisationError(table string, initErr error) error {
	return fmt.Errorf("initialising %v failed %e", table, initErr)
}

func WrapDAOCaclulationError(table string, calcErr error) error {
	return fmt.Errorf("calculating %v failed %e", table, calcErr)
}

func WrapDAOValidationError(table string, valErr error) error {
	return fmt.Errorf("validating %v failed %e", table, valErr)
}

func WrapDAOUpdateError(table string, updateErr error) error {
	return fmt.Errorf("updating %v failed %e", table, updateErr)
}

func WrapDAODeleteError(table, field string, value any, deleteErr error) error {
	return fmt.Errorf("deleting %v failed (%v=%v) %e", table, field, value, deleteErr)
}

func WrapDAOReadError(table, field string, value any, readErr error) error {
	return fmt.Errorf("reading %v failed (%v=%v) %e", table, field, value, readErr)
}

func WrapDAOLookupError(table, field string, value any, lookupErr error) error {
	return fmt.Errorf("builing looking up for %v failed (key=%v,value=%v) %e", table, field, value, lookupErr)
}

func WrapDAONotInitialisedError(table, action string) error {
	return fmt.Errorf("%v DAO not initialised (Action=%v)", table, action)
}
