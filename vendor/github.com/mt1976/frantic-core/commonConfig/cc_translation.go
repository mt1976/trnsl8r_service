package commonConfig

import "strconv"

// / TRANSLATION

func (s *Settings) GetTranslationServer_Host() string {
	if s.Translation.Host == "" {
		return "localhost"
	}
	return s.Translation.Host
}

func (s *Settings) GetTranslationServer_Port() int {
	if s.Translation.Port == 0 {
		return s.GetDatabase_Port() + 1
	}
	return s.Translation.Port
}

func (s *Settings) GetTranslation_Locale() string {
	if s.Translation.Locale == "" {
		return s.GetApplication_Locale()
	}
	return s.Translation.Locale
}

func (s *Settings) GetTranslationServer_Protocol() string {
	if s.Translation.Protocol == "" {
		return "http"
	}
	return s.Translation.Protocol
}

func (s *Settings) GetTranslationServer_PortString() string {
	return strconv.Itoa(s.GetTranslationServer_Port())
}

func (s *Settings) GetTranslation_PermittedOrigins() []string {
	var origins []string
	for _, v := range s.Translation.Permitted.Origins { // / SECURITY SESSIONs
		if v.Name != "" {
			origins = append(origins, v.Name)
		}
	}
	return origins
}

func (s *Settings) GetTranslation_PermittedLocales() []struct {
	Key  string "toml:\"key\""
	Name string "toml:\"name\""
} {
	return s.Translation.Permitted.Locales
}

func (s *Settings) IsPermittedTranslationLocale(in string) bool {
	for _, v := range s.Translation.Permitted.Locales {
		if v.Key == in {
			return true
		}
	}
	return false
}

func (s *Settings) IsPermittedTranslationOrigin(in string) bool {
	for _, v := range s.Translation.Permitted.Origins {
		if v.Name == in {
			return true
		}
	}
	return false
}

func (s *Settings) GetLocaleName(in string) string {
	for _, v := range s.Translation.Permitted.Locales {
		if v.Key == in {
			return v.Name
		}
	}
	return ""
}
