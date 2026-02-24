// Data Access Object for the TextStore table
// Template Version: 0.6.00 - 2026-02-14
// Generated 
// Date: 24/02/2026 & 10:04
// Who : matttownsend (orion)

package textStore

import (
	"github.com/goforj/godump"
	"github.com/mt1976/frantic-core/logHandler"
)

// Spew outputs the contents of the record to the Trace log.
func (record *TextStore) Spew() {
	logHandler.Trace.Printf("[%v] Record=[%+v]", tableName, godump.DumpStr(record))
}
