package textStore

import "github.com/mt1976/frantic-plum/dao/audit"

const (
	tableName = "texts" // table name in the database
)

// TextStore represents a TextStore entity.
type TextStore struct {
	ID         int               `storm:"id,increment" csv:"-"` // primary key with auto increment
	Signature  string            `csv:"-"`                      // secondary key (used a unique identifier)
	Source     string            `csv:"-"`                      // Saved for future use
	Domain     string            `csv:"-"`                      // Saved for future use
	Type       string            `csv:"-"`                      // Saved for future use
	Locale     string            `csv:"-"`                      // Saved for future use
	Original   string            `csv:"original"`               // Holds the original text, unformated.
	Message    string            `csv:"message"`                // Holds the translated text, formated.
	ConsumedBy []string          `csv:"-"`                      // Saved for future use
	Localised  map[string]string `csv:"-"`                      // Saved for future use
	Audit      audit.Audit       `csv:"-"`                      // Audit holds the audit information
}

var Field_ID = "ID"
var Field_Signature = "Signature"
var Field_Source = "Source"
var Field_Domain = "Domain"
var Field_Type = "Type"
var Field_Locale = "Locale"
var Field_Original = "Original"
var Field_Message = "Message"
var Field_ConsumedBy = "ConsumedBy"
var Field_Localised = "Localised"
var Field_Audit = "Audit"

type TextImportModel struct {
	Original string `csv:"original"`
	Message  string `csv:"message"`
}
