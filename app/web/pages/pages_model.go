package pages

import (
	dao "github.com/mt1976/frantic-core/dao/lookup"
	"github.com/mt1976/trnsl8r_service/app/dao/textstore"
)

type Page struct {
	ApplicationName         string
	ApplicationDescription  string
	ApplicationPrefix       string
	ApplicationLogo         string
	ApplicationFavicon      string
	ApplicationVersion      string
	ApplicationEnvironment  string
	ApplicationBuildDate    string
	ApplicationCopyriteDate string
	ApplicationPath         string
	ApplicationOS           string
	Delimiter               string
	PageAction              string
	PageTitle               string
	Message                 string
	MessageType             string
	SingleItem              bool
	ID                      string

	OSSeperator      string
	BackupLocation   string
	DumpLocation     string
	DatabaseLocation string

	TextList    []textstore.Text_Store
	TextItem    textstore.Text_Store
	NoTextItems int
	HostsSelect dao.Lookup
}
