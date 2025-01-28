package textStore

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/gocarina/gocsv"
	"github.com/mt1976/frantic-plum/dao/audit"
	"github.com/mt1976/frantic-plum/dao/database"
	"github.com/mt1976/frantic-plum/errors"
	"github.com/mt1976/frantic-plum/id"
	"github.com/mt1976/frantic-plum/logger"
	"github.com/mt1976/frantic-plum/paths"
	"github.com/mt1976/trnsl8r_service/app/business/domains"
)

func ExportCSV() error {

	textsFile := openTextsFile()
	defer textsFile.Close()

	texts, err := GetAll()
	if err != nil {
		logger.ErrorLogger.Printf("Error Getting all texts: %v", err.Error())
	}

	csvContent, err := gocsv.MarshalString(&texts) // Get all texts as CSV string
	if err != nil {
		logger.ErrorLogger.Printf("Error exporting texts: %v", err.Error())
	}

	fmt.Printf("csvContent: %v\n", csvContent)

	gocsv.SetCSVWriter(func(out io.Writer) *gocsv.SafeCSVWriter {
		writer := csv.NewWriter(out)
		writer.Comma = '|' // Use tab-delimited format
		return gocsv.NewSafeCSVWriter(writer)
	})

	err = gocsv.MarshalFile(&texts, textsFile) // Get all texts as CSV string
	if err != nil {
		logger.ErrorLogger.Printf("Error exporting texts: %v", err.Error())
	}

	return nil
}

func openTextsFile() *os.File {
	exportPath := paths.Defaults()
	textsFileName := fmt.Sprintf("%s%s/%s", paths.Application().String(), exportPath, "translations.csv")
	logger.InfoLogger.Printf("Export File: [%v]", textsFileName)
	// fmt.Printf("exportPath: %v\n", exportPath)
	// fmt.Printf("textsFile: %v\n", textsFileName)

	textsFile, err := os.OpenFile(textsFileName, os.O_RDWR|os.O_CREATE, os.ModePerm)
	if err != nil {
		panic(err)
	}

	return textsFile
}

func ImportCSV() error {

	csvFile := openTextsFile()
	defer csvFile.Close()
	// fmt.Printf("textsFile: %v\n", csvFile.Name())

	gocsv.SetCSVReader(func(in io.Reader) gocsv.CSVReader {
		r := csv.NewReader(in)
		r.Comma = '|'
		return r // Allows use pipe as delimiter
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

	for _, textEntry := range texts {
		//fmt.Printf("% 3v)[%v][%v][%v]\n", i, textEntry.Original, textEntry.Message, textEntry)
		logger.EventLogger.Printf("Importing text [%v] [%v]", textEntry.Original, textEntry.Message)

		existingText, _ := GetBySignature(id.Encode(textEntry.Original))
		if existingText.Signature != "" {
			logger.InfoLogger.Printf("Text already exists: [%v] [%v]", textEntry.Original, textEntry.Message)
			continue
		}

		_, err := load(textEntry.Original, textEntry.Message)
		if err != nil {
			logger.ErrorLogger.Printf("Error importing text: %v", err.Error())
		}
		logger.EventLogger.Printf("Imported text [%v] [%v]", textEntry.Original, textEntry.Message)
	}

	logger.EventLogger.Printf("Imported %v texts", len(texts))

	return nil
}

func load(original, message string) (TextStore, error) {

	//logger.InfoLogger.Printf("ACT: NEW New %v %v %v", tableName, name, destination)

	// Create a new d
	u := TextStore{}
	u.Signature = id.Encode(original)
	u.Message = message
	u.Original = message
	// Add basic attributes

	// Record the create action in the audit data
	_ = u.Audit.Action(nil, audit.IMPORT.WithMessage(fmt.Sprintf("Imported text [%v]", message)))

	// Log the dest instance before the creation
	xtext, err := u.prepare()
	if err == errors.ErrorDuplicate {
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
	if u.Signature == "" {
		logger.WarningLogger.Printf("[%v] ID is required, skipping", strings.ToUpper(tableName))
		return TextStore{}, nil
	}

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
