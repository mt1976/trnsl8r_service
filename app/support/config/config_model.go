package config

type Configuration struct {
	Application struct {
		Name        string `toml:"name"`
		Prefix      string `toml:"prefix"`
		Home        string `toml:"home"`
		Description string `toml:"description"`
		Version     string `toml:"version"`
		Environment string `toml:"environment"`
		ReleaseDate string `toml:"releaseDate"`
		Copyright   string `toml:"copyright"`
		Author      string `toml:"author"`
		License     string `toml:"license"`
		Locale      string `toml:"locale"`
	} `toml:"Application"`
	Server struct {
		Host        string `toml:"host"`
		Port        int    `toml:"port"`
		Protocol    string `toml:"protocol"`
		Environment string `toml:"environment"`
	} `toml:"Server"`
	Translation struct {
		Host     string `toml:"host"`
		Port     int    `toml:"port"`
		Locale   string `toml:"locale"`
		Protocol string `toml:"protocol"`
	} `toml:"Translation"`
	Assets struct {
		Logo    string `toml:"logo"`
		Favicon string `toml:"favicon"`
	} `toml:"Assets"`
	Dates struct {
		DateTimeFormat string `toml:"dateTimeFormat"`
		DateFormat     string `toml:"dateFormat"`
		TimeFormat     string `toml:"timeFormat"`
		Backup         string `toml:"backup"`
		BackupFolder   string `toml:"backupFolder"`
		Human          string `toml:"human"`
		DMY2           string `toml:"dmy2"`
		YMD            string `toml:"ymd"`
		Internal       string `toml:"internal"`
	} `toml:"Dates"`
	History struct {
		MaxEntries int `toml:"maxEntries"`
	} `toml:"History"`
	Message struct {
		TypeKey    string `toml:"typeKey"`
		TitleKey   string `toml:"titleKey"`
		ContentKey string `toml:"contentKey"`
		ActionKey  string `toml:"actionKey"`
	} `toml:"Message"`
	Display struct {
		Delimiter string `toml:"delim"`
	} `toml:"Display"`
}
