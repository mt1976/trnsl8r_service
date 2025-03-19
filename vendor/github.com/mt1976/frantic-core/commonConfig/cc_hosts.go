package commonConfig

func (s *Settings) GetValidHosts() []struct {
	Name string "toml:\"name\""
	FQDN string "toml:\"fqdn\""
	IP   string "toml:\"ip\""
	Zone string "toml:\"zone\""
} {
	return s.Hosts
}
