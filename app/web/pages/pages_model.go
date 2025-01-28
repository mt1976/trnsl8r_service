package pages

import (
	dao "github.com/mt1976/frantic-plum/dao/lookup"
	"github.com/mt1976/trnsl8r_service/app/dao/textStore"
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

	TextList    []textStore.TextStore
	TextItem    textStore.TextStore
	NoTextItems int
	HostsSelect dao.Lookup
}
