package dao

import (
	"context"
	"strings"

	"github.com/mt1976/frantic-core/commonConfig"
	"github.com/mt1976/frantic-core/logHandler"
	"github.com/mt1976/trnsl8r_service/app/business/text"
	"github.com/mt1976/trnsl8r_service/app/dao/textStore"
	"github.com/mt1976/trnsl8r_service/app/web/routes"
)

var (
	name      = "DAO"
	Version   = 1
	tableName = "database"
)

func Initialise(cfg *commonConfig.Settings) error {
	logHandler.Info.Printf("[%v] Initialising...", strings.ToUpper(name))

	// database.Connect()
	textStore.Initialise(context.TODO(), false)
	textStore.RegisterCreator(text.Creator)
	textStore.RegisterImporter(text.Importer)

	routes.Initialise(cfg)

	logHandler.Info.Printf("[%v] Initialised", strings.ToUpper(name))
	return nil
}
