package html

import (
	"strconv"
	"strings"

	"github.com/mt1976/trnsl8r_service/app/support/config"
	"github.com/mt1976/trnsl8r_service/app/support/logger"
)

var name = "HTML"
var d *config.Configuration

func init() {
	d = config.Get()
}

func ValueToInt(s string) int {
	if s == "" {
		return 0
	}
	i, err := strconv.Atoi(s)
	if err != nil {
		// ... handle error
		logger.WarningLogger.Printf("[%v] [Math] Error=[%v]", strings.ToUpper(name), err.Error())
		return 999999999
	}
	return i
}

func ValueToBool(s string) bool {
	return s == "on"
}
