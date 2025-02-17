package dao

import (
	"strings"

	storm "github.com/asdine/storm/v3"
	common "github.com/mt1976/frantic-core/commonConfig"
	"github.com/mt1976/frantic-core/dao/database"
	logger "github.com/mt1976/frantic-core/logHandler"
	"github.com/mt1976/trnsl8r_service/app/web/routes"
)

var name = "DAO"
var Version = 1
var DB *storm.DB
var tableName = "database"

func Initialise(cfg *common.Settings) error {
	logger.InfoLogger.Printf("[%v] Initialising...", strings.ToUpper(name))

	database.Connect()

	routes.Initialise(cfg)

	logger.InfoLogger.Printf("[%v] Initialised", strings.ToUpper(name))
	return nil
}
