package textstore

// Data Access Object Template
// Version: 0.2.0
// Updated on: 2021-09-10

import (
	audit "github.com/mt1976/frantic-core/dao/audit"
)

// Text_Store represents a Text_Store entity.
type Text_Store struct {
	ID                int               `storm:"id,increment" csv:"-"` // primary key with auto increment
	Signature         string            `csv:"-"`                      // secondary key (used a unique identifier)
	Domain            string            `csv:"-"`                      // Saved for future use
	Type              string            `csv:"-"`                      // Saved for future use
	Original          string            `csv:"original"`               // Holds the original text, unformated.
	Message           string            `csv:"message"`                // Holds the translated text, formated.
	SourceApplication string            `csv:"-"`                      // Saved for future use
	SourceLocale      string            `csv:"-"`                      // Saved for future use
	ConsumedBy        []string          `csv:"-"`                      // Saved for future use
	Localised         map[string]string `csv:"-"`                      // Saved for future use
	Audit             audit.Audit       `csv:"-"`                      // Audit holds the audit information
}

// Define the field set as names
var (
	FIELD_ID                = "ID"
	FIELD_Signature         = "Signature"
	FIELD_Domain            = "Domain"
	FIELD_Type              = "Type"
	FIELD_Original          = "Original"
	FIELD_Message           = "Message"
	FIELD_SourceApplication = "SourceApplication"
	FIELD_SourceLocale      = "SourceLocale"
	FIELD_ConsumedBy        = "ConsumedBy"
	FIELD_Localised         = "Localised"
	FIELD_Audit             = "Audit"
)

var domain = "Text"
