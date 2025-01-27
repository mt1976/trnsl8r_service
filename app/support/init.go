package support

import "github.com/mt1976/trnsl8r_service/app/support/config"

var TRUE = "true"
var FALSE = "false"

var d *config.Configuration

func init() {
	d = config.Get()
}
