package routes

import (
	"github.com/mt1976/frantic-core/commonConfig"
	"github.com/mt1976/frantic-core/logHandler"
)

var settings *commonConfig.Settings
var appName string
var msgTypeKey string
var msgTitleKey string
var msgContentKey string
var msgActionKey string
var trnsServerProtocol string
var trnsServerHost string
var trnsServerPort int
var trnsLocale string
var serverPort string
var serverHost string
var serverProtocol string

func init() {
	logHandler.EventLogger.Println("Loading Routes")
	settings = commonConfig.Get()
	appName = settings.GetApplication_Name()
	msgTypeKey = settings.GetMessageKey_Type()
	msgTitleKey = settings.GetMessageKey_Title()
	msgContentKey = settings.GetMessageKey_Content()
	msgActionKey = settings.GetMessageKey_Action()
	trnsServerProtocol = settings.GetTranslationServer_Protocol()
	trnsServerHost = settings.GetTranslationServer_Host()
	trnsServerPort = settings.GetTranslationServer_Port()
	trnsLocale = settings.GetTranslation_Locale()
	serverPort = settings.GetServer_PortString()
	serverHost = settings.GetServer_Host()
	serverProtocol = settings.GetServer_Protocol()

	//io.Dump("settings", paths.Dumps(), "test", "common", settings)

}
