package textstore

// Data Access Object Template
// Version: 0.2.0
// Updated on: 2021-09-10

import (
	"context"

	"github.com/mt1976/frantic-core/dao/actions"
	"github.com/mt1976/frantic-core/dao/database"
	"github.com/mt1976/frantic-core/logHandler"
	"github.com/mt1976/frantic-core/timing"
)

var activeDB *database.DB
var initialised bool = false // default to false

func Initialise(ctx context.Context) {
	logHandler.InfoLogger.Printf("Initialising %v", domain)
	timing := timing.Start(domain, actions.INITIALISE.GetCode(), "Initialise")
	//	cfg := commonConfig.Get()

	// For a specific database connection, use NamedConnect, otherwise use Connect
	//activeDB = database.ConnectToNamedDB("Template")
	activeDB = database.Connect()
	initialised = true
	timing.Stop(1)
	logHandler.InfoLogger.Printf("Initialised %v", domain)
}
