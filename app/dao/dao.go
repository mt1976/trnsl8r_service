package dao

import (
	"strings"

	storm "github.com/asdine/storm/v3"
	"github.com/mt1976/trnsl8r_service/app/support/config"
	"github.com/mt1976/trnsl8r_service/app/support/logger"
	"github.com/mt1976/trnsl8r_service/app/web/routes"
)

var name = "DAO"
var Version = 1
var DB *storm.DB
var tableName = "database"

func Initialise(cfg *config.Configuration) error {
	logger.InfoLogger.Printf("[%v] Initialising...", strings.ToUpper(name))

	routes.Initialise(cfg)

	logger.InfoLogger.Printf("[%v] Initialised", strings.ToUpper(name))
	return nil
}
