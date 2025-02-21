package textstore

import (
	"context"
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
	"time"

	"github.com/gocarina/gocsv"
	"github.com/mt1976/frantic-core/commonConfig"
	"github.com/mt1976/frantic-core/commonErrors"
	"github.com/mt1976/frantic-core/dao/audit"
	"github.com/mt1976/frantic-core/idHelpers"
	"github.com/mt1976/frantic-core/logHandler"
	"github.com/mt1976/frantic-core/paths"
	"github.com/mt1976/frantic-core/stringHelpers"
	"github.com/mt1976/trnsl8r_service/app/business/domains"
)

type TextImportModel struct {
	Original string `csv:"original"`
	Message  string `csv:"message"`
}

var COMMA = '|'

func ExportCSV() error {
	logHandler.ExportLogger.Printf("Exporting texts")
	Initialise(context.TODO())

	textsFile := openTextsFile("export", logHandler.ExportLogger)
	defer textsFile.Close()

	texts, err := GetAll()
	if err != nil {
		logHandler.ExportLogger.Printf("Error Getting all texts: %v", err.Error())
	}

	gocsv.SetCSVWriter(func(out io.Writer) *gocsv.SafeCSVWriter {
		writer := csv.NewWriter(out)
		writer.Comma = COMMA // Use tab-delimited format
		writer.UseCRLF = true
		return gocsv.NewSafeCSVWriter(writer)
	})

	_, err = gocsv.MarshalString(texts) // Get all texts as CSV string
	if err != nil {
		logHandler.ExportLogger.Printf("Error exporting texts: %v", err.Error())
	}
	err = gocsv.MarshalFile(&texts, textsFile) // Get all texts as CSV string
	if err != nil {
		logHandler.ExportLogger.Printf("Error exporting texts: %v", err.Error())
	}

	msg := fmt.Sprintf("# Generated (%v) texts at %v on %v", len(texts), time.Now().Format("15:04:05"), time.Now().Format("2006-01-02"))
	textsFile.WriteString(msg)

	textsFile.Close()

	logHandler.ExportLogger.Printf("Exported (%v) texts", len(texts))

	return nil
}

func openTextsFile(in string, log *log.Logger) *os.File {
	exportPath := paths.Defaults()
	textsFileName := fmt.Sprintf("%s%s/%s", paths.Application().String(), exportPath, "translations.csv")

	// fmt.Printf("exportPath: %v\n", exportPath)
	// fmt.Printf("textsFile: %v\n", textsFileName)

	textsFile, err := os.OpenFile(textsFileName, os.O_RDWR|os.O_CREATE, os.ModePerm)
	if err != nil {
		log.Panicf("Error opening file: %v", err)
		panic(err)
	}
	//fmt.Printf("textsFile.Name(): %v\n", textsFile.Name())
	log.Printf("Import/Export=[%v] File=[%v]", in, textsFile.Name())
	return textsFile
}

func ImportCSV() error {
	logHandler.ImportLogger.Printf("Importing texts")
	Initialise(context.TODO())
	csvFile := openTextsFile("import", logHandler.ImportLogger)
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
		logHandler.ImportLogger.Printf("Importing %v: %v - No Content", domains.TEXT.String(), err.Error())
		csvFile.Close()
		return nil
	}

	if _, err := csvFile.Seek(0, 0); err != nil { // Go to the start of the file
		panic(err)
	}
	noTexts := len(texts)
	for thisPos, textEntry := range texts {
		//fmt.Printf("% 3v)[%v][%v][%v]\n", i, textEntry.Original, textEntry.Message, textEntry)
		logHandler.ImportLogger.Printf("importing text (%v/%v) [%v]", thisPos+1, noTexts, textEntry.Message)

		existingText, _ := GetBySignature(idHelpers.Encode(textEntry.Original))
		if existingText.Signature != "" {
			//logger.InfoLogger.Printf("Text already exists: [%v]", textEntry.Message)
			continue
		}

		_, err := load(textEntry.Original, textEntry.Message)
		if err != nil {
			logHandler.ImportLogger.Panicf("importing text: %v", err.Error())
		}
		logHandler.ImportLogger.Printf("imported text [%v] [%v]", textEntry.Original, textEntry.Message)
	}

	logHandler.ImportLogger.Printf("imported (%v) texts", len(texts))
	csvFile.Close()
	return nil
}

func load(original, message string) (Text_Store, error) {

	//logger.InfoLogger.Printf("ACT: NEW New %v %v %v", tableName, name, destination)
	u := Text_Store{}
	u.Signature = idHelpers.Encode(strings.ToUpper(original))
	u.Message = message
	u.Original = message
	// Add basic attributes

	// Record the create action in the audit data
	_ = u.Audit.Action(context.TODO(), audit.IMPORT.WithMessage(fmt.Sprintf("Imported text [%v]", message)))

	dupe, err := u.dup(u.Signature)
	// Log the dest instance before the creation
	if err == commonErrors.ErrorDuplicate {
		// This is OK, do nothing as this is a duplicate record
		// we ignore duplicate destinations.
		logHandler.ImportLogger.Printf("[DUPLICATE] %v already available", stringHelpers.DQuote(message))
		return dupe, nil
	}

	//u.Dump(!,"Post-Prepare-DupCheck")

	if err != nil {
		if err != commonErrors.ErrorDuplicate {
			// Log and return the error if there is an error checking for duplicates
			logHandler.ImportLogger.Panicf("Duplicated Detected=[%s]", err.Error())
			return Text_Store{}, err
		}
		logHandler.ImportLogger.Panicf("Error=[%s]", err.Error())
		return Text_Store{}, err
	}

	// Save the dest instance to the database
	if u.Signature == "" {
		logHandler.ImportLogger.Printf("[%v] ID is required, skipping", strings.ToUpper(domain))
		return Text_Store{}, nil
	}

	set := commonConfig.Get()

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

	err = activeDB.Create(&u)
	if err != nil {
		// Log and panic if there is an error creating the dest instance
		logHandler.ImportLogger.Panicf("[%v] Create %s", strings.ToUpper(domain), err.Error())
		panic(err)
	}

	//u.Dump(!,fmt.Sprintf("PostNew_dest_%d", u.ID))
	msg := fmt.Sprintf("Imported text translation available Id=[%v] Message=[%v]", original, message)
	logHandler.ImportLogger.Println(msg)
	// Return the created dest and nil error
	return u, nil
}
