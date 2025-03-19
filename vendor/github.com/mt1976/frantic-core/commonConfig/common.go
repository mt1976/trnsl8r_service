package commonConfig

import (
	"fmt"
	"os"
	"strings"

	"github.com/BurntSushi/toml"
	"github.com/mt1976/frantic-core/paths"
)

// var Data ConfigurationModel
var name = "common"
var filename = ""
var commonSettingsFile = "common"

var TRUE = "true"
var FALSE = "false"

func Get() *Settings {

	var thisConfig Settings
	filename = paths.Application().String() + paths.Config().String() + string(os.PathSeparator) + commonSettingsFile + ".toml"
	content, err := os.ReadFile(filename)
	if err != nil {
		fmt.Printf("[%v] Error=[%v]", strings.ToUpper(name), err.Error())
		panic(err)
	}

	err = toml.Unmarshal(content, &thisConfig)
	if err != nil {
		panic(err)
	}

	return &thisConfig
}

func (s *Settings) Spew() {
	nm := strings.ToUpper(name)
	fmt.Printf("[%v] Config loaded from file: %v\n", nm, filename)
	fmt.Printf("[%v][APPLICATION]\n", nm)
	fmt.Printf("[%v] Name: %+v\n", nm, s)
}

func isTrueFalse(s string) bool {
	// We only disable the logging if the value is "true"/"t" or "yes"/"y"

	if s == "" {
		return false
	}

	logTrue := "true"
	if strings.EqualFold(s[:1], "y") {
		logTrue = "yes"
	}

	if strings.EqualFold(s, logTrue[:1]) {
		return true
	}
	if strings.EqualFold(s, logTrue) {
		return true
	}

	return false
}
