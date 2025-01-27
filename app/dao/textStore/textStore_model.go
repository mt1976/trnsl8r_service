package textStore

import (
	audit "github.com/mt1976/trnsl8r_service/app/dao/support/audit"
)

const (
	tableName = "texts" // table name in the database
)

// TextStore represents a TextStore entity.
type TextStore struct {
	ID        int         `storm:"id,increment" csv:"-"` // primary key with auto increment
	Signature string      `csv:"-"`                      // secondary key (used a unique identifier)
	Source    string      `csv:"-"`                      // Saved for future use
	Domain    string      `csv:"-"`                      // Saved for future use
	Type      string      `csv:"-"`                      // Saved for future use
	Locale    string      `csv:"-"`                      // Saved for future use
	Original  string      `csv:"original"`
	Message   string      `csv:"message"`
	Audit     audit.Audit `csv:"-"`
}

var Field_ID = "ID"
var Field_Signature = "Signature"
var Field_Source = "Source"
var Field_Domain = "Domain"
var Field_Type = "Type"
var Field_Locale = "Locale"
var Field_Original = "Original"
var Field_Message = "Message"
var Field_Audit = "Audit"

type TextImportModel struct {
	Original string `csv:"original"`
	Message  string `csv:"message"`
}
