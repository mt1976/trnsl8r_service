package pages

import (
	"github.com/mt1976/frantic-plum/application"
	"github.com/mt1976/frantic-plum/config"
	"github.com/mt1976/frantic-plum/logger"
	"github.com/mt1976/frantic-plum/paths"
	"github.com/mt1976/trnsl8r_service/app/business/translation"
)

var cfg *config.Configuration

func init() {
	cfg = config.Get()
}

func New(title, action string) *Page {
	logger.EventLogger.Printf("Create New Page [%v],[%v]", title, action)

	p := Page{}
	p.ApplicationLogo = cfg.AssetsLogo()
	p.ApplicationName = translation.Get(cfg.ApplicationName())
	p.ApplicationDescription = translation.Get(cfg.ApplicationDescription())
	p.ApplicationPrefix = translation.Get(cfg.ApplicationPrefix())
	p.ApplicationFavicon = cfg.AssetsFavicon()

	p.Message = ""
	p.MessageType = ""
	p.SingleItem = true
	p.ID = ""

	p.Delimiter = cfg.DisplayDelimiter()
	p.ApplicationVersion = cfg.ApplicationVersion()
	p.ApplicationEnvironment = cfg.ApplicationEnvironment()
	p.ApplicationBuildDate = cfg.ApplicationReleaseDate()
	p.ApplicationCopyriteDate = cfg.ApplicationCopyright()
	p.OSSeperator = paths.Seperator()
	p.BackupLocation = paths.Backups().String()
	p.DumpLocation = paths.Dumps().String()
	p.DatabaseLocation = paths.Database().String()
	p.ApplicationPath = paths.Application().String()
	p.ApplicationOS = application.OS()

	p.PageAction = translation.Get(action)
	p.PageTitle = translation.Get(title)

	return &p
}
