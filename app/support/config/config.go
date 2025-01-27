package config

import (
	"fmt"
	"os"
	"strings"

	"github.com/BurntSushi/toml"
	"github.com/mt1976/trnsl8r_service/app/support/paths"
)

// var Data ConfigurationModel
var name = "config"
var filename = ""
var settingsFile = "trnsl8r_service"

func Get() *Configuration {
	var thisConfig Configuration
	filename = paths.Application().String() + paths.Config().String() + string(os.PathSeparator) + settingsFile + ".toml"
	content, err := os.ReadFile(filename)
	if err != nil {
		fmt.Printf("[%v] Error=[%v]", strings.ToUpper(name), err.Error())
		panic(err)
	}

	toml.Unmarshal(content, &thisConfig)

	return &thisConfig
}

func (d *Configuration) Spew() {
	nm := strings.ToUpper(name)
	f2 := "[%v] %-15v : %v\n"
	fmt.Printf("[%v] Config loaded from file: %v\n", nm, filename)
	fmt.Printf("[%v][APPLICATION]\n", nm)
	fmt.Printf(f2, nm, "Name", d.Application.Name)
	fmt.Printf(f2, nm, "Version", d.Application.Version)
	fmt.Printf(f2, nm, "Description", d.Application.Description)
	fmt.Printf(f2, nm, "Prefix", d.Application.Prefix)
	fmt.Printf(f2, nm, "Environment", d.Application.Environment)
	fmt.Printf(f2, nm, "ReleaseDate", d.Application.ReleaseDate)
	fmt.Printf(f2, nm, "Copyright", d.Application.Copyright)
	fmt.Printf(f2, nm, "License", d.Application.License)
	fmt.Printf(f2, nm, "Author", d.Application.Author)

	fmt.Printf("[%v][SERVER]\n", nm)

	fmt.Printf(f2, nm, "Port", d.Server.Port)
	fmt.Printf(f2, nm, "Protocol", d.Server.Protocol)
	fmt.Printf(f2, nm, "Environment", d.Server.Environment)

	fmt.Printf("[%v][ASSETS]\n", nm)

	fmt.Printf(f2, nm, "Logo", d.Assets.Logo)
	fmt.Printf(f2, nm, "Favicon", d.Assets.Favicon)

	fmt.Printf("[%v][DATES]\n", nm)

	fmt.Printf(f2, nm, "DateTimeFormat", d.Dates.DateTimeFormat)
	fmt.Printf(f2, nm, "DateFormat", d.Dates.DateFormat)
	fmt.Printf(f2, nm, "TimeFormat", d.Dates.TimeFormat)
	fmt.Printf(f2, nm, "Backup", d.Dates.Backup)
	fmt.Printf(f2, nm, "BackupFolder", d.Dates.BackupFolder)
	fmt.Printf(f2, nm, "Human", d.Dates.Human)
	fmt.Printf(f2, nm, "DMY2", d.Dates.DMY2)
	fmt.Printf(f2, nm, "YMD", d.Dates.YMD)
	fmt.Printf(f2, nm, "Internal", d.Dates.Internal)

	fmt.Printf("[%v][HISTORY]\n", nm)
	fmt.Printf(f2, nm, "MaxEntries", d.History.MaxEntries)

	fmt.Printf("[%v][MESSAGE]\n", nm)
	fmt.Printf(f2, nm, "TypeKey", d.Message.TypeKey)
	fmt.Printf(f2, nm, "TitleKey", d.Message.TitleKey)
	fmt.Printf(f2, nm, "ContentKey", d.Message.ContentKey)
	fmt.Printf(f2, nm, "ActionKey", d.Message.ActionKey)

	fmt.Printf("[%v][DISPLAY]\n", nm)
	fmt.Printf(f2, nm, "Delimiter", d.Display.Delimiter)

}
