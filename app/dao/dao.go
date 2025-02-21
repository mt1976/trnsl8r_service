package dao

import (
	"context"
	"strings"

	"github.com/mt1976/frantic-core/commonConfig"
	"github.com/mt1976/frantic-core/logHandler"
	"github.com/mt1976/trnsl8r_service/app/dao/textstore"
	"github.com/mt1976/trnsl8r_service/app/web/routes"
)

var name = "DAO"
var Version = 1
var tableName = "database"

func Initialise(cfg *commonConfig.Settings) error {
	logHandler.InfoLogger.Printf("[%v] Initialising...", strings.ToUpper(name))

	//database.Connect()
	textstore.Initialise(context.TODO())

	routes.Initialise(cfg)

	logHandler.InfoLogger.Printf("[%v] Initialised", strings.ToUpper(name))
	return nil
}
