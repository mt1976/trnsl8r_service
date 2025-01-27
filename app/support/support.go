package support

import (
	"strings"

	"github.com/mt1976/trnsl8r_service/app/support/logger"
)

//var SEP = config.DisplayDelimiter()

const PACKED = 1
const NOTPACKED = 0

func Banner(class, name, action string) {
	hdr := "------------------------------------------------------------------------"
	logger.InfoLogger.Println(hdr)
	logger.InfoLogger.Printf("[%v] Activity=[%v] - %v", strings.ToUpper(class), name, action)
	logger.InfoLogger.Println(hdr)
}

func ServiceBanner(class, name, action string) {
	hdr := "------------------------------------------------------------------------"
	logger.ServiceLogger.Println(hdr)
	logger.ServiceLogger.Printf("[%v] Activity=[%v] - %v", strings.ToUpper(class), name, action)
	logger.ServiceLogger.Println(hdr)
}
