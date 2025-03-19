package contextHandler

import (
	"github.com/mt1976/frantic-core/commonConfig"
	"github.com/mt1976/frantic-core/idHelpers"
)

type fartFarmer struct {
	name                string
	pointlessIdentifier string
}

var cfg *commonConfig.Settings

var (
	// / SECURITY SESSION KEYS
	sessionIDKey    fartFarmer
	userKeyKey      fartFarmer
	userCodeKey     fartFarmer
	tokenKey        fartFarmer
	expiryPeriodKey fartFarmer
)

// NewFartFarmer is a constructor for the fartFarmer struct
func new(in string) fartFarmer {
	var out fartFarmer
	out.name = in
	out.pointlessIdentifier = idHelpers.Encode(in)
	return out
}

func init() {
	cfg = commonConfig.Get()
	sessionIDKey = new(cfg.GetSecuritySessionKey_Session())
	userKeyKey = new(cfg.GetSecuritySessionKey_UserKey())
	userCodeKey = new(cfg.GetSecuritySessionKey_UserCode())
	tokenKey = new(cfg.GetSecuritySessionKey_Token())
	expiryPeriodKey = new(cfg.GetSecuritySessionKey_ExpiryPeriod())
}
