package pages

import (
	"github.com/mt1976/frantic-core/application"
	common "github.com/mt1976/frantic-core/commonConfig"
	logger "github.com/mt1976/frantic-core/logHandler"
	"github.com/mt1976/frantic-core/paths"
	"github.com/mt1976/trnsl8r_service/app/business/translate"
)

var settings *common.Settings

func init() {
	settings = common.Get()
}

func New(title, action string) *Page {
	logger.EventLogger.Printf("Create New Page [%v],[%v]", title, action)

	p := Page{}
	p.ApplicationLogo = settings.GetAssets_LogoPath()
	p.ApplicationName = translate.Get(settings.GetApplication_Name(), "")
	p.ApplicationDescription = translate.Get(settings.GetApplication_Description(), "")
	p.ApplicationPrefix = translate.Get(settings.GetApplication_Prefix(), "")
	p.ApplicationFavicon = settings.GetAssets_FaviconPath()

	p.Message = ""
	p.MessageType = ""
	p.SingleItem = true
	p.ID = ""

	p.Delimiter = settings.GetDisplayDelimiter()
	p.ApplicationVersion = settings.GetApplication_Version()
	p.ApplicationEnvironment = settings.GetApplication_Environment()
	p.ApplicationBuildDate = settings.GetApplication_ReleaseDate()
	p.ApplicationCopyriteDate = settings.GetApplication_Copyright()
	p.OSSeperator = paths.Seperator()
	p.BackupLocation = paths.Backups().String()
	p.DumpLocation = paths.Dumps().String()
	p.DatabaseLocation = paths.Database().String()
	p.ApplicationPath = paths.Application().String()
	p.ApplicationOS = application.OS()

	p.PageAction = translate.Get(action, "")
	p.PageTitle = translate.Get(title, "")

	return &p
}
