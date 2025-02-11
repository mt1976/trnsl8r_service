package textStore

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"strings"
	"time"

	"github.com/gocarina/gocsv"
	"github.com/mt1976/frantic-core/common"
	"github.com/mt1976/frantic-core/commonErrors"
	"github.com/mt1976/frantic-core/dao/audit"
	"github.com/mt1976/frantic-core/dao/database"
	"github.com/mt1976/frantic-core/id"
	"github.com/mt1976/frantic-core/logger"
	"github.com/mt1976/frantic-core/paths"
	"github.com/mt1976/trnsl8r_service/app/business/domains"
)

var COMMA = '|'

func ExportCSV() error {

	textsFile := openTextsFile("export")
	defer textsFile.Close()

	texts, err := GetAll()
	if err != nil {
		logger.ErrorLogger.Printf("Error Getting all texts: %v", err.Error())
	}

	gocsv.SetCSVWriter(func(out io.Writer) *gocsv.SafeCSVWriter {
		writer := csv.NewWriter(out)
		writer.Comma = COMMA // Use tab-delimited format
		writer.UseCRLF = true
		return gocsv.NewSafeCSVWriter(writer)
	})

	_, err = gocsv.MarshalString(texts) // Get all texts as CSV string
	if err != nil {
		logger.ErrorLogger.Printf("Error exporting texts: %v", err.Error())
	}
	err = gocsv.MarshalFile(&texts, textsFile) // Get all texts as CSV string
	if err != nil {
		logger.ErrorLogger.Printf("Error exporting texts: %v", err.Error())
	}

	msg := fmt.Sprintf("# Generated (%v) texts at %v on %v", len(texts), time.Now().Format("15:04:05"), time.Now().Format("2006-01-02"))
	textsFile.WriteString(msg)

	textsFile.Close()

	logger.EventLogger.Printf("Exported (%v) texts", len(texts))

	return nil
}

func openTextsFile(in string) *os.File {
	exportPath := paths.Defaults()
	textsFileName := fmt.Sprintf("%s%s/%s", paths.Application().String(), exportPath, "translations.csv")

	// fmt.Printf("exportPath: %v\n", exportPath)
	// fmt.Printf("textsFile: %v\n", textsFileName)

	textsFile, err := os.OpenFile(textsFileName, os.O_RDWR|os.O_CREATE, os.ModePerm)
	if err != nil {
		panic(err)
	}
	//fmt.Printf("textsFile.Name(): %v\n", textsFile.Name())
	logger.InfoLogger.Printf("Import/Export=[%v] File=[%v]", in, textsFile.Name())
	return textsFile
}

func ImportCSV() error {

	csvFile := openTextsFile("import")
	defer csvFile.Close()
	// fmt.Printf("textsFile: %v\n", csvFile.Name())

	gocsv.SetCSVReader(func(in io.Reader) gocsv.CSVReader {
		r := csv.NewReader(in)    // Allows use pipe as delimiter
		r.Comma = COMMA           // Use tab-delimited format
		r.Comment = '#'           // Ignore comment lines
		r.TrimLeadingSpace = true // Trim leading space
		return r                  // Allows use pipe as delimiter
	})

	texts := []*TextImportModel{}

	if err := gocsv.UnmarshalFile(csvFile, &texts); err != nil { // Load clients from file
		logger.WarningLogger.Printf("Importing %v: %v - No Content", domains.TEXT.String(), err.Error())
		csvFile.Close()
		return nil
	}

	if _, err := csvFile.Seek(0, 0); err != nil { // Go to the start of the file
		panic(err)
	}
	noTexts := len(texts)
	for thisPos, textEntry := range texts {
		//fmt.Printf("% 3v)[%v][%v][%v]\n", i, textEntry.Original, textEntry.Message, textEntry)
		logger.ServiceLogger.Printf("Importing text (%v/%v) [%v]", thisPos+1, noTexts, textEntry.Message)

		existingText, _ := GetBySignature(id.Encode(textEntry.Original))
		if existingText.Signature != "" {
			//logger.InfoLogger.Printf("Text already exists: [%v]", textEntry.Message)
			continue
		}

		_, err := load(textEntry.Original, textEntry.Message)
		if err != nil {
			logger.ErrorLogger.Printf("Error importing text: %v", err.Error())
		}
		logger.ServiceLogger.Printf("Imported text [%v] [%v]", textEntry.Original, textEntry.Message)
	}

	logger.ServiceLogger.Printf("Imported (%v) texts", len(texts))
	csvFile.Close()
	return nil
}

func load(original, message string) (TextStore, error) {

	//logger.InfoLogger.Printf("ACT: NEW New %v %v %v", tableName, name, destination)
	u := TextStore{}
	u.Signature = id.Encode(strings.ToUpper(original))
	u.Message = message
	u.Original = message
	// Add basic attributes

	// Record the create action in the audit data
	_ = u.Audit.Action(nil, audit.IMPORT.WithMessage(fmt.Sprintf("Imported text [%v]", message)))

	// Log the dest instance before the creation
	xtext, err := u.validateRecord()
	if err == commonErrors.ErrorDuplicate {
		// This is OK, do nothing as this is a duplicate record
		// we ignore duplicate destinations.
		logger.WarningLogger.Printf("DUPLICATE %v available in use as [%v]", message, u.Signature)
		return xtext, nil
	}

	//u.Dump(!,"Post-Prepare-DupCheck")

	if err != nil {
		logger.ErrorLogger.Printf("Error=[%s]", err.Error())
		return TextStore{}, err
	}

	// Save the dest instance to the database
	if u.Signature == "" {
		logger.WarningLogger.Printf("[%v] ID is required, skipping", strings.ToUpper(tableName))
		return TextStore{}, nil
	}

	set := common.Get()

	locales := set.GetLocales()
	//noLocales := len(locales)

	newTextLocalised := make(map[string]string)

	for _, locale := range locales {

		// if the locale map is empty create it
		if u.Localised == nil {
			u.Localised = make(map[string]string)
		}

		newTextLocalised[locale.Key] = u.Localised[locale.Key]

	}
	u.Localised = newTextLocalised

	err = database.Create(&u)
	if err != nil {
		// Log and panic if there is an error creating the dest instance
		logger.ErrorLogger.Printf("[%v] Create %s", strings.ToUpper(tableName), err.Error())
		panic(err)
	}

	//u.Dump(!,fmt.Sprintf("PostNew_dest_%d", u.ID))
	msg := fmt.Sprintf("Imported text translation available Id=[%v] Message=[%v]", original, message)
	logger.TranslationLogger.Println(msg)
	// Return the created dest and nil error
	logger.AuditLogger.Printf("[%v] [%v] ID=[%v] Notes[%v]", strings.ToUpper(tableName), audit.CREATE.Code(), original, msg)

	return u, nil
}
