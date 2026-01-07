package textStore

// Data Access Object Text
// Version: 0.3.0
// Updated on: 2025-12-31

//TODO: RENAME "Text" TO THE NAME OF THE DOMAIN ENTITY
//TODO: Update the Text_Store struct to match the domain entity
//TODO: Update the Fields. constants to match the domain entity

import (
	audit "github.com/mt1976/frantic-core/dao/audit"
	"github.com/mt1976/frantic-core/dao/database"
)

var Domain = "Text"
var TableName = Domain + "Store"

// TextStore represents a User entity.
type TextStore struct {
	// First three fields are mandatory for all DAO entities
	ID  int    `storm:"id,increment=100"` // primary key with auto increment
	Key string `storm:"unique"`           // key
	Raw string `storm:"unique"`           // raw ID before encoding
	// Add your domain entity fields below
	Signature         string            `csv:"-"`        // secondary key (used a unique identifier)
	Domain            string            `csv:"-"`        // Saved for future use
	Type              string            `csv:"-"`        // Saved for future use
	Original          string            `csv:"original"` // Holds the original text, unformated.
	Message           string            `csv:"message"`  // Holds the translated text, formated.
	SourceApplication string            `csv:"-"`        // Saved for future use
	SourceLocale      string            `csv:"-"`        // Saved for future use
	ConsumedBy        []string          `csv:"-"`        // Saved for future use
	Localised         map[string]string `csv:"-"`        // Saved for future use
	// Last field is mandatory for all DAO entities
	Audit audit.Audit `csv:"-"` // audit data
}

// Fields provides a structured way to reference model field names.
type fieldNames struct {
	// First four fields are mandatory for all DAO entities
	ID    database.Field
	Key   database.Field
	Raw   database.Field
	Audit database.Field
	// Add your domain entity fields below
	Signature         database.Field
	Domain            database.Field
	Type              database.Field
	Original          database.Field
	Message           database.Field
	SourceApplication database.Field
	SourceLocale      database.Field
	ConsumedBy        database.Field
	Localised         database.Field
}

// Fields provides a structured way to reference model field names.
var Fields = fieldNames{
	// First four fields are mandatory for all DAO entities
	ID:    "ID",
	Key:   "Key",
	Raw:   "Raw",
	Audit: "Audit",
	// Add your domain entity fields below
	Signature:         "Signature",
	Domain:            "Domain",
	Type:              "Type",
	Original:          "Original",
	Message:           "Message",
	SourceApplication: "SourceApplication",
	SourceLocale:      "SourceLocale",
	ConsumedBy:        "ConsumedBy",
	Localised:         "Localised",
}
