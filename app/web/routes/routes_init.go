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
	appName = settings.GetApplicationName()
	msgTypeKey = settings.GetMessageTypeKey()
	msgTitleKey = settings.GetMessageTitleKey()
	msgContentKey = settings.GetMessageContentKey()
	msgActionKey = settings.GetMessageActionKey()
	trnsServerProtocol = settings.GetTranslationServerProtocol()
	trnsServerHost = settings.GetTranslationServerHost()
	trnsServerPort = settings.GetTranslationServerPort()
	trnsLocale = settings.GetTranslationLocale()
	serverPort = settings.GetServerPortAsString()
	serverHost = settings.GetServerHost()
	serverProtocol = settings.GetServerProtocol()

	//io.Dump("settings", paths.Dumps(), "test", "common", settings)

}
