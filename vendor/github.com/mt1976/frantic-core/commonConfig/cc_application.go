package commonConfig

import "strings"

func (s *Settings) GetApplication_Name() string {
	return s.Application.Name
}

func (s *Settings) GetApplication_Prefix() string {
	return s.Application.Prefix
}

func (s *Settings) GetApplication_HomePath() string {
	return s.Application.Home
}

func (s *Settings) GetApplication_Description() string {
	return s.Application.Description
}

func (s *Settings) GetApplication_Version() string {
	return s.Application.Version
}

func (s *Settings) GetApplication_Environment() string {
	return s.Application.Environment
}

func (s *Settings) GetApplication_ReleaseDate() string {
	return s.Application.ReleaseDate
}

func (s *Settings) GetApplication_Copyright() string {
	return s.Application.Copyright
}

func (s *Settings) GetApplication_Author() string {
	return s.Application.Author
}

func (s *Settings) GetApplication_License() string {
	return s.Application.License
}

func (s *Settings) GetApplication_Locale() string {
	if s.Application.Locale == "" {
		return "en_GB"
	}
	return s.Application.Locale
}

func (s *Settings) IsApplicationMode(inMode MODE) bool {
	// If first three chars of environment are "dev" then return "development"
	if strings.EqualFold(s.Server.Environment[:3], inMode.name[:3]) {
		return true
	}
	return false
}
