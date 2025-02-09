package pages

import (
	"github.com/mt1976/frantic-plum/application"
	"github.com/mt1976/frantic-plum/common"
	"github.com/mt1976/frantic-plum/logger"
	"github.com/mt1976/frantic-plum/paths"
	"github.com/mt1976/trnsl8r_service/app/business/translation"
)

var settings *common.Settings

func init() {
	settings = common.Get()
}

func New(title, action string) *Page {
	logger.EventLogger.Printf("Create New Page [%v],[%v]", title, action)

	p := Page{}
	p.ApplicationLogo = settings.GetLogoPath()
	p.ApplicationName = translation.Get(settings.GetApplicationName(), "")
	p.ApplicationDescription = translation.Get(settings.GetApplicationDescription(), "")
	p.ApplicationPrefix = translation.Get(settings.GetApplicationPrefix(), "")
	p.ApplicationFavicon = settings.GetFaviconPath()

	p.Message = ""
	p.MessageType = ""
	p.SingleItem = true
	p.ID = ""

	p.Delimiter = settings.GetDisplayDelimiter()
	p.ApplicationVersion = settings.GetApplicationVersion()
	p.ApplicationEnvironment = settings.GetApplicationEnvironment()
	p.ApplicationBuildDate = settings.GetApplicationReleaseDate()
	p.ApplicationCopyriteDate = settings.GetApplicationCopyright()
	p.OSSeperator = paths.Seperator()
	p.BackupLocation = paths.Backups().String()
	p.DumpLocation = paths.Dumps().String()
	p.DatabaseLocation = paths.Database().String()
	p.ApplicationPath = paths.Application().String()
	p.ApplicationOS = application.OS()

	p.PageAction = translation.Get(action, "")
	p.PageTitle = translation.Get(title, "")

	return &p
}
