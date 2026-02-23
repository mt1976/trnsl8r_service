// Data Access Object for the TextStore table
// Template Version: 0.6.00 - 2026-02-14
// Generated
// Date: 23/02/2026 & 12:36
// Who : matttownsend (orion)

package textStore

import (
	"sync"

	"github.com/mt1976/frantic-amphora/dao/audit"
	"github.com/mt1976/frantic-amphora/dao/entities"
)

// TableName is the canonical DAO table identifier for this package.
var (
	TableName              = entities.Table("TextStore")
	tableName              = TableName.String()
	URI       entities.URI = entities.URI("text")
	PK        entities.KEY = entities.KEY("textkey")
	RAW       entities.KEY = entities.KEY("textalt")
)

// The TextStore struct defines the data model for the TextStore table.
// Adjust domain fields and tags as required in the TextStore.definitions file.
type TextStore struct {
	// The primary key field(s), managed by the framework, DO NOT MODIFY
	ID  int    `storm:"id,increment=100"`
	Key string `storm:"index,unique"`
	Raw string `storm:"index,unique"`
	// Audit information, managed by the framework, DO NOT MODIFY
	Audit audit.Audit `csv:"-"`
	// Locking information, managed by the framework, DO NOT MODIFY
	Lock sync.Mutex `csv:"-"` // Add this field to enable record locking for concurrent updates

	// Domain specific fields
	//
	Signature         string            `csv:"-"`        // secondary key (used a unique identifier)
	Domain            string            `csv:"-"`        // Saved for future use
	Type              string            `csv:"-"`        // Saved for future use
	Original          string            `csv:"original"` // Holds the original text, unformated.
	Message           string            `csv:"message"`  // Holds the translated text, formated.
	SourceApplication string            `csv:"-"`        // Saved for future use
	SourceLocale      string            `csv:"-"`        // Saved for future use
	ConsumedBy        []string          `csv:"-"`        // Saved for future use
	Localised         map[string]string `csv:"-"`        // Saved for future use
	// Add no more fields below this line
}

type fieldNames struct {
	// The primary key field(s), managed by the framework, DO NOT MODIFY
	ID  entities.Field
	Key entities.Field
	Raw entities.Field
	// The audit information, managed by the framework, DO NOT MODIFY
	Audit entities.Field
	// Domain specific fields
	Signature         entities.Field
	Domain            entities.Field
	Type              entities.Field
	Original          entities.Field
	Message           entities.Field
	SourceApplication entities.Field
	SourceLocale      entities.Field
	ConsumedBy        entities.Field
	Localised         entities.Field

	// Add no more fields below this line
}

// Fields provides strongly-typed field names for use with GetBy/GetAllWhere/etc.
//
// Example: GetBy(Fields.Key, "abc")
//
// Note: the values are the struct field names as stored in Storm.
var Fields = fieldNames{
	// The primary key field(s), managed by the framework, DO NOT MODIFY
	ID:  "ID",
	Key: "Key",
	Raw: "Raw",
	// The audit information, managed by the framework, DO NOT MODIFY
	Audit: "Audit",
	// tableName-specific fields, please modify as required
	Signature:         "Signature",
	Domain:            "Domain",
	Type:              "Type",
	Original:          "Original",
	Message:           "Message",
	SourceApplication: "SourceApplication",
	SourceLocale:      "SourceLocale",
	ConsumedBy:        "ConsumedBy",
	Localised:         "Localised",
	// Add no more fields below this line
}
